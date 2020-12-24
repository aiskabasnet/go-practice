package utils

import (
	"math/rand"
	"time"
)

var digitsAndNumbers = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomInvitationCode generate random string
func GenerateRandomInvitationCode(codeLength int) string {
	code := make([]rune, codeLength)
	for i := range code {
		code[i] = digitsAndNumbers[rand.Intn(len(digitsAndNumbers))]
	}
	return string(code)
}
