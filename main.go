package main

import (
	"go-practice/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading")
	}
	r := Routes.SetupRouter()
	//running
	log.Fatal(r.Run(os.Getenv("PORT")))
}
