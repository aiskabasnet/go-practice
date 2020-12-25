package middleware

import (
	"fmt"
	"strings"
	"net/http"
	"go-practice/api/service"
	"github.com/gin-gonic/gin"
)

func Auth(s service.FirebaseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		header := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		}
		idToken := strings.TrimSpace(strings.Replace(header, "Bearer","",1))
		token, err := s.VerifyToken(idToken)
		if err != nil{
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("ID", token.UID)
		c.Next()

	}
}
