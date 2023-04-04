package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func ImRouter(r *gin.Engine) {
	imRouter := r.Group("/im", middleware.JWTAuthMiddleware()) //, middleware.JWTAuthMiddleware()
	{
		imRouter.GET("", service.GetIm)
		imRouter.POST("", service.AddIm)
		imRouter.PUT("", service.UpdateIm)
		imRouter.DELETE("", service.DeleteIm)
		imRouter.POST("/line", service.GetLins)
	}
}
