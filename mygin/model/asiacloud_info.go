package model

import (
	"gorm.io/gorm"
)

type VhostResult struct {
	ResponseBody
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

type AsiaCloudAccount struct {
	gorm.Model
	Uid    string `form:"uid";gorm:"type:varchar(50);uniqueIndex;not null;comment:uid"`
	Apikey string `form:"apikey";gorm:"type:varchar(50);not null;comment:密钥"`
}

// 亚洲云海
type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LoginResult struct {
	ResponseBody
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type LoginPayload struct {
	Sign string `json:"sign"`
	Time int64  `json:"time"`
	Uid  string `json:"uid"`
}
