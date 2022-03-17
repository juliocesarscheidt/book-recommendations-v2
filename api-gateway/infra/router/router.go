package router

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"net/http"
	"io/ioutil"
	"crypto/rsa"
	"path/filepath"
	"encoding/json"

	"github.com/gorilla/mux"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/juliocesarscheidt/apigateway/infra/controller"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

type Logger struct {
	Format string
}

func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	absLocation, _ := filepath.Abs(location)
	keyData, err := ioutil.ReadFile(absLocation)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func AuthenticateMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header["Authorization"])

			tokenString := r.Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: "Unauthorized"})
				return
			}
			// log.Printf("tokenString :: %s\n", tokenString)

			publicKey, _ := LoadRSAPublicKeyFromDisk("keys/jwtencyptionkey.pub")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return publicKey, nil
			})

			if err != nil {
				if errors.Is(err, jwt.ErrTokenMalformed) {
					fmt.Println("That's not even a token")
				} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
					// Token is either expired or not active yet
					fmt.Println("Timing is everything")
				}
				log.Printf("ERROR :: %s\n", err)

				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: "Unauthorized"})
				return
			}

			log.Printf("token :: %s\n", token)

			log.Printf("token.Valid :: %v\n", token.Valid)

			log.Printf("token.Signature :: %v\n", token.Signature)

			claims, ok := token.Claims.(jwt.MapClaims)
			log.Printf("claims :: %v\n", claims)
			log.Printf("claims[\"aud\"] :: %v\n", claims["aud"])
			log.Printf("claims[\"iss\"] :: %v\n", claims["iss"])

			if claims["aud"] != "urn:book-recommendations:api-gateway" || claims["iss"] != "urn:book-recommendations:users-microservice" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: "Unauthorized"})
				return
			}

			log.Printf("ok :: %v\n", ok)

			if ! ok || ! token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: "Unauthorized"})
				return
			}

			// call next handler
			next.ServeHTTP(w, r)
		})
	}
}

func (logger *Logger) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(logger.Format, r.RequestURI)
		// call next handler
		next.ServeHTTP(w, r)
	})
}

// GetRouter - it returns the mux Router with the injected middlewares
func GetRouter() *mux.Router {
	router := mux.NewRouter()

	logger := Logger{Format: "[URI] %s"}
	router.Use(logger.LoggerMiddleware)

	return router
}

func InjectRoutes(router *mux.Router, grpcClient adapter.GrpcClient, amqpClient *adapter.AmqpClient) {
	// auth routes
	router.Path("/v1/auth/signup").HandlerFunc(controller.UserSignUp(grpcClient)).Methods("POST")
	router.Path("/v1/auth/signin").HandlerFunc(controller.UserSignIn(grpcClient)).Methods("POST")

	subRouterUser := router.PathPrefix("/v1/user").Subrouter()
	subRouterUser.Use(AuthenticateMiddleware())
	// users routes
	subRouterUser.Path("").HandlerFunc(controller.CreateUser(grpcClient)).Methods("POST")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUser(grpcClient)).Methods("GET")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateUser(grpcClient)).Methods("PUT")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUser(grpcClient)).Methods("DELETE")
	subRouterUser.Path("").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUser(grpcClient)).Methods("GET")
	// users rates routes
	subRouterUser.Path("/rate").HandlerFunc(controller.UpsertUserRate(grpcClient, amqpClient)).Methods("POST")
	subRouterUser.Path("/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUserRate(grpcClient)).Methods("GET")
	subRouterUser.Path("/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUserRate(grpcClient)).Methods("DELETE")
	subRouterUser.Path("/rate").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUserRate(grpcClient)).Methods("GET")

	subRouterBook := router.PathPrefix("/v1/book").Subrouter()
	subRouterBook.Use(AuthenticateMiddleware())
	// books routes
	subRouterBook.Path("/v1/book").HandlerFunc(controller.CreateBook(amqpClient)).Methods("POST")
	subRouterBook.Path("/v1/book/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetBook(amqpClient)).Methods("GET")
	subRouterBook.Path("/v1/book/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateBook(amqpClient)).Methods("PUT")
	subRouterBook.Path("/v1/book/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteBook(amqpClient)).Methods("DELETE")
	subRouterBook.Path("/v1/book").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListBook(amqpClient)).Methods("GET")

	subRouterRecommendation := router.PathPrefix("/v1/recommendation").Subrouter()
	subRouterRecommendation.Use(AuthenticateMiddleware())
	// recommendations routes
	subRouterRecommendation.Path("/v1/recommendation/user/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetRecommendation(grpcClient, amqpClient)).Methods("GET")
}
