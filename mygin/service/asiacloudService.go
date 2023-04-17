package service

import (
	"encoding/json"
	"fmt"
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//////////////////////////////////////////////////////////

func GetVhost(c *gin.Context) {
	client := myutils.HttpClinet()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	var vhostResult model.VhostResult
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
	var domain model.DomainDetails
	log.Println("请求的 vhost :" + vhost)
	err := c.ShouldBind(&domain)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(domain)
	client := myutils.HttpClinet()
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
	var reqbody model.ResponseBody
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

func reqDomainlist(vhost string) *model.DomainResult {
	client := myutils.HttpClinet()
	asiaCloudToken := myutils.GetAsiaCloudToken()
	var domainResult model.DomainResult
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

func reqDeleteDomain(id string, vhost string) *model.DomainResult {
	deletedomaintUrl := fmt.Sprintf("https://cdnportal.myasiacloud.com/api/site/%v/domain/%v", vhost, id)
	client := myutils.HttpClinet()
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
	var domainResult model.DomainResult
	jsonerr := json.Unmarshal(resq.Body(), &domainResult)
	if jsonerr != nil {
		log.Println(jsonerr.Error())
	}
	return &domainResult
}

func UpdateDomain(c *gin.Context) {
	var domainDetails model.DomainDetails
	c.ShouldBind(&domainDetails)
	log.Println(domainDetails)
}
