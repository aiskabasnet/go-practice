package service

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type GoogleCloudStorageService interface {
	UploadFileToCloudStorage(multipart.File, string) (string, error)
	// UploadThumbnailToCloudStorage(image.Image, string, string) (string, error)
	GenerateV4GetObjectSignedURL(string) (string, error)
}

type googleService struct{}

func NewGoogleCloudStorageService() GoogleCloudStorageService {
	return &googleService{}
}
func (g *googleService) UploadFileToCloudStorage(file multipart.File, fileName string) (string, error) {
	var bucketName = os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		fmt.Println("No bucket name in env")
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("serviceAccountKey.json"))
	if err != nil {
		return "", err
	}
	_, err = client.Bucket(bucketName).Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		return "", fmt.Errorf("%v", err)
	}
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}
	u, err := url.ParseRequestURI("/" + bucketName + "/" + wc.Attrs().Name)
	if err != nil {
		return "", err
	}
	return u.EscapedPath(), nil
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
