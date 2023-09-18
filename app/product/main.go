package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	ProductDao "micro_shopping/app/product/dao"
	"micro_shopping/app/product/server"
	"micro_shopping/config"
	"micro_shopping/idl/pb"
	"net"
	"strconv"
)

func main() {
	config.ReadConfig()
	ProductDao.InitSQL()
	ipport := fmt.Sprintf("%s:%s", config.RdConfigFile.ProductService.Host, strconv.Itoa(config.RdConfigFile.ProductService.Port))
	consulConfig := api.DefaultConfig()

	consulClient, err_con := api.NewClient(consulConfig)
	if err_con != nil {
		fmt.Println("创建consul客户端失败", err_con)
		return
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"ProductService"},
		Name:    "ProductService",
		Address: config.RdConfigFile.ProductService.Host,
		Port:    config.RdConfigFile.ProductService.Port,
	}

	err_reg := consulClient.Agent().ServiceRegister(&reg)
	if err_reg != nil {
		fmt.Println("consul注册失败")
		return
	}

	grpcService := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcService, new(server.ProductService))

	lis, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer lis.Close()
	fmt.Println("服务器启动成功!")
	err = grpcService.Serve(lis)
	if err != nil {
		fmt.Println("服务器启动报错", err)
		return
	}
}
