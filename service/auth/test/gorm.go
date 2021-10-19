/*
 *@time       2021/10/19 21:22
 *@version    1.0.0
 *@author     11726
 */

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println(err)
	}
	s, err := db.DB()
	if err != nil {
		fmt.Println(s, err)
	}
}
