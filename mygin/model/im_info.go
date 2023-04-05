package model

import "gorm.io/gorm"

type Im struct {
	gorm.Model
	ProjectName string `json:"projectName";gorm:"type:varchar(50);not null;comment:项目名"`
	ServerIp    string `json:"serverip";gorm:"type:varchar(50);not null;comment:服务器ip"`
	Groupid     string `json:"groupid";gorm:"type:varchar(50);comment:cdn对应组id"`
	Jsonconfig  string `json:"jsonconfig";gorm:"type:varchar(255);comment:json配置文件地址"`
	Remark      string `json:"remark";gorm:"type:varchar(50);comment:备注"`
}
