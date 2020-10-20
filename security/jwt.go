package security

import (
	"github.com/dgrijalva/jwt-go"
	"go-flutter/model"
	"time"
)

const SECRET_KEY = "duynvh"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return result, nil
}
