package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/infra/controller"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

type Logger struct {
	Format string
}

func (logger *Logger) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Printf(r.Header["Authorization"])
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
	// users routes
	router.Path("/v1/user").HandlerFunc(controller.CreateUser(grpcClient)).Methods("POST")
	router.Path("/v1/user/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUser(grpcClient)).Methods("GET")
	router.Path("/v1/user/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateUser(grpcClient)).Methods("PUT")
	router.Path("/v1/user/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUser(grpcClient)).Methods("DELETE")
	router.Path("/v1/user").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUser(grpcClient)).Methods("GET")

	// users rates routes
	router.Path("/v1/user/rate").HandlerFunc(controller.UpsertUserRate(grpcClient, amqpClient)).Methods("POST")
	router.Path("/v1/user/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUserRate(grpcClient)).Methods("GET")
	router.Path("/v1/user/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUserRate(grpcClient)).Methods("DELETE")
	router.Path("/v1/user/rate").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUserRate(grpcClient)).Methods("GET")

	// books routes
	router.Path("/v1/book").HandlerFunc(controller.CreateBook(amqpClient)).Methods("POST")
	router.Path("/v1/book/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetBook(amqpClient)).Methods("GET")
	router.Path("/v1/book/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateBook(amqpClient)).Methods("PUT")
	router.Path("/v1/book/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteBook(amqpClient)).Methods("DELETE")
	router.Path("/v1/book").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListBook(amqpClient)).Methods("GET")

	// recommendations routes
	router.Path("/v1/recommendation/user/{uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetRecommendation(amqpClient)).Methods("GET")
}
