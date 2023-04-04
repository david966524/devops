package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func AirborneRouter(r *gin.Engine) {
	airborneRouter := r.Group("/airborne", middleware.JWTAuthMiddleware()) //, middleware.JWTAuthMiddleware()
	{
		airborneRouter.GET("", service.GetAirBorneList)
		airborneRouter.POST("", service.AddAirBorne)
		airborneRouter.PUT("", service.UpdateAirBorne)
		airborneRouter.DELETE("", service.DeleteAirBorne)
	}
}
