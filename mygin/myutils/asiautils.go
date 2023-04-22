package myutils

import (
	"fmt"
	"log"
	"mygin/model"
	"strconv"
	"time"
)

// 获取 亚洲云海token
func GetAsiaCloudToken() string {
	var accunot model.AsiaCloudAccount
	result := DBClinet.Where("id=?", 1).Find(&accunot)
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
	resp, err := HttpClinet().R().
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
