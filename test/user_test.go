package test

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"errors"
	"log"
	"testing"
)

func Repo() repository.UserRepository {
	db := app.DBConnect()
	repo := repository.NewUserRepository(db)
	return repo
}

func TestFindByEmailRepository(t *testing.T) {
	repo := Repo()
	findByEmail, err := repo.FindByEmail("aqibhh@test.com")
	if err != nil {
		log.Println(err.Error())
	}
	if findByEmail.Id == 0 {
		log.Println("User tidak ditemukan")
	} else {
		log.Println(findByEmail.Name)
	}
}

func TestLoginService(t *testing.T) {
	repo := Repo()
	serv := service.NewUserService(repo)

	login := web.LoginUserRequest{
		Email:    "ucup@test.com",
		Password: "iniUcp",
	}
	user, err := serv.Login(login)
	if err != nil {
		errors.New("Something is wrong ")
		errors.New(err.Error())
	}

	log.Println(user.Email)
	log.Println(user.Name)
}
