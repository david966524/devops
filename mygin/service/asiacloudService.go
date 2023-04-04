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

type DeleteDomaindome struct {
	Id    string `json:"id";form:"id"`
	Vhost string `json:"vhost";form:"vhost"`
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

func AddDomain(c *gin.Context) {
	vhost := c.Param("vhost")
	var domain DomainDetails
	log.Println("请求的 vhost :" + vhost)
	err := c.ShouldBind(&domain)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(domain)
	client := myutils.HttpCline()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	addDomainUrl := "https://cdnportal.myasiacloud.com/api/site/" + vhost + "/domain"
	resp, err1 := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		SetBody(&domain).
		Post(addDomainUrl)
	if err1 != nil {
		log.Println(err1.Error())
	}
	var reqbody myutils.ResponseBody
	err2 := json.Unmarshal(resp.Body(), &reqbody)
	log.Println(reqbody)
	if err2 != nil {
		log.Println(err2.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"code": reqbody.Code,
		"msg":  reqbody.Message,
	})
}

func DeleteAsiacloudDomain(c *gin.Context) {
	vhost := c.Param("vhost")
	id := c.Param("id")
	result := reqDeleteDomain(id, vhost)

	c.JSON(http.StatusOK, gin.H{
		"code": result.Code,
		"msg":  result.Message,
	})
}

func reqDomainlist(vhost string) *DomainResult {
	client := myutils.HttpCline()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	var domainResult DomainResult
	getdomaintUrl := "https://cdnportal.myasiacloud.com/api/site/" + vhost + "/domain/page"
	resp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Get(getdomaintUrl)

	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(resp.Body(), &domainResult)
	return &domainResult
}

func reqDeleteDomain(id string, vhost string) *DomainResult {
	deletedomaintUrl := fmt.Sprintf("https://cdnportal.myasiacloud.com/api/site/%v/domain/%v", vhost, id)
	client := myutils.HttpCline()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	resq, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+asiaCloudToken).
		Delete(deletedomaintUrl)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var domainResult DomainResult
	jsonerr := json.Unmarshal(resq.Body(), &domainResult)
	if jsonerr != nil {
		log.Println(jsonerr.Error())
	}
	return &domainResult
}
