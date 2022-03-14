package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func UpdateBook(updateBookRequestDTO dto.UpdateBookRequestDTO, amqpClient *adapter.AmqpClient) (error) {
	updateBookRequestDTOJson, err := json.Marshal(&updateBookRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return err
	}
	fmt.Println(string(updateBookRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(updateBookRequestDTOJson), "")
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.update_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return err
	}

	return nil
}
