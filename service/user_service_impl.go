package service

import (
	"BWA-CAMPAIGN-APP/helper"
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
	helper.ReturnIfError(err)

	user.PasswordHash = string(bytesPass)
	user.Role = "user"

	save, err := s.repo.Save(user)
	helper.ReturnIfError(err)

	return save, nil
}

func (s *userServiceImpl) Login(request web.LoginUserRequest) (*domain.User, error) {
	email := request.Email
	password := request.Password

	findByEmail, err := s.repo.FindByEmail(email)
	helper.ReturnIfError(err)

	if findByEmail.Id == 0 {
		return findByEmail, errors.New("No user found on that email ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(findByEmail.PasswordHash), []byte(password))
	helper.ReturnIfError(err)

	return findByEmail, nil
}

func (s *userServiceImpl) IsEmailAvailable(input web.CheckEmailInput) (bool, error) {
	email := input.Email
	findByEmail, err := s.repo.FindByEmail(email)
	helper.ReturnIfError(err)

	if findByEmail.Id == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userServiceImpl) UpdateAvatar(id int, avatarLocation string) (*domain.User, error) {
	findById, err := s.repo.FindById(id)
	helper.ReturnIfError(err)

	findById.Avatar = avatarLocation

	update, err := s.repo.Update(findById)
	helper.ReturnIfError(err)

	return update, nil
}
