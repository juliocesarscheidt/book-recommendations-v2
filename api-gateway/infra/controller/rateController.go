package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/application/usecase"
	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	userpb "github.com/juliocesarscheidt/apigateway/pb"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func UpsertUserRate(grpcClient adapter.GrpcClient, amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userRateRequest userpb.UserRateRequest
		if err := json.NewDecoder(r.Body).Decode(&userRateRequest); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		// check if user exists
		errCheckUserExists := CheckUserExists(userRateRequest.GetUserUuid(), grpcClient); if errCheckUserExists != nil {
			HandleError(w, errCheckUserExists)
			return
		}

		// check if book exists
		getBookRequestDTO := dto.GetBookRequestDTO{
			Uuid: userRateRequest.GetBookUuid(),
		}
		_, err := usecase.GetBook(getBookRequestDTO, amqpClient)
		if err != nil {
			ThrowNotFound(w, "Not Found")
			return
		}

		req := &userpb.UpsertUserRateRequest{
			UserRateRequest: &userRateRequest,
		}
		uuid, err := usecase.UpsertUserRate(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}
		fmt.Printf("Response :: %v\n", uuid)

		// publish on recommendations exchange the recommendation.calculate
		calculateRecommendationRequestDTO := &dto.CalculateRecommendationRequestDTO{
			UserUuid: userRateRequest.GetUserUuid(),
		}
		calculateRecommendationRequestDTOJson, err := json.Marshal(&calculateRecommendationRequestDTO)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		fmt.Println(string(calculateRecommendationRequestDTOJson))

		message := amqpClient.BuildMessageJson(string(calculateRecommendationRequestDTOJson), "")
		if err := amqpClient.PublishMessage(message, "recommendations_exchange", "recommendations.calculate", 1); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: uuid})
	}
}

func GetUserRate(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		userUuid, _ := params["user_uuid"]

		// check if user exists
		errCheckUserExists := CheckUserExists(userUuid, grpcClient); if errCheckUserExists != nil {
			HandleError(w, errCheckUserExists)
			return
		}

		req := &userpb.GetUserRateRequest{
			UserUuid: userUuid,
		}
		user, err := usecase.GetUserRate(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", user)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: user})
	}
}

func DeleteUserRate(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		userUuid, _ := params["user_uuid"]

		// check if user exists
		errCheckUserExists := CheckUserExists(userUuid, grpcClient); if errCheckUserExists != nil {
			HandleError(w, errCheckUserExists)
			return
		}

		req := &userpb.DeleteUserRateRequest{
			UserUuid: userUuid,
		}
		err := usecase.DeleteUserRate(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ListUserRate(grpcClient adapter.GrpcClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := strconv.ParseInt(r.FormValue("page"), 10, 64)
		size, _ := strconv.ParseInt(r.FormValue("size"), 10, 64)

		req := &userpb.ListUserRateRequest{
			Page: page,
			Size: size,
		}
		userRates, err := usecase.ListUserRate(req, grpcClient)
		if err != nil {
			HandleError(w, err)
			return
		}

		fmt.Printf("Response :: %v\n", userRates)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: userRates})
	}
}
