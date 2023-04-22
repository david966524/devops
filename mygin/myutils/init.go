package myutils

import (
	"fmt"
	"log"
	"mygin/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DBClinet
var DBClinet *gorm.DB

// aliyun 账号
var aliyunaccess model.AliyunAccess

// 腾讯云
var tencentyun model.TencentAccess

// 定义全局logger
var Mylog *logrus.Logger

// init 初始化 logger
func init() {
	InitLogger()
	Mylog = GetLogger()
}

// init 初始化 DBclinet
func init() {
	fmt.Println("run dbuitls")
	clinet, err := ConnectMysqlByDatabaseSql()
	if err != nil {
		GetLogger().Info(err.Error())
	}
	DBClinet = clinet
}

// init 初始化 阿里云
func init() {
	result := DBClinet.Where("type = ?", 1).First(&aliyunaccess)
	if result.Error != nil {
		GetLogger().Info(result.Error.Error())
	}
	GetLogger().Info(aliyunaccess)
}

// init 初始化 腾讯云
func init() {
	resule := DBClinet.Where("type=?", 1).Find(&tencentyun)
	if resule.Error != nil {
		log.Println(resule.Error.Error())
	}
	log.Println(tencentyun)
}
