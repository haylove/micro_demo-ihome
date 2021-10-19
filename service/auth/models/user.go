/*
 *@time       2021/10/19 20:47
 *@version    1.0.0
 *@author     11726
 */

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model          //id 编号
	Name         string `gorm:"size:32;unique"` //用户姓名
	PasswordHash string `gorm:"size:128"`       //密码哈希值
	Mobile       string `gorm:"size:11;unique"` //手机号码
	RealName     string `gorm:"size:32"`        //真实姓名，实名认证
	IDCard       string `gorm:"size:20;unique"` //身份证号码，实名认证
	AvatarUrl    string `gorm:"size:256"`       //头像地址，使用 fast_dfs 存储的路径
}
