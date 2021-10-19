/*
 *@time       2021/10/20 2:52
 *@version    1.0.0
 *@author     11726
 */

package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Encryption(str string) string {
	//password 加密
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
