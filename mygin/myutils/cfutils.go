package myutils

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"gorm.io/gorm"
)

// cloudflare client
func GetCfApi() *cloudflare.API {
	var cfAccount CloudFlareAccount
	DBClinet.Where("id=?", 1).Find(&cfAccount)
	log.Println("cfAccount_email: ", cfAccount.Email)
	log.Println("cfAccount_key: ", cfAccount.Key)
	api, err := cloudflare.New(cfAccount.Key, cfAccount.Email)
	if err != nil {
		fmt.Println(err)
	}
	return api
}

type CloudFlareAccount struct {
	gorm.Model
	Email string `gorm:"type:varchar(50);not null;comment:邮箱"`
	Key   string `gorm:"type:varchar(50);not null;comment:密钥"`
}
