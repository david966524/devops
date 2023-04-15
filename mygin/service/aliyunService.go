package service

import (
	"log"
	"mygin/model"
	"mygin/myutils"
	"net/http"

	ecs "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
)

func Getecslist(c *gin.Context) {
	ecslist := ResqToMy()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": ecslist,
	})
}

func ResqToMy() []*model.Instance {
	client, err := myutils.CreateClient()
	if err != nil {
		log.Println(err.Error())
	}
	request := &ecs.DescribeInstancesRequest{
		PageSize: tea.Int32(100),
		RegionId: tea.String(string("cn-hongkong")),
	}
	//log.Println(client)
	resp, err := client.DescribeInstances(request)
	instances := resp.Body.Instances.Instance
	if err != nil {
		log.Println(err.Error())
	}

	var instanceList []*model.Instance

	for _, instance := range instances {
		var myinstance = new(model.Instance)
		myinstance.ImageId = *instance.ImageId
		myinstance.InstanceId = *instance.InstanceId

		if instance.InstanceName != nil {
			myinstance.InstanceName = *instance.InstanceName
		}

		if instance.InstanceType != nil {
			myinstance.InstanceType = *instance.InstanceType
		}

		if instance.EipAddress != nil && instance.EipAddress.IpAddress != nil {
			myinstance.IpAddress = *instance.EipAddress.IpAddress
		}

		if instance.KeyPairName != nil {
			myinstance.KeyPairName = *instance.KeyPairName
		}

		if instance.OSName != nil {
			myinstance.OSName = *instance.OSName
		}

		instanceList = append(instanceList, myinstance)
	}
	return instanceList
}
