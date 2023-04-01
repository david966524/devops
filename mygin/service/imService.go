package service

import (
	"fmt"
	"log"
	"mygin/cfutils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Im struct {
	gorm.Model
	ProjectName string `json:"projectName";gorm:"type:varchar(50);not null;comment:项目名"`
	ServerIp    string `json:"serverip";gorm:"type:varchar(50);not null;comment:服务器ip"`
	Groupid     string `json:"groupid";gorm:"type:varchar(50);comment:cdn对应组id"`
	Jsonconfig  string `json:"jsonconfig";gorm:"type:varchar(255);comment:json配置文件地址"`
	Remark      string `json:"remark";gorm:"type:varchar(50);comment:备注"`
}

func AddIm(c *gin.Context) {
	var im Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
	}
	db, err := cfutils.ConnectMysqlByDatabaseSql()
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
	var im Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := cfutils.ConnectMysqlByDatabaseSql()
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
	var ims []Im
	db, err := cfutils.ConnectMysqlByDatabaseSql()
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
	var im Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := cfutils.ConnectMysqlByDatabaseSql()
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
