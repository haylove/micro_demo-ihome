/*
 *@time       2021/10/15 9:08
 *@version    1.0.0
 *@author     11726
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"helloworld/web/response"
	"helloworld/web/utils"
)

func GetSession(ctx *gin.Context) {
	response.Err(ctx, http.StatusOK, utils.RECODE_SESSIONERR)
}
