package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func AsiacloudRouter(r *gin.Engine) {
	asiacloudRouter := r.Group("/asiacloud", middleware.JWTAuthMiddleware()) //, middleware.JWTAuthMiddleware()
	{
		asiacloudRouter.GET("/vhost", service.GetVhost)
		asiacloudRouter.GET("/vhost/:vhost", service.GetVhostDomainlist)
		asiacloudRouter.POST("/vhost/:vhost", service.AddDomain)
		// asiacloudRouter.PUT("", service.UpdateIm)
		asiacloudRouter.DELETE("/vhost/:vhost/:id", service.DeleteAsiacloudDomain)
		// asiacloudRouter.POST("/line", service.GetLins)
	}
}
