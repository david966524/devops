package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func AliyunRouter(r *gin.Engine) {
	aliyunRouter := r.Group("/aliyun", middleware.JWTAuthMiddleware()) //, middleware.JWTAuthMiddleware()
	{
		aliyunRouter.GET("", service.Getecslist)

	}
}
