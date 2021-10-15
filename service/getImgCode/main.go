package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"getImgCode/handler"
	getImgCode "getImgCode/proto/getImgCode"
)

func main() {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.getImgCode"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = getImgCode.RegisterGetImgCodeHandler(service.Server(), new(handler.GetImgCode))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
