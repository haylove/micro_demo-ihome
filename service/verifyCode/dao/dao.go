/*
 *@time       2021/10/15 19:16
 *@version    1.0.0
 *@author     11726
 */
// 数据库访问对象
package dao

import (
	"strings"

	"github.com/gomodule/redigo/redis"
)

var redisPool = redis.Pool{
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", "127.0.0.1:6379")
	},
	MaxIdle:         20,
	MaxActive:       50,
	IdleTimeout:     60,
	MaxConnLifetime: 60 * 5,
}

func StoreCode(key, Code string) error {
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", key, 60*5, Code)
	if err != nil {
		return err
	}
	return nil
}

func CheckCode(key, Code string) (isExist bool, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	code, err := redis.String(conn.Do("get", key))
	return strings.ToLower(code) == strings.ToLower(Code), nil
}
