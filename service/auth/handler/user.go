package handler

import (
	"context"

	"auth/dao"
	user "auth/proto/user"
)

type User struct {
	*dao.UserManger
}

func NewUser(driverName, dataSourceName string) (*User, error) {
	userManger, err := dao.NewUserManger(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &User{userManger}, nil
}

func (u *User) Register(ctx context.Context, user *user.User, userEntity *user.UserEntity) error {
	createUser, err := u.CreateUser(user)
	if err != nil {
		return err
	}

	userEntity.Id = uint32(createUser.ID)
	userEntity.User = user
	userEntity.User.Password = createUser.PasswordHash

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
