package handler

import (
	"context"
	"encoding/json"
	"getImgCode/dao"
	"getImgCode/proto/getImgCode"
	"getImgCode/utils"
	"github.com/afocus/captcha"
)

type GetImgCode struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImgCode) Call(ctx context.Context, req *getImgCode.Request, rsp *getImgCode.Response) error {
	c := captcha.New()
	// 设置字体
	err := c.SetFont("handler/static/font/comic.ttf")
	if err != nil {
		return err
	}

	ImgCodeStr := utils.RandStr(6)
	//TODO should log db error
	go func() {
		_ = dao.StoreImgCode(req.Uuid, ImgCodeStr)
	}()

	img := c.CreateCustom(ImgCodeStr)

	stream, _ := json.Marshal(img)

	rsp.Img = stream
	return nil
}
