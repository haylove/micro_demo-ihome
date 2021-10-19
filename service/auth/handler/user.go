package handler

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"auth/dao"
	user "auth/proto/user"
	"auth/utils/token"
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

func (u *User) Login(ctx context.Context, req *user.LoginReq, res *user.LoginRes) error {
	getUser, isExist := u.GetUser(req)
	if !isExist {
		res.LoginSuccess = isExist
		return nil

	}
	file, err := os.Open("config/test.privateKey")
	if err != nil {
		return err
	}
	defer file.Close()

	readAll, err := ioutil.ReadAll(file)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readAll)
	if err != nil {
		return err
	}

	gen := token.NewJWTTokenGen("ihome", privateKey)
	expire := time.Hour * 24

	generateToken, err := gen.GenerateToken(strconv.Itoa(int(getUser.ID)), expire)
	if err != nil {
		return err
	}
	res.LoginSuccess = isExist
	res.Msg = &user.Msg{
		Token:       generateToken,
		ExpireInSec: int32(expire.Seconds()),
	}

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
