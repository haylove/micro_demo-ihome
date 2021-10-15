/*
 *@time       2021/10/15 9:29
 *@version    1.0.0
 *@author     11726
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"helloworld/web/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.Static("/home", "static")

	api := router.Group("/api/v1.0")
	{
		api.GET("/session", controllers.GetSession)
		api.GET("/imagecode/:uuid", controllers.GetImageCode)

	}

	return router
}