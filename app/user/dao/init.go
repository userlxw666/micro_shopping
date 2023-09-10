package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"micro_shopping/config"
)

var DB *gorm.DB

func InitSQL() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.RdConfigFile.Username,
		config.RdConfigFile.Password, config.RdConfigFile.Host, config.RdConfigFile.Port, config.RdConfigFile.DatabaseName)
	fmt.Println(dns)
	var ormLogger logger.Interface
	ormLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger:         ormLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		fmt.Println("mysql 数据库连接失败", err)
		return
	}
	DB = db
	fmt.Println("数据库连接成功!")
}
