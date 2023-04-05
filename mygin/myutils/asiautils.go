package myutils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

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

func GetAsiaCloudToken() string {
	db, err := ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	var accunot AsiaCloudAccount
	result := db.Where("id=?", 1).Find(&accunot)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	log.Println("asiacloud :" + accunot.Uid + "   ----------" + accunot.Apikey)
	apiskeyMd5 := MD5(accunot.Uid + accunot.Apikey)
	timestamp := time.Now().Unix()
	fullapiMd5 := MD5(apiskeyMd5 + strconv.Itoa(int(timestamp)))
	fmt.Println("当前时间戳  ", timestamp)
	fmt.Println(fullapiMd5)
	loginUrl := "https://cdnportal.myasiacloud.com/api/user/login/token"
	var myresule LoginResult
	mypayload := LoginPayload{
		Sign: fullapiMd5,
		Time: timestamp,
		Uid:  accunot.Uid,
	}
	fmt.Println(mypayload)
	resp, err := HttpCline().R().
		SetHeader("Content-Type", "application/json").
		SetResult(&myresule).
		SetBody(&mypayload).
		Post(loginUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(resp.Body()))
	asiaCloudToken := myresule.Data.Token
	return asiaCloudToken
}
