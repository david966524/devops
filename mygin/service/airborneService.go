package service

import (
	"fmt"
	"log"
	"mygin/cfutils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AirBorne struct {
	gorm.Model
	ProjectName string `json:"projectName";form:"projectName";gorm:"type:varchar(50);not null;comment:项目名"`
	ServerIp    string `json:"serverip";form:"serverip";gorm:"type:varchar(50);not null;comment:服务器ip"`
	Groupid     string `json:"groupid";form:"groupid";gorm:"type:varchar(50);comment:方能对应组id"`
	Remark      string `json:"remark";form:"remark";gorm:"type:varchar(50);comment:备注"`
}

func GetAirBorneList(c *gin.Context) {

	var airBorneList []AirBorne
	db, err := cfutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Find(&airBorneList)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": airBorneList,
	})
}

func AddAirBorne(c *gin.Context) {
	var airBorne AirBorne
	err := c.ShouldBind(&airBorne)
	if err != nil {
		log.Println(err.Error())
		return
	}

	db, err := cfutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Create(&airBorne)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功 添加站点 ",
	})
}

func UpdateAirBorne(c *gin.Context) {
	var ab AirBorne
	err := c.ShouldBind(&ab)
	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := cfutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := db.Model(&ab).Where("id=?", ab.ID).Save(&ab)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

func DeleteAirBorne(c *gin.Context) {
	var ab AirBorne
	err := c.ShouldBind(&ab)
	log.Printf("%v --------------------------", &ab)

	if err != nil {
		log.Println(err.Error())
		return
	}
	db, err := cfutils.ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(db)
	result := db.Delete(&ab)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}

}
