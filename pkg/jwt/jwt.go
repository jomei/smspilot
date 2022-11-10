package jwt

import (
	"SmsPilot2/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	IdUser int `json:"id-user"`
	jwt.StandardClaims
}

func Get(userId int, expiresTime int64) (string, error) {
	// load SigningKey
	mySigningKey := []byte(config.Conf.SIGNINGKEY)
	// Create the Claims
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresTime, // Unix time string
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func Decode(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.SIGNINGKEY), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, err
	} else {
		return claims, err
	}
}
