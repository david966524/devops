package myutils

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {

	fmt.Println("get config")
	v := viper.New()
	v.SetConfigName("dbconfig")
	v.SetConfigType("toml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("配置文件未找到！%v\n", err)

		} else {
			fmt.Printf("找到配置文件,但是解析错误,%v\n", err)

		}
	}
	return v
}
