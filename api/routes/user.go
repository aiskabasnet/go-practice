package routes

import (
	controllers "go-practice/api/controllers/users"
	"go-practice/api/middleware"
	repository "go-practice/api/repository/user"
	fbService "go-practice/api/service"
	service "go-practice/api/service/user"
	"log"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.RouterGroup, db *gorm.DB, fb *auth.Client) {
	fbAuthService := fbService.NewFirebaseService(fb)
	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}
	// gCloudService := gcloudService.NewGoogleCloudStorageService()
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, fbAuthService)

	route.POST("/", userController.CreateUser)
	route.GET("/users", middleware.Auth(fbAuthService), userController.GetUsers)
}
