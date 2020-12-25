package service

import (
	"fmt"
	"image"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
)

type GoogleCloudStorageService interface {
	UploadFileToCloudStorage(multipart.File, string) (string, error)
	UploadThumbnailToCloudStorage(image.Image, string, string) (string, error)
	GenerateV4GetObjectSignedURL(string) (string, error)
}

type googleService struct{}

func NewGoogleCloudStorageService() GoogleCloudStorageService {
	return &googleService{}
}
func (g *googleService) GenerateV4GetObjectSignedURL(object string) (string, error) {
	var bucketName = os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		fmt.Println("No bucket name in env")
	}
	jsonKey, err := ioutil.ReadFile("key.json")
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}

	u, err := storage.SignedURL(bucketName, object, opts)
	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	return u, nil
}
