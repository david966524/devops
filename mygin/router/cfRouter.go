package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func CfRouter(r *gin.Engine) {
	cfRouter := r.Group("/cloudflare", middleware.JWTAuthMiddleware())
	{
		cfRouter.GET("", service.GetCloudFlare)
		cfRouter.GET("/:domainname", service.QueryCloudFlare)
		cfRouter.POST("", service.AddCloudflare)
		cfRouter.DELETE("", service.DeleteCloudflare)
		cfRouter.GET("/dns/:domainname", service.GetDomainDns)
		cfRouter.POST("/dns/:domainId", service.AddDnsRecord)
		cfRouter.DELETE("/dns/", service.DeleteDnsRecord)
	}
}
