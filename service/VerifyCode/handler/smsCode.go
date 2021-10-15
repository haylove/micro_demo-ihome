/*
 *@time       2021/10/16 4:44
 *@version    1.0.0
 *@author     11726
 */

package handler

import (
	"context"
	"fmt"

	"VerifyCode/dao"
	smscode "VerifyCode/proto/smsCode"
	"VerifyCode/utils"
)

type SmsCode struct{}

func (s SmsCode) SendSms(_ context.Context, request *smscode.SmsRequest, response *smscode.SmsResponse) error {
	//TODO go sendSms

	randStr := utils.RandStr(4)
	fmt.Println(randStr)

	go func() {
		_ = dao.StoreCode("sms_"+request.Phone, randStr)
	}()
	response.IsSend = true
	return nil
}
