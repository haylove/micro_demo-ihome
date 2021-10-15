/*
 *@time       2021/10/15 9:35
 *@version    1.0.0
 *@author     11726
 */

package code

import (
	"context"
	"encoding/json"
	"image/png"
	"net/http"
	"regexp"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"web/proto/imgCode"
	"web/response"
	"web/utils"
)

var phoneCompile *regexp.Regexp

func init() {
	phoneCompile = regexp.MustCompile(`^1[3-9]/d{9}$`)
}

func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	service := getImgServer()

	call, err := service.Get(context.TODO(), &imgCode.Request{Uuid: uuid})
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_UNKNOWERR)
		return
	}

	var img captcha.Image
	err = json.Unmarshal(call.Img, &img)
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, utils.RECODE_UNKNOWERR)
		return
	}

	_ = png.Encode(ctx.Writer, img)
}

func getImgServer() imgCode.ImgCodeService {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	client := micro.NewService(
		micro.Registry(newRegistry),
	)

	service := imgCode.NewImgCodeService("go.micro.service.code", client.Client())
	return service
}

func checkImgCode(key, code string) (correct bool, err error) {
	service := getImgServer()
	checkResponse, err := service.Check(context.TODO(), &imgCode.CheckRequest{ImgCode: code, Uuid: key})
	if err != nil {
		return false, err
	}
	return checkResponse.IsCorrect, nil
}
