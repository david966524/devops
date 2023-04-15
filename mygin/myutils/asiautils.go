package myutils

import (
	"fmt"
	"log"
	"mygin/model"
	"strconv"
	"time"
)

func GetAsiaCloudToken() string {
	db, err := ConnectMysqlByDatabaseSql()
	if err != nil {
		log.Println(err.Error())
	}
	var accunot model.AsiaCloudAccount
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
	var myresule model.LoginResult
	mypayload := model.LoginPayload{
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
