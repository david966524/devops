package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		fmt.Println(err.Error())
	}
	result := db.Create(&im)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func UpdateIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	result := db.Model(&im).Where("id=?", im.ID).Save(&im)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func GetIm(c *gin.Context) {
	var ims []model.Im
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	result := db.Find(&ims)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": ims,
	})
}

func DeleteIm(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := myutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Table("ims").Delete(&im)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}
}
