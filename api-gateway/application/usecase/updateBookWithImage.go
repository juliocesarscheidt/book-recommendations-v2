package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func UpdateBookWithImage(updateBookWithImageRequestDTO dto.UpdateBookWithImageRequestDTO, amqpClient *adapter.AmqpClient) (error) {
	updateBookWithImageRequestDTOJson, err := json.Marshal(&updateBookWithImageRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return err
	}
	fmt.Println(string(updateBookWithImageRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(updateBookWithImageRequestDTOJson), "")
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.update_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return err
	}

	return nil
}
