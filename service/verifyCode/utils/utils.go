/*
 *@time       2021/10/15 11:57
 *@version    1.0.0
 *@author     11726
 */

package utils

import (
	"math/rand"
	"strings"
	"time"
)

const AlphaStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandStr(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(AlphaStr)
	ret := strings.Builder{}
	for i := 0; i < length; i++ {
		hash := r.Int() % l
		ret.WriteByte(AlphaStr[hash])
	}
	return ret.String()
}
