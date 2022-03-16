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
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

func CreateBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var createBookRequestDTO dto.CreateBookRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&createBookRequestDTO); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		uuid, err := usecase.CreateBook(createBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: uuid})
	}
}

func GetBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		book_uuid, _ := params["book_uuid"]

		getBookRequestDTO := dto.GetBookRequestDTO{
			Uuid: book_uuid,
		}
		book, err := usecase.GetBook(getBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		fmt.Printf("Response :: %v\n", book)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: book})
	}
}

func UpdateBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		book_uuid, _ := params["book_uuid"]

		var updateBookRequestDTO dto.UpdateBookRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&updateBookRequestDTO); err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		updateBookRequestDTO.Uuid =  book_uuid
		err := usecase.UpdateBook(updateBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func DeleteBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		book_uuid, _ := params["book_uuid"]

		deleteBookRequestDTO := dto.DeleteBookRequestDTO{
			Uuid: book_uuid,
		}
		err := usecase.DeleteBook(deleteBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ListBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		page, _ := strconv.ParseInt(r.FormValue("page"), 10, 64)
		size, _ := strconv.ParseInt(r.FormValue("size"), 10, 64)

		listBookRequestDTO := dto.ListBookRequestDTO{
			Page: page,
			Size: size,
		}
		books, err := usecase.ListBook(listBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		fmt.Printf("Response :: %v\n", books)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: books})
	}
}
