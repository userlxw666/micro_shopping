package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"micro_shopping/app/user/dao"
	"micro_shopping/app/user/server"
	"micro_shopping/config"
	"micro_shopping/idl/pb"
	"net"
	"strconv"
)

func main() {
	config.ReadConfig()
	dao.InitSQL()
	ipport := fmt.Sprintf("%s:%s", config.RdConfigFile.UserService.Host, strconv.Itoa(config.RdConfigFile.UserService.Port))

	consulConfig := api.DefaultConfig()

	conclient, err_conul := api.NewClient(consulConfig)
	if err_conul != nil {
		fmt.Println("创建consul失败", err_conul)
		return
	}

	reg := api.AgentServiceRegistration{
		Tags:    []string{"UserService"},
		Name:    "UserService",
		Address: config.RdConfigFile.UserService.Host,
		Port:    config.RdConfigFile.UserService.Port,
	}

	err_agent := conclient.Agent().ServiceRegister(&reg)
	if err_agent != nil {
		fmt.Println("注册失败", err_agent)
		return
	}

	grpcService := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcService, new(server.UserService))

	lis, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("监听失败", err)
		return
	}
	defer lis.Close()
	fmt.Println("服务启动成功")
	err = grpcService.Serve(lis)
	if err != nil {
		fmt.Println("服务器启动报错", err)
		return
	}
}
