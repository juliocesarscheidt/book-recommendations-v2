package controller

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/application/usecase"
	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func GetRecommendation(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		uuid, _ := params["uuid"]

		getRecommendationRequestDTO := dto.GetRecommendationRequestDTO{
			UserUuid: uuid,
		}
		recommendations, err := usecase.GetRecommendation(getRecommendationRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		fmt.Printf("Response :: %v\n", recommendations)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: recommendations})
	}
}
