package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func GetRecommendation(getRecommendationRequestDTO dto.GetRecommendationRequestDTO, amqpClient *adapter.AmqpClient) (dto.GetRecommendationResponseDTO, error) {
	emptyResponse := dto.GetRecommendationResponseDTO{}
	tempQueue, err := amqpClient.MakeTempQueue()
	if err != nil {
		fmt.Errorf("Failed to make temp queue: %v", err)
		return emptyResponse, err
	}

	getRecommendationRequestDTOJson, err := json.Marshal(&getRecommendationRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return emptyResponse, err
	}
	fmt.Println(string(getRecommendationRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(getRecommendationRequestDTOJson), tempQueue.Name)
	if err := amqpClient.PublishMessage(message, "recommendations_exchange", "recommendations.get", 2); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return emptyResponse, err
	}

	consumerName := fmt.Sprintf("consumer_%s", tempQueue.Name)
	messageBody, err := amqpClient.ConsumeMessageFromTempQueue(tempQueue.Name, consumerName)
	if err != nil {
		fmt.Errorf("Failed to consume from temp queue: %v", err)
		return emptyResponse, err
	}

	bodyDTO := dto.GetRecommendationResponseDTO{}
	if err := json.Unmarshal(messageBody, &bodyDTO); err != nil {
		fmt.Errorf("Failed to unmarshal JSON: %v", err)
		return emptyResponse, err
	}

	return bodyDTO, nil
}
