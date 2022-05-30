package common

import (
	"os"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() (*s3.S3, error) {
	region := os.Getenv("AWS_DEFAULT_REGION")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	return s3.New(sess), nil
}

func PutS3File(file io.ReadSeeker, bucketName string, filenameBucketKey string, contentType string) (*s3.PutObjectOutput, error) {
	s3Client, err := GetS3Client()
	if err != nil {
		return nil, err
	}
	input := &s3.PutObjectInput{
		Body: file,
		Bucket: aws.String(bucketName),
		Key: aws.String(filenameBucketKey),
		ContentType: aws.String(contentType),
		ServerSideEncryption: aws.String("AES256"), // SSE-S3 Encryption
	}
	objectOutput, err := s3Client.PutObject(input)
	if err != nil {
		return nil, err
	}
	return objectOutput, nil
}

func GenerateS3PresignUrl(bucketName string, filenameBucketKey string) (string, error) {
	s3Client, err := GetS3Client()
	if err != nil {
		return "", err
	}
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(filenameBucketKey),
	})
	url, err := req.Presign(60*60*time.Second) // 1 hour
	if err != nil {
		return "", err
	}
	return url, nil
}
