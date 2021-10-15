/*
 *@time       2021/10/15 7:57
 *@version    1.0.0
 *@author     11726
 */

package utils

type ErrCode string

const (
	RECODE_OK        ErrCode = "0"
	RECODE_DBERR     ErrCode = "4001"
	RECODE_NODATA    ErrCode = "4002"
	RECODE_DATAEXIST ErrCode = "4003"
	RECODE_DATAERR   ErrCode = "4004"

	RECODE_SESSIONERR ErrCode = "4101"
	RECODE_LOGINERR   ErrCode = "4102"
	RECODE_PARAMERR   ErrCode = "4103"
	RECODE_USERONERR  ErrCode = "4104"
	RECODE_ROLEERR    ErrCode = "4105"
	RECODE_PWDERR     ErrCode = "4106"
	RECODE_USERERR    ErrCode = "4107"
	RECODE_SMSERR     ErrCode = "4108"
	RECODE_MOBILEERR  ErrCode = "4109"

	RECODE_REQERR    ErrCode = "4201"
	RECODE_IPERR     ErrCode = "4202"
	RECODE_THIRDERR  ErrCode = "4301"
	RECODE_IOERR     ErrCode = "4302"
	RECODE_SERVERERR ErrCode = "4500"
	RECODE_UNKNOWERR ErrCode = "4501"
)

var recodeText = map[ErrCode]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_USERONERR:  "用户已经注册",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_SMSERR:     "短信失败",
	RECODE_MOBILEERR:  "手机号错误",
}

func RecodeText(code ErrCode) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
