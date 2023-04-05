package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {

	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	var users []model.User
	result := db.Find(&users)
	fmt.Println(result.RowsAffected)
	log.Println(users)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "get all",
		"data": users,
	})
}

func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户添加成功",
	})

}

func ChangePassword(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("----change passwod" + user.UserName + user.PassWord)
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Table("users").Where("id=?", user.ID).Update("pass_word", user.PassWord)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": 200,
		"msg":  "修改密码成功",
	})
}
