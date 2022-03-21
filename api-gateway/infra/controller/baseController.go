package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"google.golang.org/grpc/status"
	// "google.golang.org/grpc/codes"

	"github.com/juliocesarscheidt/apigateway/application/usecase"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func CheckUserExists(uuid string, grpcClient adapter.GrpcClient) error {
	// check if user exists
	getUserRequest := &userpb.GetUserRequest{
		Uuid: uuid,
	}
	user, err := usecase.GetUser(getUserRequest, grpcClient)
	if err != nil {
		return err
	}
	fmt.Printf("Response :: %v\n", user)

	return nil
}

func ThrowInternalServerError(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: message})
}

func ThrowBadRequest(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: message})
}

func ThrowUnauthorized(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: message})
}

func ThrowNotFound(w http.ResponseWriter, message string) {
	fmt.Println(message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&httpmodule.HttpResponseError{Message: message})
}

func HandleError(w http.ResponseWriter, err error) {
	stat := status.Convert(err)
	fmt.Println(fmt.Sprintf("Code %v", stat.Code()))
	fmt.Println(fmt.Sprintf("Message %v", stat.Message()))

	if stat.Message() == "Not Found" {
		ThrowNotFound(w, stat.Message())
		return

	} else if stat.Message() == "Bad Request" {
		ThrowBadRequest(w, stat.Message())
		return

	}	else if stat.Message() == "Invalid Rate" {
		ThrowBadRequest(w, stat.Message())
		return

	} else if stat.Message() == "Invalid Email or Password" {
		ThrowBadRequest(w, stat.Message())
		return

	}	else if stat.Message() == "Duplicate Email" {
		ThrowBadRequest(w, stat.Message())
		return

	} else if stat.Message() == "Internal Server Error" {
		ThrowInternalServerError(w, stat.Message())
		return
	}

	ThrowInternalServerError(w, err.Error())
}


