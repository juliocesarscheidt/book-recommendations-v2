package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func ListBook(listBookRequestDTO dto.ListBookRequestDTO, amqpClient *adapter.AmqpClient) ([]dto.GetBookResponseDTO, error) {
	emptyResponse := []dto.GetBookResponseDTO{}
	tempQueue, err := amqpClient.MakeTempQueue()
	if err != nil {
		fmt.Errorf("Failed to make temp queue: %v", err)
		return emptyResponse, err
	}

	listBookRequestDTOJson, err := json.Marshal(&listBookRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return emptyResponse, err
	}
	fmt.Println(string(listBookRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(listBookRequestDTOJson), tempQueue.Name)
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.list_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return emptyResponse, err
	}

	consumerName := fmt.Sprintf("consumer_%s", tempQueue.Name)
	messageBody, err := amqpClient.ConsumeMessageFromTempQueue(tempQueue.Name, consumerName)
	if err != nil {
		fmt.Errorf("Failed to consume from temp queue: %v", err)
		return emptyResponse, err
	}

	bodyDTO := dto.ListBookResponseDTO{}
	if err := json.Unmarshal(messageBody, &bodyDTO); err != nil {
		fmt.Errorf("Failed to unmarshal JSON: %v", err)
		return emptyResponse, err
	}

	response := []dto.GetBookResponseDTO{}
	if bodyDTO.Books != nil {
		response = bodyDTO.Books
	}

	return response, nil
}
