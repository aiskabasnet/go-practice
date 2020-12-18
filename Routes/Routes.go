package Routes

import (
	"fmt"
	"go-practice/handler"
	"go-practice/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome"})
	})
	userHandler := handler.NewUserHandler()
	fmt.Println(userHandler)
	productHandler := handler.NewProductHandler()
	orderHandler := handler.NewOrderHandler()
	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/signin", userHandler.SignIn)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.Authorize())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:user", userHandler.GetUser)
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductsOrdered)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.Authorize())
	{
		productRoutes.GET("/", productHandler.GetProducts)
		productRoutes.GET("/:product", productHandler.GetProductById)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.PUT("/:product", productHandler.UpdateProduct)
		productRoutes.DELETE("/:product", productHandler.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", middleware.Authorize())
	{
		orderRoutes.POST("/product/:productId/:userId/:quantity", orderHandler.OrderProduct)
	}
	return r
}
