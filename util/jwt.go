package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var myKey = []byte("AilzrMusicManager-jwt-key")

func GenerateToken(identity, name string) (string, error) {
	userClaims := &UserClaims{
		identity,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Hour * 24 * 7)},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	//fmt.Println(tokenString)
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaims, nil
}
