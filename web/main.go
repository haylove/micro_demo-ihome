package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "helloWorld",
		})
	})
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
