package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `form:"username";gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名"`
	PassWord string `form:"password";gorm:"type:varchar(255);not null;comment:密码"`
}
