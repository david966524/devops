package router

import (
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func CheckDomainRouter(r *gin.Engine) {
	checkDomainRouter := r.Group("/check") //, middleware.JWTAuthMiddleware()
	{
		checkDomainRouter.GET("/:url", service.CheckDomain)
	}
}
