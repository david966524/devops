package service

import (
	"log"
	"mygin/model"
	"mygin/myjwt"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login 校验  发送token
func LoginUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
	}

	var myuser model.User
	if result := myutils.DBClinet.Where("user_name=?", user.UserName).First(&myuser); result.Error != nil {
		log.Println(myuser.PassWord)
		log.Println(err)
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "username or password is incorrect"})
		return
	}
	if user.PassWord != myuser.PassWord {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "username or password is incorrect"})
		return
	}
	token, err := myjwt.GenToken(user.UserName)
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "login success",
		"data": gin.H{
			"token":    token,
			"username": user.UserName,
		},
	})
	log.Println(user)

}
