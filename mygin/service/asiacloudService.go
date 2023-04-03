package service

import (
	"encoding/json"
	"fmt"
	"log"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VhostResult struct {
	myutils.ResponseBody
	Data struct {
		Total int `json:"total"`
		List  []struct {
			Name   string `json:"name"`
			Pid    int    `json:"pid"`
			Status int    `json:"status"`
		} `json:"list"`
	} `json:"data"`
}

type DomainResult struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DomainData `json:"data"`
}

type DomainData struct {
	Total int             `json:"total"`
	List  []DomainDetails `json:"list"`
}

type DomainDetails struct {
	ID            int                 `json:"id"`
	Domain        string              `json:"domain"`
	Host          string              `json:"host"`
	Vhost         string              `json:"vhost"`
	Status        int                 `json:"status"`
	Cname         string              `json:"cname"`
	PublicSetting DomainPublicSetting `json:"publicSetting"`
}

type DomainPublicSetting struct {
	Sslcsr        string `json:"sslcsr"`
	Sslkey        string `json:"sslkey"`
	Hash          string `json:"hash"`
	Hsts          int    `json:"hsts"`
	ForceSSL      int    `json:"force_ssl"`
	MaxErrorCount string `json:"max_error_count"`
	ErrorTryTime  string `json:"error_try_time"`
	Portmap       int    `json:"portmap"`
}

//////////////////////////////////////////////////////////

func GetVhost(c *gin.Context) {
	client := myutils.HttpCline()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	var vhostResult VhostResult
	vhostUrl := "https://cdnportal.myasiacloud.com/api/proxy/vhost"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Get(vhostUrl)

	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(string(resp.Body()))
	json.Unmarshal(resp.Body(), &vhostResult)
	log.Println(vhostResult.Data.List)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": vhostResult.Message,
		"data":    vhostResult.Data.List,
	})
}

func GetVhostDomainlist(c *gin.Context) {
	vhost := c.Param("vhost")
	log.Println("请求的 vhost :" + vhost)
	v := reqDomainlist(vhost)
	c.JSON(http.StatusOK, gin.H{
		"code":    v.Code,
		"message": v.Message,
		"data":    v.Data.List,
	})
}

func reqDomainlist(vhost string) *DomainResult {
	client := myutils.HttpCline()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	var domainResult DomainResult
	getdomaintUrl := "https://cdnportal.myasiacloud.com/api/site/" + vhost + "/domain/page"
	resp1, err1 := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Get(getdomaintUrl)

	if err1 != nil {
		fmt.Println(err1.Error())
	}
	json.Unmarshal(resp1.Body(), &domainResult)
	return &domainResult
}
