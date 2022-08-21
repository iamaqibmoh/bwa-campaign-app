package test

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {
	db := app.DBConnect()
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)

	user := web.RegisterUser{
		Name:       "Sayidin",
		Email:      "sayidin@test.com",
		Occupation: "UI/UX",
		Password:   "sayidinAjah",
	}

	regis, _ := serv.Register(user)
	fmt.Println("Register Success")
	fmt.Println(regis)
}
