package middleware

import (
	"go-practice/api/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(s service.FirebaseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		}
		idToken := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))
		token, err := s.VerifyToken(idToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		ctx.Set("ID", token.UID)
		ctx.Next()

	}
}
