/*
 *@time       2021/10/15 9:08
 *@version    1.0.0
 *@author     11726
 */

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"web/response"
	"web/utils"
)

func GetSession(ctx *gin.Context) {
	response.Err(ctx, http.StatusOK, utils.RECODE_SESSIONERR)
}
