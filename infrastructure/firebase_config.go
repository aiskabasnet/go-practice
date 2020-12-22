package infrastructure

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

//InitializeFirebase -> initialize the firebase
func InitializeFirebase() *auth.Client {
	ctx := context.Background()
	serviceAccountKeyPath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey.json")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Firebase Authentication: %v", err)
	}

	return firebaseAuth

}
