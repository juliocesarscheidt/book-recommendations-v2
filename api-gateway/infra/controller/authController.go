package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	"github.com/juliocesarscheidt/apigateway/application/usecase"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func UserSignUp(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userSignUp userpb.UserSignUp
		if err := json.NewDecoder(r.Body).Decode(&userSignUp); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		req := &userpb.UserSignUpRequest{
			UserSignUp: &userSignUp,
		}

		access_token, err := usecase.UserSignUp(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", access_token)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: access_token})
	}
}

func UserSignIn(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userSignIn userpb.UserSignIn
		if err := json.NewDecoder(r.Body).Decode(&userSignIn); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		req := &userpb.UserSignInRequest{
			UserSignIn: &userSignIn,
		}

		access_token, err := usecase.UserSignIn(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", access_token)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: access_token})
	}
}
