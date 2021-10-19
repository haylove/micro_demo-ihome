/*
 *@time       2021/10/19 20:58
 *@version    1.0.0
 *@author     11726
 */

package dao

import (
	"auth/utils/hash"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auth/models"
	"auth/proto/user"
)

type UserManger struct {
	db *gorm.DB
}

func NewUserManger(driverName, dataSourceName string) (*UserManger, error) {
	var (
		db  *gorm.DB
		err error
	)

	switch driverName {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dataSourceName))
	default:
		err = fmt.Errorf("not supported driver: %s", driverName)
	}

	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&models.User{})

	return &UserManger{db: db}, nil
}

func (m *UserManger) CreateUser(user *auth.User) (*models.User, error) {

	password := hash.Encryption(user.Password)

	u := &models.User{
		Name:         user.MobilePhone,
		PasswordHash: password,
		Mobile:       user.MobilePhone,
	}

	result := m.db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (m *UserManger) UpdateUser() {}

func (m *UserManger) DeleteUser() {}

func (m *UserManger) GetUser(user *auth.LoginReq) (*models.User, bool) {
	var u models.User
	u.Name = user.Username
	u.PasswordHash = hash.Encryption(user.Password)
	result := m.db.First(&u)
	if result.Error != nil {
		return nil, false
	}
	return &u, true
}
