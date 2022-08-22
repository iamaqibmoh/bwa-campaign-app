package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
)

type UserService interface {
	Register(request web.RegisterUserRequest) (*domain.User, error)
	Login(request web.LoginUserRequest) (*domain.User, error)
	IsEmailAvailable(input web.CheckEmailInput) (bool, error)
}
