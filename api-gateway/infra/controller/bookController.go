package controller

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"net/http"
	// "mime/multipart"
	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/juliocesarscheidt/apigateway/common"
	"github.com/juliocesarscheidt/apigateway/application/usecase"
	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

// routes
func CreateBook(amqpClient *adapter.AmqpClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		file, header, err := r.FormFile("image")
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		defer file.Close()

		filenameParts := strings.Split(header.Filename, ".")
		filename := strings.Trim(filenameParts[0], "")
		fileExtension := strings.Trim(filenameParts[1], "")
		fileSizeMb := float64(header.Size) / 1024 / 1024
		fileContentType := header.Header["Content-Type"][0]

		if err := common.ValidateImageFileForm(filename, fileExtension, fileSizeMb); err != nil {
			ThrowBadRequest(w, err.Error())
			return
		}
		filenameRand, filenameBucketKey := common.GenerateFilenameRandBucketKey(filename, fileExtension)

		bucketName := os.Getenv("BUCKET_NAME")
		_, err = common.PutS3File(file, bucketName, filenameBucketKey, fileContentType)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		createBookRequestDTO := dto.CreateBookRequestDTO{
			Title: r.FormValue("title"),
			Author: r.FormValue("author"),
			Genre: r.FormValue("genre"),
			Image: filenameRand,
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

func GetBookPresignUrl(amqpClient *adapter.AmqpClient, redisClient adapter.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		book_uuid, _ := params["book_uuid"]

		// get from cache
		cacheKey := fmt.Sprintf("/book/%s/image/url", book_uuid)
		result, err := redisClient.Get(cacheKey)
		if (err != nil) {
			fmt.Errorf("Failed to retrieve from cache: %v", err)
		}
		if (result != nil) {
			fmt.Println("Cache Hit :: " + cacheKey)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: result})
			return
		}
		fmt.Println("Cache Miss :: " + cacheKey)

		// from here, it wasn't on cache yet
		getBookRequestDTO := dto.GetBookRequestDTO{
			Uuid: book_uuid,
		}
		book, err := usecase.GetBook(getBookRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		fmt.Printf("Response :: %v\n", book)

		filenameBucketKey := common.AppendBucketPathToFilename(book.Image)

		bucketName := os.Getenv("BUCKET_NAME")
		url, err := common.GenerateS3PresignUrl(bucketName, filenameBucketKey)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		// set on cache
		if err := redisClient.Set(cacheKey, url); err != nil {
			fmt.Errorf("Failed to set on cache: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&httpmodule.HttpResponse{Data: url})
	}
}

func UpdateBookWithImage(amqpClient *adapter.AmqpClient, redisClient adapter.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		fmt.Println(params)
		book_uuid, _ := params["book_uuid"]

		file, header, err := r.FormFile("image")
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		defer file.Close()

		filenameParts := strings.Split(header.Filename, ".")
		filename := strings.Trim(filenameParts[0], "")
		fileExtension := strings.Trim(filenameParts[1], "")
		fileSizeMb := float64(header.Size) / 1024 / 1024
		fileContentType := header.Header["Content-Type"][0]

		if err := common.ValidateImageFileForm(filename, fileExtension, fileSizeMb); err != nil {
			ThrowBadRequest(w, err.Error())
			return
		}
		filenameRand, filenameBucketKey := common.GenerateFilenameRandBucketKey(filename, fileExtension)

		bucketName := os.Getenv("BUCKET_NAME")
		_, err = common.PutS3File(file, bucketName, filenameBucketKey, fileContentType)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		updateBookWithImageRequestDTO := dto.UpdateBookWithImageRequestDTO{
			Uuid: book_uuid,
			Title: r.FormValue("title"),
			Author: r.FormValue("author"),
			Genre: r.FormValue("genre"),
			Image: filenameRand,
		}

		err = usecase.UpdateBookWithImage(updateBookWithImageRequestDTO, amqpClient)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}

		// delete presign url from cache if it was there
		cacheKey := fmt.Sprintf("/book/%s/image/url", book_uuid)
		if err := redisClient.Del(cacheKey); err != nil {
			fmt.Errorf("Failed to delete from cache: %v", err)
		}

		w.WriteHeader(http.StatusAccepted)
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
		updateBookRequestDTO.Uuid = book_uuid
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
