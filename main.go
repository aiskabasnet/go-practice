package main

import (
	Routes "go-practice/api/routes"
	"go-practice/api/seeds"
	"go-practice/infrastructure"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading")
	}
	fb := infrastructure.InitializeFirebase()
	seeds.LoadAdmin(fb)
	log.Println("Firebase done", fb)
	r := Routes.SetupRouter()
	//running
	log.Fatal(r.Run(os.Getenv("PORT")))
}
