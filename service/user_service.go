package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
)

type UserService interface {
	Register(user web.RegisterUserRequest) (domain.User, error)
}
