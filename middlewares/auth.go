package middlewares

import (
	"net/http"

	"example.com/booking/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") //gets the token from the header for validation

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unnauthorized action"}) //no further status messages will be sent when this one is active
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unnauthorized action"})
		return
	}

	context.Set("userId", userId)

	context.Next() //next handler in line will execute correctly
}
