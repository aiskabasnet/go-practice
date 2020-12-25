package service

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// FirebaseService : represent the firebase's services
type FirebaseService interface {
	VerifyToken(string) (*auth.Token, error)
	CreateUser(string, string) (string, error)
	DeleteUser(string) error
	SetClaim(string, gin.H) error
	GetUser(string) (*auth.UserRecord, error)
	UpdateUser(string, bool) (string, error)
	CreateToken(string) (string, error)
}

type firebaseService struct {
	Firebase *auth.Client
}

// NewFirebaseService : get injected firebase
func NewFirebaseService(fb *auth.Client) FirebaseService {
	return &firebaseService{Firebase: fb}
}

func (fb *firebaseService) VerifyToken(idToken string) (*auth.Token, error) {
	token, err := fb.Firebase.VerifyIDToken(context.Background(), idToken)
	return token, err
}

func (fb *firebaseService) CreateUser(email string, password string) (string, error) {
	params := (&auth.UserToCreate{}).Email(email).Password(password)
	u, err := fb.Firebase.CreateUser(context.Background(), params)
	if err != nil {
		return "", err
	}
	return u.UID, err
}

func (fb *firebaseService) DeleteUser(id string) error {
	err := fb.Firebase.DeleteUser(context.Background(), id)
	return err
}
func (fb *firebaseService) SetClaim(id string, claims gin.H) error {
	err := fb.Firebase.SetCustomUserClaims(context.Background(), id, claims)
	return err
}
func (fb *firebaseService) GetUser(uid string) (*auth.UserRecord, error) {
	user, err := fb.Firebase.GetUser(context.Background(), uid)
	return user, err
}
func (fb *firebaseService) UpdateUser(uid string, verified bool) (string, error) {
	params := (&auth.UserToUpdate{}).EmailVerified(verified)
	u, err := fb.Firebase.UpdateUser(context.Background(), uid, params)
	if err != nil {
		return "", err
	}
	return u.UID, err
}
func (fb *firebaseService) CreateToken(id string) (string, error) {
	token, err := fb.Firebase.CustomToken(context.Background(), id)
	return token, err
}
