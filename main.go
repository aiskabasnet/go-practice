package main

import (
	"fmt"
	"go-practice/Routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
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

func ProvideString() CustomString {
	var str CustomString = "Hello world"
	return str
}
func ProvideLogger() LoggerInterface {
	return &MyLogger{Prefix: "Konichiwa"}
}
func RunMe(logger LoggerInterface, data CustomString) {
	logData := logger.LogMe("dependency Injection working")
	fmt.Println(logData)
	fmt.Println("Received string", data)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading")
	}
	fx.New(
		fx.Provide(ProvideString),
		fx.Provide(ProvideLogger),
		fx.Invoke(RunMe),
	).Run()
	r := Routes.SetupRouter()
	//running
	log.Fatal(r.Run(os.Getenv("PORT")))
}
