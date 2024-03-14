package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l077j/JWT/jwt"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		accessToken := context.Request.Header.Get("access_token")
		data := &jwt.Data{}
		err := jwt.Verify(accessToken, data)
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{
				"message": "身份认证失败",
			})
			context.Abort()
		}
		context.Set("authInfo", data)
		context.Next()
	}
}
