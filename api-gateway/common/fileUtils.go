package common

import (
	"fmt"
	"time"
	"errors"
)

const (
	MAX_FILE_SIZE = 5.00 // 5 MB
)

func ValidateImageFileForm(filename string, fileExtension string, fileSizeMb float64) error {
	if fileExtension != "png" && fileExtension != "jpg" && fileExtension != "jpeg" {
		return errors.New("Invalid File Extension")
	}
	if fileSizeMb > MAX_FILE_SIZE {
		return errors.New("Invalid File Size")
	}
	return nil
}

func AppendBucketPathToFilename(filename string) string {
	return fmt.Sprintf("book-recommendations-files/%s", filename)
}

func GenerateFilenameRandBucketKey(filename string, fileExtension string) (string, string) {
	now := time.Now()
	nowSecs := now.Unix()
	filenameRand := fmt.Sprintf("%s-%d.%s", filename, nowSecs, fileExtension)
	filenameBucketKey := AppendBucketPathToFilename(filenameRand)
	return filenameRand, filenameBucketKey
}
