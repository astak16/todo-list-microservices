package middleware

import (
	"api-gateway/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		var code int = 200
		if auth == "" {
			code = 400
		} else {
			clamis, err := utils.ParseToken(auth)
			fmt.Println(clamis, "2222")
			if err != nil {
				code = 20001
			} else if time.Now().Unix() > clamis.ExpiresAt {
				code = 20002
			}
		}
		if code != 200 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": "token验证失败",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
