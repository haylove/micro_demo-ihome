package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"google.golang.org/protobuf/types/known/emptypb"

	"getImgCode/proto/getImgCode"
	"getImgCode/utils"
)

type GetImgCode struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImgCode) Call(_ context.Context, _ *emptypb.Empty, rsp *getImgCode.Response) error {
	c := captcha.New()
	// 设置字体
	err := c.SetFont("handler/static/font/comic.ttf")
	if err != nil {
		return err
	}

	ImgCodeStr := utils.RandStr(6)
	img := c.CreateCustom(ImgCodeStr)

	stream, _ := json.Marshal(img)

	rsp.Img = stream
	return nil
}
