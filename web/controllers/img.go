/*
 *@time       2021/10/15 9:35
 *@version    1.0.0
 *@author     11726
 */

package controllers

import (
	"context"
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"google.golang.org/protobuf/types/known/emptypb"

	"helloworld/web/proto/getImgCode"
	"helloworld/web/response"
	"helloworld/web/utils"
)

func GetImageCode(ctx *gin.Context) {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	client := micro.NewService(
		micro.Registry(newRegistry),
	)

	service := getImgCode.NewGetImgCodeService("go.micro.service.getImgCode", client.Client())

	call, err := service.Call(context.TODO(), &emptypb.Empty{})
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

	png.Encode(ctx.Writer, img)
}
