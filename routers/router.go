package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/l077j/JWT/middleware"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	// api.Use(middleware.Cors(), middleware.Auth())
	api.Use(middleware.Auth())
	// 课程相关接口
	InitCourse(api)
	// 用户相关接口
	InitUser(api)

	notAuthApi := r.Group("/api")
	notAuthApi.Use()
	// 登录接口
	InitLogin(notAuthApi)
}
