package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"house/handler"
	"house/subscriber"

	house "house/proto/house"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.house"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	house.RegisterHouseHandler(service.Server(), new(handler.House))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.house", service.Server(), new(subscriber.House))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
