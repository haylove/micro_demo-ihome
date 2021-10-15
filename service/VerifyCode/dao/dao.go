/*
 *@time       2021/10/15 19:16
 *@version    1.0.0
 *@author     11726
 */
// 数据库访问对象
package dao

import "github.com/gomodule/redigo/redis"

func StoreCode(key, Code string) error {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Do("setex", key, 60*5, Code)
	if err != nil {
		return err
	}
	return nil
}

func CheckImgCode(key, imgCode string) (isExist bool, err error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	code, err := redis.String(conn.Do("get", key))
	return code == imgCode, nil
}
