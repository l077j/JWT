package routers

import "github.com/gin-gonic/gin"

func InitLogin(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	{
		v1.GET("/login", login.Login)
	}
}
