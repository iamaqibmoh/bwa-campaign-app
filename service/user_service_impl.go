package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) Register(request web.RegisterUserRequest) (*domain.User, error) {
	user := domain.User{}
	user.Name = request.Name
	user.Email = request.Email
	user.Occupation = request.Occupation

	bytesPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return &user, err
	}

	user.PasswordHash = string(bytesPass)
	user.Role = "user"

	save, err := s.repo.Save(user)

	if err != nil {
		return save, err
	}

	return save, nil
}

func (s *userServiceImpl) Login(request web.LoginUserRequest) (*domain.User, error) {
	email := request.Email
	password := request.Password

	findByEmail, err := s.repo.FindByEmail(email)
	if err != nil {
		return findByEmail, err
	}

	if findByEmail.Id == 0 {
		return findByEmail, errors.New("No user found on that email ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.PasswordHash), []byte(password))
	if err != nil {
		return findByEmail, err
	}

	return findByEmail, nil
}

func (s *userServiceImpl) IsEmailAvailable(input web.CheckEmailInput) (bool, error) {
	email := input.Email
	findByEmail, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if findByEmail.Id == 0 {
		return true, nil
	}

	return false, nil
}
