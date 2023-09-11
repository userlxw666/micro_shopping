package utils

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserName string
	Password string
	jwt.StandardClaims
}

var MySecret = "microShopping"

func CreateToken(username, password string) (string, error) {
	c := MyClaims{
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "lxw",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(MySecret))
	if err != nil {
		fmt.Println("create tokenString error", err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) *MyClaims {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println("parse token error", err)
		return nil
	}
	if myClaim, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return myClaim
	} else {
		fmt.Println("断言失败", errors.New("断言失败"))
		return nil
	}
}
