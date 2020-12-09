package Routes

import (
	"go-practice/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}

	grp2 := r.Group("/article")
	{
		grp2.GET("articles", Controllers.GetArticles)
		grp2.POST("article", Controllers.CreateArticles)
		grp2.GET("article/:id", Controllers.GetArticleByID)
		grp2.PUT("article/:id", Controllers.UpdateArticle)
		grp2.DELETE("article/:id", Controllers.DeleteArticle)

	}
	return r
}
