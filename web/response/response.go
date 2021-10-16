/*
 *@time       2021/10/15 8:59
 *@version    1.0.0
 *@author     11726
 */

package response

import (
	"github.com/gin-gonic/gin"
	"web/utils"
)

func Err(ctx *gin.Context, statusCode int, ErrCode utils.ErrCode) {
	ctx.JSON(statusCode,
		gin.H{
			"errno":  ErrCode,
			"errmsg": utils.RecodeText(ErrCode),
		},
	)
}

func Normal(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode,
		gin.H{
			"data": data,
		},
	)
}
