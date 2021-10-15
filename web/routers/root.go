/*
 *@time       2021/10/15 9:29
 *@version    1.0.0
 *@author     11726
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"web/controllers/code"
	"web/controllers/user"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.Static("/home", "static")

	api := router.Group("/api/v1.0")
	{
		api.GET("/session", user.GetSession)
		api.GET("/imagecode/:uuid", code.GetImageCode)
		api.GET("/smscode/:phone", code.GetSmsCode)
	}

	return router
}
