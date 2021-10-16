package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"VerifyCode/handler"
	imgCode "VerifyCode/proto/imgCode"
	smsCode "VerifyCode/proto/smsCode"
)

func main() {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.code"),
		micro.Version("latest"),
		micro.Registry(newRegistry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = imgCode.RegisterImgCodeHandler(service.Server(), new(handler.ImgCode))
	_ = smsCode.RegisterSmsCodeHandler(service.Server(), new(handler.SmsCode))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
