package Controllers

import (
	"fmt"
	"go-practice/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	var articles []Models.Article
	err := Models.GetAllArticles(&articles)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, articles)
		return
	}
}

func CreateArticles(c *gin.Context) {
	var article Models.Article
	fmt.Println("article", article)
	c.BindJSON(&article)
	err := Models.CreateArticle(&article)
	if err != nil {
		fmt.Println("Error", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println("No any errprr", article)
		c.JSON(http.StatusOK, article)
	}
}

func GetArticleByID(c *gin.Context) {
	var article Models.Article
	id := c.Params.ByName("id")
	err := Models.GetArticleById(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, article)
	}
}

func UpdateArticle(c *gin.Context) {
	var article Models.Article
	id := c.Params.ByName("id")
	err := Models.GetArticleById(&article, id)
	if err != nil {
		c.JSON(http.StatusNotFound, article)
	}
	c.BindJSON(&article)
	err = Models.UpdateArticle(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, article)
	}
}

func DeleteArticle(c *gin.Context) {
	var article Models.Article
	id := c.Params.ByName("id")

	err := Models.DeleteArticle(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
