package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 检测 域名 ipc
func CheckDomain(c *gin.Context) {
	url := c.Param("url")
	checkDomainInfo := reqFreeApi(url)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": checkDomainInfo})
}

// 使用httpclient 请求查询接口
func reqFreeApi(url string) *model.CheckDomainInfo {
	var respinfo model.CheckDomainInfo
	requrl := fmt.Sprintf("https://v.api.aa1.cn/api/api-lj-gf/index.php?url=%v", url)
	client := myutils.HttpClinet()
	resp, err := client.R().SetResult(&respinfo).EnableTrace().Get(requrl)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
	return &respinfo
}
