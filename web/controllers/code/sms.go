/*
 *@time       2021/10/16 5:34
 *@version    1.0.0
 *@author     11726
 */

package code

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"net/http"

	"github.com/gin-gonic/gin"

	smscode "web/proto/smsCode"
	"web/response"
	"web/utils"
)

func GetSmsCode(ctx *gin.Context) {
	phone := ctx.Param("phone")
	if len(phoneCompile.FindString(phone)) == 0 {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_MOBILEERR)
		return
	}

	verifyCode := ctx.Query("text")

	imgUuid := ctx.Query("id")
	correct, err := checkImgCode(imgUuid, verifyCode)
	if err != nil || !correct {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_PARAMERR)
		return
	}

	server := getSmsServer()
	_, err = server.SendSms(context.TODO(), &smscode.SmsRequest{Phone: phone})
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_UNKNOWERR)
		return
	}
}

func getSmsServer() smscode.SmsCodeService {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	client := micro.NewService(
		micro.Registry(newRegistry),
	)

	service := smscode.NewSmsCodeService("go.micro.service.code", client.Client())
	return service
}
