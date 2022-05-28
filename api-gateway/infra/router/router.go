package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	"github.com/juliocesarscheidt/apigateway/infra/controller"
	"github.com/juliocesarscheidt/apigateway/infra/middleware"
)

// GetRouter - it returns the mux Router with the injected middlewares
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// inject main middleware
	logger := middleware.Logger{Format: "[URI] %s"}
	router.Use(logger.LoggerMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	return router
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "alive"})
}

func InjectRoutes(router *mux.Router, grpcClient adapter.GrpcClient, amqpClient *adapter.AmqpClient, redisClient adapter.RedisClient) {
	// healthcheck route (no authentication)
	router.Path("/api/healthcheck").HandlerFunc(HealthCheckHandler).Methods(http.MethodGet)

	// auth routes (no authentication)
	router.Path("/api/auth/signup").HandlerFunc(controller.UserSignUp(grpcClient)).Methods(http.MethodPost)
	router.Path("/api/auth/signin").HandlerFunc(controller.UserSignIn(grpcClient)).Methods(http.MethodPost)

	// user sub router
	subRouterUser := router.PathPrefix("/api/user").Subrouter()
	subRouterUser.Use(middleware.AuthenticationMiddleware(redisClient))
	// user routes (/api/user)
	subRouterUser.Path("/me").HandlerFunc(controller.GetCurrentUserInfo(grpcClient)).Methods(http.MethodGet)
	subRouterUser.Path("").HandlerFunc(controller.CreateUser(grpcClient)).Methods(http.MethodPost)
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUser(grpcClient)).Methods(http.MethodGet)
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateUser(grpcClient)).Methods(http.MethodPut)
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUser(grpcClient)).Methods(http.MethodDelete)
	subRouterUser.Path("").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUser(grpcClient)).Methods(http.MethodGet)
	// user rates routes (/api/user/rating)
	subRouterUser.Path("/rating").HandlerFunc(controller.UpsertUserRate(grpcClient, amqpClient)).Methods(http.MethodPost)
	subRouterUser.Path("/rating/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUserRate(grpcClient)).Methods(http.MethodGet)
	subRouterUser.Path("/rating/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUserRate(grpcClient)).Methods(http.MethodDelete)
	subRouterUser.Path("/rating").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUserRate(grpcClient)).Methods(http.MethodGet)

	// book sub router
	subRouterBook := router.PathPrefix("/api/book").Subrouter()
	subRouterBook.Use(middleware.AuthenticationMiddleware(redisClient))
	// book routes (/api/book)
	subRouterBook.Path("").HandlerFunc(controller.CreateBook(amqpClient)).Methods(http.MethodPost)
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetBook(amqpClient)).Methods(http.MethodGet)
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}/file/url").HandlerFunc(controller.GetBookPresignUrl(amqpClient, redisClient)).Methods(http.MethodGet)
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateBook(amqpClient)).Methods(http.MethodPut)
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}/file").HandlerFunc(controller.UpdateBookWithFile(amqpClient, redisClient)).Methods(http.MethodPut)
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteBook(amqpClient)).Methods(http.MethodDelete)
	subRouterBook.Path("").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListBook(amqpClient)).Methods(http.MethodGet)

	// recommendation sub router
	subRouterRecommendation := router.PathPrefix("/api/recommendation").Subrouter()
	subRouterRecommendation.Use(middleware.AuthenticationMiddleware(redisClient))
	// recommendation routes (/api/recommendation/user)
	subRouterRecommendation.Path("/user/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetRecommendation(grpcClient, amqpClient)).Methods(http.MethodGet)
}
