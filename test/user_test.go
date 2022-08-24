package test

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"encoding/json"
	"errors"
	"fmt"
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

func TestErr(t *testing.T) {
	err := 0
	if err < 1 {
		helper.ReturnIfError(errors.New("Ups"))
		fmt.Println("Ini gajalan")
	}
}

func TestAvatarUpdate(t *testing.T) {
	repo := Repo()
	serv := service.NewUserService(repo)
	user, err := serv.UpdateAvatar(1, "/images/1-avatar.jpg")
	helper.ReturnIfError(err)
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
}

func TestJWTGenerate(t *testing.T) {
	authServ := app.NewAuthService()
	token, _ := authServ.GenerateToken(1001)
	log.Println(token)
}

func TestValidateToken(t *testing.T) {
	authService := app.NewAuthService()
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.Htu9vGjyUuiHupXLaSxrt3FKEoShkvqVZgWsHYB0oYU")
	helper.ReturnIfError(err)
	log.Println("Valid")
	log.Println(token)
}
