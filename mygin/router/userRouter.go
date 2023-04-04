package router

import (
	"mygin/middleware"
	"mygin/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("/system", middleware.JWTAuthMiddleware())
	{
		userRouter.GET("/user", service.GetUserList)
		userRouter.POST("/user", service.AddUser)
		userRouter.PUT("/user", service.ChangePassword)
	}
}
