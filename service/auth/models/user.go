/*
 *@time       2021/10/19 20:47
 *@version    1.0.0
 *@author     11726
 */

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primarykey" json:"user_id"` //id 编号
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"size:32;unique" json:"username"` //用户姓名
	PasswordHash string         `gorm:"size:128" json:"-"`              //密码哈希值
	Mobile       string         `gorm:"size:11;unique" json:"mobile"`   //手机号码
	RealName     string         `gorm:"size:32" json:"real_name"`       //真实姓名，实名认证
	IDCard       string         `gorm:"size:20;unique" json:"id_card"`  //身份证号码，实名认证
	AvatarUrl    string         `gorm:"size:256" json:"avatar_url"`     //头像地址，使用 fast_dfs 存储的路径
}
