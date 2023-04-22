package service

import (
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get 空降
func GetAirBorneList(c *gin.Context) {
	var airBorneList []model.AirBorne
	result := myutils.DBClinet.Find(&airBorneList)
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

// 添加空降
func AddAirBorne(c *gin.Context) {
	var airBorne model.AirBorne
	err := c.ShouldBind(&airBorne)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result := myutils.DBClinet.Create(&airBorne)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "成功 添加站点 ",
	})
}

// update 空降
func UpdateAirBorne(c *gin.Context) {
	var ab model.AirBorne
	err := c.ShouldBind(&ab)
	if err != nil {
		log.Println(err.Error())
		return
	}

	result := myutils.DBClinet.Model(&ab).Where("id=?", ab.ID).Save(&ab)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
}

// delete 空降
func DeleteAirBorne(c *gin.Context) {
	var ab model.AirBorne
	err := c.ShouldBind(&ab)
	log.Printf("%v --------------------------", &ab)
	if err != nil {
		log.Println(err.Error())
		return
	}
	result := myutils.DBClinet.Delete(&ab)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}

}
