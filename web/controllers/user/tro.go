/*
 *@time       2021/10/20 17:13
 *@version    1.0.0
 *@author     11726
 */

package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (c *AuthController) Login(ctx *gin.Context) {
	var ReqData struct {
		Username string `json:"username"`
		PassWord string `json:"password"`
	}
	_ = ctx.Bind(&ReqData)
	if ReqData.Username == "" || ReqData.PassWord == "" {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_PARAMERR)
		return
	}
	login, err := c.userService.Login(context.TODO(), &user.LoginReq{
		Username: ReqData.Username,
		Password: ReqData.PassWord,
	})
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_SERVERERR)
		return
	}
	if !login.LoginSuccess {
		response.Err(ctx, http.StatusUnauthorized, utils.RECODE_LOGINERR)
		return
	}
	response.Normal(ctx, http.StatusOK, login)

}
