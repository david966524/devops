package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取 user 列表
func GetUserList(c *gin.Context) {

	var users []model.User
	result := myutils.DBClinet.Find(&users)
	fmt.Println(result.RowsAffected)
	myutils.GetLogger().Info(users)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "get all",
		"data": users,
	})
}

// 添加 user
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := myutils.DBClinet.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户添加成功",
	})
}

// 修改用户 密码
func ChangePassword(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("----change passwod" + user.UserName + user.PassWord)
	result := myutils.DBClinet.Table("users").Where("id=?", user.ID).Update("pass_word", user.PassWord)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": 200,
		"msg":  "修改密码成功",
	})
}
