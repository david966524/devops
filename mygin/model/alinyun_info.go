package model

import "gorm.io/gorm"

type AliyunAccess struct {
	gorm.Model
	AccessKeyId     string
	AccessKeySecret string
}
