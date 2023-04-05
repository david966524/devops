package service

import (
	"fmt"
	"log"
	"mygin/model"
	"mygin/mycloudflare"
	"net/http"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

func GetCloudFlare(c *gin.Context) {
	fmt.Println("/cloudflare")
	allDomainSlice := mycloudflare.GetAllDomain()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "get success",
		"data": allDomainSlice,
	})

}

func QueryCloudFlare(c *gin.Context) {
	domainName := c.Param("domainname")
	fmt.Println(domainName)
	MyzoneSlice := mycloudflare.QueryDomain(domainName)
	c.JSON(http.StatusOK, gin.H{
		"data": MyzoneSlice,
	})

}

// 定义了一个名为CloudflareRequest的结构体，用于存储要接收的POST参数。然后，我们使用c.BindJSON()方法将POST请求中的JSON参数绑定到该结构体上，并将其存储在一个名为req的变量中。

// 需要注意的是，在结构体的每个字段上，我们使用了一个带有"json"标签的注释。这样可以告诉Gin框架使用POST参数的值来填充每个字段。例如，如果POST参数中存在名为"domainName"的参数，它的值将被填充到CloudflareRequest结构体的DomainName字段中。

// 另外，如果您想手动解析POST参数而不是使用结构体绑定功能，可以使用context对象的PostForm()方法来获取单个POST参数，或者使用context对象的Request.PostForm属性来获取所有的POST参数。具体实现方式与之前的回答相同，这里不再赘述。

func AddCloudflare(c *gin.Context) {

	//fmt.Println(result)
	var req model.CloudflareRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	fmt.Printf("Domain name: %s\n", req.DomainName)
	myzoneList := mycloudflare.AddDomain(req.DomainName)
	log.Println(myzoneList)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "add success",
		"data": myzoneList,
	})
}

func DeleteCloudflare(c *gin.Context) {
	var req model.CloudflareRequest
	c.ShouldBind(&req)
	log.Println("需要删除的domain ID:" + req.DomainId)
	result := mycloudflare.DeleteDomain(req.DomainId)
	if result {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  result,
		})
	}
}

func GetDomainDns(c *gin.Context) {
	domainName := c.Param("domainname")
	log.Println(domainName)
	dnsRecord := mycloudflare.GetDNSRecord(domainName)
	fmt.Println(dnsRecord)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "qurey success",
		"data": dnsRecord,
	})
}

func AddDnsRecord(c *gin.Context) {
	// 添加 DNS 解析记录
	// dnsRecord := cloudflare.CreateDNSRecordParams{
	// 	Type:    "A",
	// 	Name:    "example.com",
	// 	Content: "1.2.3.4",
	// 	TTL:     120,
	// }
	domainId := c.Param("domainId")
	var dnsRecordParams cloudflare.CreateDNSRecordParams
	err := c.ShouldBindJSON(&dnsRecordParams)
	log.Println("id: " + domainId)
	log.Println(err)
	log.Println(dnsRecordParams.Proxied)
	result, msg := mycloudflare.AddDnsRecord(domainId, dnsRecordParams)
	if result {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 666,
		"msg":  msg,
	})
}
func DeleteDnsRecord(c *gin.Context) {
	//domainId := c.Param("domainId")
	//fmt.Println(domainId)
	var req model.CloudflareRequest
	c.ShouldBind(&req)
	log.Println("要删除的domainid :" + req.DomainId)
	log.Println("要删除的dnsid :" + req.DnsRecordId)
	result, msg := mycloudflare.DeleteDnsRecord(req.DomainId, req.DnsRecordId)
	if result {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  msg,
		})
	}
}
