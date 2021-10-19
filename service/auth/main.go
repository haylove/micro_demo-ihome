package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"auth/handler"
	user "auth/proto/user"
)

func main() {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)

	// Initialise service
	service.Init()

	dsn := "root:9506@tcp(127.0.0.1:3306)/ihome?charset=utf8mb4&parseTime=True&loc=Local"
	newUser, err := handler.NewUser("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// Register Handler
	_ = user.RegisterAuthHandler(service.Server(), newUser)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
