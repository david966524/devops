package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLins(c *gin.Context) {
	var im model.Im
	err := c.ShouldBind(&im)
	if err != nil {
		log.Println(err.Error())
		return
	}
	linelist := ReqLines(im.Jsonconfig)
	if linelist != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Success",
			"data": linelist,
		})
	}
}

func ReqLines(url string) []model.Line {
	client := myutils.HttpCline()
	var datas model.Datas
	resp, err := client.R().SetResult(&datas).EnableTrace().Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//b1, _ := json.MarshalIndent(datas.Data.Lines, "", "    ")
	//fmt.Println(string(b1))
	return datas.Data.Lines
}
