package ProductDao

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"micro_shopping/config"
)

var DB *gorm.DB

func InitSQL() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.RdConfigFile.Mysql.Username,
		config.RdConfigFile.Mysql.Password, config.RdConfigFile.Mysql.Host, config.RdConfigFile.Mysql.Port, config.RdConfigFile.Mysql.DatabaseName)
	fmt.Println(dns)
	var ormLogger logger.Interface
	ormLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	DB = db
	Migration()
}

func NewSqlClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}
