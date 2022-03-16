package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	"github.com/juliocesarscheidt/apigateway/application/usecase"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func CreateUser(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userRequest userpb.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		req := &userpb.CreateUserRequest{
			UserRequest: &userRequest,
		}

		uuid, err := usecase.CreateUser(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", uuid)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: uuid})
	}
}

func GetUser(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		user_uuid, _ := params["user_uuid"]

		req := &userpb.GetUserRequest{
			Uuid: user_uuid,
		}
		user, err := usecase.GetUser(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", user)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: user})
	}
}

func UpdateUser(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		user_uuid, _ := params["user_uuid"]

		var userRequest userpb.UserRequest
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			HandleError(w, err)
			return
		}
		req := &userpb.UpdateUserRequest{
			Uuid: user_uuid,
			UserRequest: &userRequest,
		}

		err := usecase.UpdateUser(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func DeleteUser(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		user_uuid, _ := params["user_uuid"]

		req := &userpb.DeleteUserRequest{
			Uuid: user_uuid,
		}
		err := usecase.DeleteUser(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ListUser(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := strconv.ParseInt(r.FormValue("page"), 10, 64)
		size, _ := strconv.ParseInt(r.FormValue("size"), 10, 64)

		req := &userpb.ListUserRequest{
			Page: page,
			Size: size,
		}
		users, err := usecase.ListUser(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", users)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: users})
	}
}
