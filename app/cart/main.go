package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"micro_shopping/app/cart/dao"
	"micro_shopping/app/cart/server"
	"micro_shopping/config"
	"micro_shopping/idl/pb"
	"net"
	"strconv"
)

func main() {
	config.ReadConfig()
	Cartdao.InitSQL()
	consulConfig := api.DefaultConfig()

	consulClient, err_cli := api.NewClient(consulConfig)
	if err_cli != nil {
		fmt.Println("创建consul失败", err_cli)
		return
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"CartService"},
		Name:    "CartService",
		Address: config.RdConfigFile.CartService.Host,
		Port:    config.RdConfigFile.CartService.Port,
	}

	err_reg := consulClient.Agent().ServiceRegister(&reg)
	if err_reg != nil {
		fmt.Println("consul注册失败", err_reg)
		return
	}

	grpcService := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcService, new(server.CartService))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.RdConfigFile.CartService.Host,
		strconv.Itoa(config.RdConfigFile.CartService.Port)))
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
