package middleware

import (
	"log"
	"fmt"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/juliocesarscheidt/apigateway/common"
	"github.com/juliocesarscheidt/apigateway/infra/controller"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func AuthenticationMiddleware(redisClient adapter.RedisClient) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := common.ExtractTokenString(r.Header.Get("Authorization"))
			if err != nil {
				controller.ThrowUnauthorized(w, "Invalid Authorization Header")
				return
			}

			token, err := common.ParseTokenUsingPublicKey(tokenString)
			if err != nil {
				log.Printf("ERROR :: %s\n", err)
				if errors.Is(err, jwt.ErrTokenMalformed) {
					controller.ThrowUnauthorized(w, "Malformed Token")
				} else if (errors.Is(err, jwt.ErrTokenExpired) ||
					errors.Is(err, jwt.ErrTokenNotValidYet)) {
					controller.ThrowUnauthorized(w, "Expired Token")
				} else {
					controller.ThrowUnauthorized(w, "Unauthorized")
				}
				return
			}

			claims, err := common.ValidateClaims(token)
			if err != nil {
				controller.ThrowUnauthorized(w, err.Error())
				return
			}

			// check token on redis
			userUuid := claims["uuid"]
			log.Printf("userUuid :: %v\n", userUuid)
			result, err := redisClient.Get(fmt.Sprintf("/user/bearer/%s", userUuid))
			if (result == nil || err != nil) {
				controller.ThrowUnauthorized(w, "Unauthorized")
				return
			}

			// set user uuid on header
			r.Header.Add("X-Session-User", userUuid.(string))

			// call next handler
			next.ServeHTTP(w, r)
		})
	}
}
