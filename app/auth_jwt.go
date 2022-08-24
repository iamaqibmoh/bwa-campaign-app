package app

import (
	"BWA-CAMPAIGN-APP/helper"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(Id int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

var secret_key = []byte("ini_s3creT_K3y")

func (a *AuthServiceImpl) GenerateToken(Id int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = Id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedString, err := token.SignedString(secret_key)
	helper.ReturnIfError(err)

	return signedString, nil
}

func (a *AuthServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token Method")
		}
		return []byte(secret_key), nil
	})
	helper.ReturnIfError(err)
	return token, nil
}
