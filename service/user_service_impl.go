package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &UserServiceImpl{repo: repo}
}

func (service *UserServiceImpl) Register(request web.RegisterUserRequest) (domain.User, error) {
	user := domain.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation

	bytesPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(bytesPass)
	user.Role = "user"

	save, err := service.repo.Save(user)

	if err != nil {
		return save, err
	}

	return save, nil
}
