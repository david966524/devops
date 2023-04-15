package router

import (
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func AliyunRouter(r *gin.Engine) {
	aliyunRouter := r.Group("/aliyun") //, middleware.JWTAuthMiddleware()
	{
		aliyunRouter.GET("", service.Getecslist)

	}
}
