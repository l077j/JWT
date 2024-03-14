package routers

import (
	"github.com/l077j/JWT/course"

	"github.com/gin-gonic/gin"
)

func InitCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	{
		v1.GET("/course/:id", course.Get)
		v1.POST("/course", course.Add)
		v1.PUT("/course", course.Update)
		v1.DELETE("/course", course.Delete)
	}
}
