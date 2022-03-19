package usecase

import (
	"fmt"
	"errors"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func GetBook(getBookRequestDTO dto.GetBookRequestDTO, amqpClient *adapter.AmqpClient) (dto.GetBookResponseDTO, error) {
	emptyResponse := dto.GetBookResponseDTO{}
	tempQueue, err := amqpClient.MakeTempQueue()
	if err != nil {
		fmt.Errorf("Failed to make temp queue: %v", err)
		return emptyResponse, err
	}

	getBookRequestDTOJson, err := json.Marshal(&getBookRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return emptyResponse, err
	}
	fmt.Println(string(getBookRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(getBookRequestDTOJson), tempQueue.Name)
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.get_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return emptyResponse, err
	}

	consumerName := fmt.Sprintf("consumer_%s", tempQueue.Name)
	messageBody, err := amqpClient.ConsumeMessageFromTempQueue(tempQueue.Name, consumerName)
	if err != nil {
		fmt.Errorf("Failed to consume from temp queue: %v", err)
		return emptyResponse, err
	}

	bodyDTO := dto.GetBookResponseDTO{}
	if err := json.Unmarshal(messageBody, &bodyDTO); err != nil {
		fmt.Errorf("Failed to unmarshal JSON: %v", err)
		return emptyResponse, err
	}

	if (bodyDTO == emptyResponse) {
		return emptyResponse, errors.New("Not Found")
	}

	return bodyDTO, nil
}
