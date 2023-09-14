package rpc

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"micro_shopping/idl/pb"
	"strconv"
)

func GetGrpcConn(ServiceName string, ServiceTags string, consulClient *api.Client) *grpc.ClientConn {
	service, _, service_err := consulClient.Health().Service(ServiceName, ServiceTags, true, nil)
	if service_err != nil {
		fmt.Println("获取健康服务报错", service_err)
		return nil
	}

	s := service[0].Service
	address := s.Address + ":" + strconv.Itoa(s.Port)
	fmt.Printf("address:%v\n", address)
	grpcConn, _ := grpc.Dial(address, grpc.WithInsecure())
	return grpcConn
}

var (
	UserService     pb.UserServiceClient
	CategoryService pb.CategoryServiceClient
)

func InitRPC() {
	consulConfig := api.DefaultConfig()
	consulClient, err_consul := api.NewClient(consulConfig)
	if err_consul != nil {
		fmt.Println("consul创建对象失败", err_consul)
		return
	}
	UserService = pb.NewUserServiceClient(GetGrpcConn("UserService", "UserService", consulClient))
	CategoryService = pb.NewCategoryServiceClient(GetGrpcConn("CategoryService", "CategoryService", consulClient))
}
