/*
 *@time       2021/10/15 9:08
 *@version    1.0.0
 *@author     11726
 */

package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	user "web/proto/auth"
	smscode "web/proto/smsCode"
	"web/response"
	"web/utils"
)

type AuthController struct {
	userService user.AuthService
	smsService  smscode.SmsCodeService
}

func NewAuthController(userService user.AuthService, smsService smscode.SmsCodeService) *AuthController {
	return &AuthController{userService: userService, smsService: smsService}
}

func initDefault() *AuthController {

	authService := user.NewAuthService("go.micro.service.user", utils.NewClient().Client())

	smsCodeService := smscode.NewSmsCodeService("go.micro.service.code", utils.NewClient().Client())
	return NewAuthController(authService, smsCodeService)
}

var DefaultController = initDefault()

func (c *AuthController) Register(ctx *gin.Context) {

	var ReqData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	_ = ctx.Bind(&ReqData)

	if ReqData.Mobile == "" || ReqData.PassWord == "" || ReqData.SmsCode == "" {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_PARAMERR)
		return
	}

	checkResponse, err := c.smsService.Check(context.TODO(), &smscode.CheckRequest{
		Phone:   ReqData.Mobile,
		SmsCode: ReqData.SmsCode,
	})

	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_SERVERERR)
		return
	}
	if !checkResponse.IsCorrect {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_DATAERR)
		return
	}

	userEntity, err := c.userService.Register(context.TODO(), &user.User{
		Username:    ReqData.Mobile,
		MobilePhone: ReqData.Mobile,
		Password:    ReqData.PassWord,
	})
	fmt.Println("user", err)
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_SERVERERR)
		return
	}
	response.Normal(ctx, http.StatusOK, userEntity)
}

func GetSession(ctx *gin.Context) {
	response.Err(ctx, http.StatusOK, utils.RECODE_SESSIONERR)
}

func Register(ctx *gin.Context) {
	DefaultController.Register(ctx)
}
