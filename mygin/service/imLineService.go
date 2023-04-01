package service

import (
	"fmt"
	"log"
	"mygin/cfutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Line struct {
	BaseUrl    string `json:"base_url"`
	ResUrl     string `json:"res_url"`
	SocketIP   string `json:"socket_ip"`
	SocketPort int    `json:"socket_port"`
	Timeout    int    `json:"timeout"`
	Ssl        int    `json:"ssl"`
	Remark     string `json:"remark"`
	Type       int    `json:"type"`
}

type Data struct {
	Lines []Line `json:"lines"`
}

type Datas struct {
	Data Data `json:"data"`
	Ok   bool `json:"ok"`
}

func GetLins(c *gin.Context) {
	var im Im
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

func ReqLines(url string) []Line {
	client := cfutils.HttpCline()
	var datas Datas
	resp, err := client.R().SetResult(&datas).EnableTrace().Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//b1, _ := json.MarshalIndent(datas.Data.Lines, "", "    ")
	//fmt.Println(string(b1))
	return datas.Data.Lines
}
