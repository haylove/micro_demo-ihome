/*
 *@time       2021/10/16 5:34
 *@version    1.0.0
 *@author     11726
 */

package code

import (
	"context"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"

	smscode "web/proto/smsCode"
	"web/response"
	"web/utils"
)

var (
	smsService   smscode.SmsCodeService
	phoneCompile *regexp.Regexp
)

func init() {
	phoneCompile = regexp.MustCompile(`^1[3-9]\d{9}$`)
	client := utils.NewClient()
	smsService = smscode.NewSmsCodeService("go.micro.service.code", client.Client())
}

func GetSmsCode(ctx *gin.Context) {
	phone := ctx.Param("phone")

	if len(phoneCompile.FindStringSubmatch(phone)) == 0 {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_MOBILEERR)
		return
	}

	verifyCode := ctx.Query("text")

	imgUuid := ctx.Query("id")
	correct, err := checkImgCode(imgUuid, verifyCode)
	//fmt.Println(correct)
	if err != nil || !correct {
		response.Err(ctx, http.StatusBadRequest, utils.RECODE_PARAMERR)
		return
	}

	_, _ = smsService.SendSms(context.TODO(), &smscode.SmsRequest{Phone: phone})

	response.Normal(ctx, http.StatusOK, gin.H{"errno": "0000", "errmsg": "发送成功"})
}

func CheckSmsCode(key, code string) (correct bool, err error) {
	checkResponse, err := smsService.Check(context.TODO(), &smscode.CheckRequest{SmsCode: code, Phone: key})
	if err != nil {
		return false, err
	}
	return checkResponse.IsCorrect, nil
}
