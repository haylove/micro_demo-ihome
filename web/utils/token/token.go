/*
 *@time       2021/10/20 3:28
 *@version    1.0.0
 *@author     11726
 */

package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Verifier interface {
	Verify(interface{}) (interface{}, error)
}

//JWTTokenVerifier which verify token and return account id
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (j *JWTTokenVerifier) Verify(token interface{}) (interface{}, error) {
	tokenString := token.(string)
	type cusClaim struct {
		jwt.StandardClaims
		Data interface{} `json:"data"`
	}
	t, err := jwt.ParseWithClaims(tokenString, &cusClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return j.PublicKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("cannot parser token:%v", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("token not valid")
	}

	claims, ok := t.Claims.(*cusClaim)
	if !ok {
		return nil, fmt.Errorf("token claims err")
	}
	if err = claims.Valid(); err != nil {
		return nil, fmt.Errorf("token not valid:%v", err)
	}

	return claims.Subject, nil
}
