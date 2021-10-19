/*
 *@time       2021/10/15 9:35
 *@version    1.0.0
 *@author     11726
 */

package code

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"web/proto/imgCode"
	"web/response"
	"web/utils"
)

var (
	imgServer imgCode.ImgCodeService
)

func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	call, err := imgServer.Get(context.TODO(), &imgCode.Request{Uuid: uuid})
	fmt.Println(err)
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

func init() {
	client := utils.NewClient()
	imgServer = imgCode.NewImgCodeService("go.micro.service.code", client.Client())
}

func checkImgCode(key, code string) (correct bool, err error) {
	checkResponse, err := imgServer.Check(context.TODO(), &imgCode.CheckRequest{ImgCode: code, Uuid: key})
	if err != nil {
		return false, err
	}
	return checkResponse.IsCorrect, nil
}
