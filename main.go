package main

import (
	"fmt"
	"go-practice/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type CustomString string

type MyLogger struct {
	Prefix string
}

func (lg *MyLogger) LogMe(data string) string {
	return fmt.Sprintf("%v:%v\n", lg.Prefix, data)
}

type LoggerInterface interface {
	LogMe(str string) string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading")
	}
	r := Routes.SetupRouter()
	//running
	log.Fatal(r.Run(os.Getenv("PORT")))
}
