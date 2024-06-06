package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/utils"
)

func Authenticate(ctx *gin.Context) {
	tokenWithBearer := ctx.Request.Header.Get("Authorization")

	if tokenWithBearer == ""{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":"Not authorized",
		})
		return
	}

	token := strings.TrimPrefix(tokenWithBearer, "Bearer ")
	uId, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":"Not authorized",
		})
		return
	}

	ctx.Set("userId", uId)
	ctx.Next()
}
