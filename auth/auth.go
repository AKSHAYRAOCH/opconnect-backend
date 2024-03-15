package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generatejwt(userid, usertype string) (string, error) {

	claims := jwt.MapClaims{
		"username": userid,
		"usertype": usertype,
		"iat":      time.Now().Unix(),
	}

	jwt, err := generatewithwithclaims(claims)

	if err != nil {
		return "", err
	}
	return jwt, nil
}

func generatewithwithclaims(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedtoken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return signedtoken, nil
}
