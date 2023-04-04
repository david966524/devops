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
	r.Use(cors.New(config))

	//router
	r.POST("/login", service.LoginUser)
	router.CfRouter(r)
	router.UserRouter(r)
	router.AsiacloudRouter(r)
	router.AirborneRouter(r)
	router.ImRouter(r)

	logger := myutils.GetLogger()
	logger.Info("this is  日志")

	return r
}

func main() {
	myutils.InitLogger()
	startServer().Run()

}
