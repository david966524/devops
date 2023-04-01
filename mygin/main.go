package main

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	//router
	// r.GET("/cloudflare", service.GetCloudFlare)
	// r.GET("/cloudflare/:domainname", service.QueryCloudFlare)
	// r.POST("/cloudflare", service.AddCloudflare)
	// r.DELETE("/cloudflare", service.DeleteCloudflare)
	// r.GET("/cloudflare/dns/:domainname", service.GetDomainDns)
	// r.POST("/cloudflare/dns/:domainId", service.AddDnsRecord)
	// r.DELETE("/cloudflare/dns/", service.DeleteDnsRecord)

	// r.POST("/user/login", service.LoginUser)
	r.POST("/login", service.LoginUser)
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
	userRouter := r.Group("/system", middleware.JWTAuthMiddleware())
	{
		userRouter.GET("/user", service.GetUserList)
		userRouter.POST("/user", service.AddUser)
		userRouter.PUT("/user", service.ChangePassword)
	}
	airborneRouter := r.Group("/airborne", middleware.JWTAuthMiddleware()) //, middleware.JWTAuthMiddleware()
	{
		airborneRouter.GET("", service.GetAirBorneList)
		airborneRouter.POST("", service.AddAirBorne)
		airborneRouter.PUT("", service.UpdateAirBorne)
		airborneRouter.DELETE("", service.DeleteAirBorne)
	}
	imRouter := r.Group("/im") //, middleware.JWTAuthMiddleware()
	{
		imRouter.GET("", service.GetIm)
		imRouter.POST("", service.AddIm)
		imRouter.PUT("", service.UpdateIm)
		imRouter.DELETE("", service.DeleteIm)
		imRouter.POST("/line", service.GetLins)
	}

	r.Run()
}
