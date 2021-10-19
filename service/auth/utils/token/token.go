/*
 *@time       2021/10/20 2:38
 *@version    1.0.0
 *@author     11726
 */

package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Generator interface {
	GenerateToken(subject string, expire time.Duration) (token string, err error)
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
func (j *JWTTokenGen) GenerateToken(subject string, expire time.Duration) (token string, err error) {
	genTime := j.timeFunc()
	claims := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Subject:   subject,
		Issuer:    j.issuer,
		IssuedAt:  genTime.Unix(),
		ExpiresAt: genTime.Add(expire).Unix(),
	})
	return claims.SignedString(j.privateKey)
}
