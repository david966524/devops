package main

import (
	"mygin/myutils"
	"mygin/router"
	"mygin/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startServer() *gin.Engine {
	r := gin.Default()

	// 跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config), myutils.LoggerToFile())

	//router
	r.POST("/login", service.LoginUser)
	router.CfRouter(r)
	router.UserRouter(r)
	router.AsiacloudRouter(r)
	router.AirborneRouter(r)
	router.ImRouter(r)
	router.AliyunRouter(r)
	router.CheckDomainRouter(r)

	return r
}

func main() {
	myutils.Mylog.Info("gin runing........")
	startServer().Run(":8089")
}
