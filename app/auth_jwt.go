package app

import (
	"BWA-CAMPAIGN-APP/helper"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(Id int) (string, error)
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
