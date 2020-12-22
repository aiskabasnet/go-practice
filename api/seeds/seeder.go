package seeds

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
)

//LoadAdmin => Add admin
func LoadAdmin(firebaseAuth *auth.Client) {
	_, err := firebaseAuth.GetUserByEmail(context.Background(), "admin@go-practice.com")
	if err != nil {
		log.Println("Creating admin...")
		params := (&auth.UserToCreate{}).Email("admin@go-practice.com").Password("aiska111").EmailVerified(true)
		user, err := firebaseAuth.CreateUser(context.Background(), params)
		if err != nil {
			log.Fatalf("Error creating admin: %v\n", err)
			return
		}

		claims := map[string]interface{}{"admin": true}
		err = firebaseAuth.SetCustomUserClaims(context.Background(), user.UID, claims)
		if err != nil {
			log.Fatalf("Error adding claim: %v\n", err)
			return
		}
	}
	log.Println("Admin already exists")
}
