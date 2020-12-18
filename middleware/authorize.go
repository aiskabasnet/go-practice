package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-practice/handler"
	"fmt"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if(authHeader == ""){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"}
		)
		}
		tokenString := authHeader[len(BearerSchema):]
		if token, err := handler.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		}else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)

			} else {
				if token.Valid {
					ctx.Set("userID", claims["userID"])
					fmt.Println("during authorization", claims["userID"])
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}

			}
		}

	}
}
