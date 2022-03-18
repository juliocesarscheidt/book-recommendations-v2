package router

import (
	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/infra/middleware"
	"github.com/juliocesarscheidt/apigateway/infra/controller"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

// GetRouter - it returns the mux Router with the injected middlewares
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// inject main middleware
	logger := middleware.Logger{Format: "[URI] %s"}
	router.Use(logger.LoggerMiddleware)

	return router
}

func InjectRoutes(router *mux.Router, grpcClient adapter.GrpcClient, amqpClient *adapter.AmqpClient, redisClient adapter.RedisClient) {
	// auth routes (no authentication)
	router.Path("/v1/auth/signup").HandlerFunc(controller.UserSignUp(grpcClient)).Methods("POST")
	router.Path("/v1/auth/signin").HandlerFunc(controller.UserSignIn(grpcClient)).Methods("POST")

	// user sub router
	subRouterUser := router.PathPrefix("/v1/user").Subrouter()
	subRouterUser.Use(middleware.AuthenticationMiddleware(redisClient))
	// user routes (/v1/user)
	subRouterUser.Path("").HandlerFunc(controller.CreateUser(grpcClient)).Methods("POST")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUser(grpcClient)).Methods("GET")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateUser(grpcClient)).Methods("PUT")
	subRouterUser.Path("/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUser(grpcClient)).Methods("DELETE")
	subRouterUser.Path("").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUser(grpcClient)).Methods("GET")
	// user rates routes (/v1/user/rate)
	subRouterUser.Path("/rate").HandlerFunc(controller.UpsertUserRate(grpcClient, amqpClient)).Methods("POST")
	subRouterUser.Path("/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetUserRate(grpcClient)).Methods("GET")
	subRouterUser.Path("/rate/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteUserRate(grpcClient)).Methods("DELETE")
	subRouterUser.Path("/rate").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListUserRate(grpcClient)).Methods("GET")

	// book sub router
	subRouterBook := router.PathPrefix("/v1/book").Subrouter()
	subRouterBook.Use(middleware.AuthenticationMiddleware(redisClient))
	// book routes (/v1/book)
	subRouterBook.Path("").HandlerFunc(controller.CreateBook(amqpClient)).Methods("POST")
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetBook(amqpClient)).Methods("GET")
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.UpdateBook(amqpClient)).Methods("PUT")
	subRouterBook.Path("/{book_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.DeleteBook(amqpClient)).Methods("DELETE")
	subRouterBook.Path("").Queries("page", "{page}", "size", "{size}").HandlerFunc(controller.ListBook(amqpClient)).Methods("GET")

	// recommendation sub router
	subRouterRecommendation := router.PathPrefix("/v1/recommendation").Subrouter()
	subRouterRecommendation.Use(middleware.AuthenticationMiddleware(redisClient))
	// recommendation routes (/v1/recommendation)
	subRouterRecommendation.Path("/user/{user_uuid:[a-zA-Z0-9]{24}}").HandlerFunc(controller.GetRecommendation(grpcClient, amqpClient)).Methods("GET")
}
