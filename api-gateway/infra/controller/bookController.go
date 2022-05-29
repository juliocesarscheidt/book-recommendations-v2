package controller

import (
	"io"
	"fmt"
	"time"
	"errors"
	"strconv"
	"strings"
	"net/http"
	// "mime/multipart"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/juliocesarscheidt/apigateway/application/usecase"
	"github.com/juliocesarscheidt/apigateway/application/dto"
	"github.com/juliocesarscheidt/apigateway/infra/adapter"
	httpmodule "github.com/juliocesarscheidt/apigateway/infra/http"
)

const (
	MAX_FILE_SIZE = 10.00
	BUCKET_NAME = "blackdevs-aws"
)

// TODO: move these methods to a separated library
func validateImageFileForm(filename string, fileExtension string, fileSizeMb float64) error {
	if fileExtension != "png" && fileExtension != "jpg" && fileExtension != "jpeg" {
		return errors.New("Invalid File Extension")
	}
	if fileSizeMb > MAX_FILE_SIZE {
		return errors.New("Invalid File Size")
	}
	return nil
}

func appendBucketPathToFilename(filename string) string {
	return fmt.Sprintf("book-recommendations-files/%s", filename)
}

func generateFilenameRandBucketKey(filename string, fileExtension string) (string, string) {
	now := time.Now()
	nowSecs := now.Unix()
	filenameRand := fmt.Sprintf("%s-%d.%s", filename, nowSecs, fileExtension)
	filenameBucketKey := appendBucketPathToFilename(filenameRand)
	return filenameRand, filenameBucketKey
}

func getS3Client() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1")},
	)
	if err != nil {
		return nil, err
	}
	return s3.New(sess), nil
}

func putS3File(file io.ReadSeeker, bucketName string, filenameBucketKey string) (*s3.PutObjectOutput, error) {
	s3Client, err := getS3Client()
	if err != nil {
		return nil, err
	}
	input := &s3.PutObjectInput{
		Body: file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(filenameBucketKey),
	}
	objectOutput, err := s3Client.PutObject(input)
	if err != nil {
		return nil, err
	}
	return objectOutput, nil
}

func generateS3PresignUrl(bucketName string, filenameBucketKey string) (string, error) {
	s3Client, err := getS3Client()
	if err != nil {
		return "", err
	}
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filenameBucketKey),
	})
	url, err := req.Presign(60*60*time.Second) // 1 hour
	if err != nil {
		return "", err
	}
	return url, nil
}

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

		if err := validateImageFileForm(filename, fileExtension, fileSizeMb); err != nil {
			ThrowBadRequest(w, err.Error())
			return
		}
		filenameRand, filenameBucketKey := generateFilenameRandBucketKey(filename, fileExtension)
		fmt.Println("filenameBucketKey :: " + filenameBucketKey)

		objectOutput, err := putS3File(file, BUCKET_NAME, filenameBucketKey)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		fmt.Println(objectOutput)

		createBookRequestDTO := dto.CreateBookRequestDTO{
			Title: r.FormValue("title"),
			Author: r.FormValue("author"),
			Genre: r.FormValue("genre"),
			Image: filenameRand,
		}
		fmt.Println(createBookRequestDTO)

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

		filenameBucketKey := appendBucketPathToFilename(book.Image)

		url, err := generateS3PresignUrl(BUCKET_NAME, filenameBucketKey)
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

		if err := validateImageFileForm(filename, fileExtension, fileSizeMb); err != nil {
			ThrowBadRequest(w, err.Error())
			return
		}
		filenameRand, filenameBucketKey := generateFilenameRandBucketKey(filename, fileExtension)
		fmt.Println("filenameBucketKey :: " + filenameBucketKey)

		objectOutput, err := putS3File(file, BUCKET_NAME, filenameBucketKey)
		if err != nil {
			ThrowInternalServerError(w, err.Error())
			return
		}
		fmt.Println(objectOutput)

		updateBookRequestDTO := dto.UpdateBookRequestDTO{
			Uuid: book_uuid,
			Title: r.FormValue("title"),
			Author: r.FormValue("author"),
			Genre: r.FormValue("genre"),
			Image: filenameRand,
		}
		fmt.Println(updateBookRequestDTO)

		err = usecase.UpdateBook(updateBookRequestDTO, amqpClient)
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
