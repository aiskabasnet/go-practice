package main

import (
	"go-practice/api/seeds"
	"go-practice/infrastructure"
	"go-practice/utils"
)

func main() {
	utils.LoadEnv()
	db := infrastructure.SetupModels()
	fb := infrastructure.InitializeFirebase()
	seeds.LoadAdmin(fb)
	infrastructure.SetupRoutes(db, fb)
}
