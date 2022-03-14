package usecase

import (
	"fmt"
	"encoding/json"

	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
)

func DeleteBook(deleteBookRequestDTO dto.DeleteBookRequestDTO, amqpClient *adapter.AmqpClient) (error) {
	deleteBookRequestDTOJson, err := json.Marshal(&deleteBookRequestDTO)
	if err != nil {
		fmt.Errorf("Failed to marshal JSON: %v", err)
		return err
	}
	fmt.Println(string(deleteBookRequestDTOJson))

	message := amqpClient.BuildMessageJson(string(deleteBookRequestDTOJson), "")
	if err := amqpClient.PublishMessage(message, "books_exchange", "books.delete_book", 0); err != nil {
		fmt.Errorf("Failed to publish message: %v", err)
		return err
	}

	return nil
}
