/*
 *@time       2021/10/15 9:08
 *@version    1.0.0
 *@author     11726
 */

package user

import (
	"github.com/gin-gonic/gin"

	user "web/proto/auth"
	smscode "web/proto/smsCode"
	"web/utils"
)

var DefaultController = initDefault()

func initDefault() *AuthController {

	authService := user.NewAuthService("go.micro.service.user", utils.NewClient().Client())

	smsCodeService := smscode.NewSmsCodeService("go.micro.service.code", utils.NewClient().Client())
	return NewAuthController(authService, smsCodeService)
}

func Register(ctx *gin.Context) {
	DefaultController.Register(ctx)
}

func Login(ctx *gin.Context) {
	DefaultController.Login(ctx)
}
