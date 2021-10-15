package handler

import (
	"context"
	"encoding/json"

	"github.com/afocus/captcha"

	"VerifyCode/dao"
	"VerifyCode/proto/imgCode"
	"VerifyCode/utils"
)

type ImgCode struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ImgCode) Get(_ context.Context, req *imgCode.Request, rsp *imgCode.Response) error {
	c := captcha.New()
	// 设置字体
	err := c.SetFont("conf/font/comic.ttf")
	if err != nil {
		return err
	}

	ImgCodeStr := utils.RandStr(6)
	//TODO should log db error
	go func() {
		_ = dao.StoreCode(req.Uuid, ImgCodeStr)
	}()

	img := c.CreateCustom(ImgCodeStr)

	stream, _ := json.Marshal(img)

	rsp.Img = stream
	return nil
}

func (e *ImgCode) Check(ctx context.Context, request *imgCode.CheckRequest, response *imgCode.CheckResponse) error {
	isExist, err := dao.CheckImgCode(request.Uuid, request.ImgCode)
	if err != nil {
		return err
	}

	if isExist {
		response.IsCorrect = true
	}
	return nil
}
