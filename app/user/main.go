package main

import (
	"micro_shopping/app/user/dao"
	"micro_shopping/config"
)

func main() {
	config.ReadConfig()
	dao.InitSQL()
}
