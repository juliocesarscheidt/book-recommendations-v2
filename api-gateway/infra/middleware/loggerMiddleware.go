package middleware

import (
	"log"
	"net/http"
)

type Logger struct {
	Format string
}

func (logger *Logger) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(logger.Format, r.RequestURI)
		// call next handler
		next.ServeHTTP(w, r)
	})
}
