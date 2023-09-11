package main

import (
	"micro_shopping/app/gateway/router"
	"micro_shopping/app/gateway/rpc"
)

func main() {
	rpc.InitRPC()
	router.NewRouter()
}
