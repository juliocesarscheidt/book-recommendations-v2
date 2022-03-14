package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func CreateBook(createBookRequestDTO dto.CreateBookRequestDTO, amqpClient *adapter.AmqpClient) (string, error) {
	emptyResponse := ""
	tempQueue, err := amqpClient.MakeTempQueue()
	if err != nil {
		fmt.Errorf("Failed to make temp queue: %v", err)
		return emptyResponse, err
	}

	createBookRequestDTOJson, err := json.Marshal(&createBookRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return emptyResponse, err
	}
	fmt.Println(string(createBookRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(createBookRequestDTOJson), tempQueue.Name)
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.create_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return emptyResponse, err
	}

	consumerName := fmt.Sprintf("consumer_%s", tempQueue.Name)
	messageBody, err := amqpClient.ConsumeMessageFromTempQueue(tempQueue.Name, consumerName)
	if err != nil {
		fmt.Errorf("Failed to consume from temp queue: %v", err)
		return emptyResponse, err
	}

	bodyDTO := dto.CreateBookResponseDTO{}
	if err := json.Unmarshal(messageBody, &bodyDTO); err != nil {
		fmt.Errorf("Failed to unmarshal JSON: %v", err)
		return emptyResponse, err
	}

	return bodyDTO.Uuid, nil
}
