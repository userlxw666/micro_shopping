package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigFile struct {
	Mysql
}

type Mysql struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

var RdConfigFile *ConfigFile

func ReadConfig() {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("read config file error", err)
		return
	}
	err = v.Unmarshal(&RdConfigFile)
	if err != nil {
		fmt.Println("解析到结构体失败", err)
		return
	}
}
