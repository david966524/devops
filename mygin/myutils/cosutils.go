package myutils

import (
	"log"
	"mygin/model"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var SecretId string
var SECRETKEY string

func init() {
	var tencentyun model.TencentAccess
	db, err := ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	resule := db.Where("type=?", 1).Find(&tencentyun)
	if resule.Error != nil {
		log.Println(resule.Error.Error())
	}
	log.Println(tencentyun)
	SecretId = tencentyun.AccessKeyId
	SECRETKEY = tencentyun.AccessKeySecret
}

func CosClient() *cos.Client {
	u, _ := url.Parse("https://newim-txy4-123qwer-1312026995.cos.ap-hongkong.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SecretId,
			SecretKey: SECRETKEY,
		},
	})
	return client
}
