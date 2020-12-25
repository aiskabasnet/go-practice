package infrastructure

import (
	"net/http"
	"os"

	"go-practice/api/routes"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB, fb *auth.Client) {
	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Go Practice API Server is Running..."})
	})
	routes.UserRoutes(httpRouter.Group("user"), db, fb)
	port := os.Getenv("PORT")
	if port == "" {
		httpRouter.Run()
	} else {
		httpRouter.Run(port)
	}
}
