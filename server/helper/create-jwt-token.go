package helper

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func GetJwtToken() string {
	var jwtKey = []byte("my_secret_key")
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 24*3600),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Panic(err.Error())
		return ""
	}
	return tokenString
}
