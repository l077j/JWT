package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/l077j/JWT/user"
)

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/user", user.Get)
		v1.POST("/user", user.Add)
	}

	v2 := group.Group("/v2")
	{
		v2.GET("user", user.Get)
		v2.POST("user", user.Add)
	}
}
