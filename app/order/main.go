package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	Orderdao "micro_shopping/app/order/dao"
	"micro_shopping/app/order/server"
	"micro_shopping/config"
	"micro_shopping/idl/pb"
	"net"
	"strconv"
)

func main() {
	config.ReadConfig()
	Orderdao.InitSQL()
	consulConfig := api.DefaultConfig()

	consulClient, err_cli := api.NewClient(consulConfig)
	if err_cli != nil {
		fmt.Println("创建consul失败", err_cli)
		return
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"OrderService"},
		Name:    "OrderService",
		Address: config.RdConfigFile.OrderService.Host,
		Port:    config.RdConfigFile.OrderService.Port,
	}

	err_reg := consulClient.Agent().ServiceRegister(&reg)
	if err_reg != nil {
		fmt.Println("consul注册失败", err_reg)
		return
	}

	grpcService := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcService, new(server.OrderService))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.RdConfigFile.ProductService.Host,
		strconv.Itoa(config.RdConfigFile.ProductService.Port)))
	if err != nil {
		fmt.Println("listen error", err)
		return
	}
	fmt.Println("服务器启动成功")
	err = grpcService.Serve(lis)
	if err != nil {
		fmt.Println("服务器启动报错", err)
		return
	}
}
