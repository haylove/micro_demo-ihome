/*
 *@time       2021/10/20 2:38
 *@version    1.0.0
 *@author     11726
 */

package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"os"
	"time"
)

type Generator interface {
	GenerateToken(subject interface{}, expire time.Duration) (token string, err error)
}

type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	timeFunc   func() time.Time
}

// JWTTokenGen Constructor
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		privateKey: privateKey,
		issuer:     issuer,
		timeFunc:   time.Now,
	}
}

//GenerateToken Generate A Token with RSA512
func (j *JWTTokenGen) GenerateToken(subject interface{}, expire time.Duration) (token string, err error) {
	genTime := j.timeFunc()

	claims := jwt.NewWithClaims(
		jwt.SigningMethodRS512,
		struct {
			jwt.StandardClaims
			Data interface{} `json:"data"`
		}{
			jwt.StandardClaims{
				Issuer:    j.issuer,
				IssuedAt:  genTime.Unix(),
				ExpiresAt: genTime.Add(expire).Unix()},
			subject,
		})
	return claims.SignedString(j.privateKey)
}

var DefaultTokenGen = initTokenGen()

func initTokenGen() *JWTTokenGen {
	file, _ := os.Open("config/test.privateKey")

	defer file.Close()

	readAll, _ := ioutil.ReadAll(file)
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM(readAll)

	return NewJWTTokenGen("ihome", privateKey)
}
