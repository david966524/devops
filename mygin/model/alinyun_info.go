package model

import "gorm.io/gorm"

type AliyunAccess struct {
	gorm.Model
	AccessKeyId     string `json:"accessKeyId";gorm:"type:varchar(50);not null;comment:阿里云apikey"`
	AccessKeySecret string `json:"accessKeySecret";gorm:"type:varchar(50);not null;comment:阿里云api密钥"`
	Type            int    `json:"type";gorm:"type:varchar(50);not null;comment:开启状态 1 开启 0关闭"`
}

type Instance struct {
	IpAddress    string `json:"IpAddress"`
	ImageId      string `json:"ImageId"`
	InstanceId   string `json:"InstanceId"`
	InstanceName string `json:"InstanceName"`
	InstanceType string `json:"InstanceType"`
	KeyPairName  string `json:"KeyPairName"`
	OSName       string `json:"OSName"`
}
