package main

import (
	"log"

	"web/routers"
)

func main() {
	router := routers.InitRouter()

	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}

}
