package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"micro_shopping/app/category/dao"
	"micro_shopping/app/category/server"
	"micro_shopping/config"
	"micro_shopping/idl/pb"
	"net"
	"strconv"
)

func main() {
	config.ReadConfig()
	CategoryDao.InitSQL()
	ipport := fmt.Sprintf("%s:%s", config.RdConfigFile.CategoryService.Host, strconv.Itoa(config.RdConfigFile.CategoryService.Port))

	consulConfig := api.DefaultConfig()

	consulClient, err_con := api.NewClient(consulConfig)
	if err_con != nil {
		fmt.Println("create consul error", err_con)
		return
	}

	req := api.AgentServiceRegistration{
		Name:    "CategoryService",
		Tags:    []string{"CategoryService"},
		Port:    config.RdConfigFile.CategoryService.Port,
		Address: config.RdConfigFile.CategoryService.Host,
	}

	err_reg := consulClient.Agent().ServiceRegister(&req)
	if err_reg != nil {
		fmt.Println("consul register error", err_reg)
		return
	}

	grpcService := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpcService, new(server.CategoryService))

	lis, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	defer lis.Close()
	fmt.Println("服务器启动成功")
	err = grpcService.Serve(lis)
	if err != nil {
		fmt.Println("服务器启动报错", err)
		return
	}
}
