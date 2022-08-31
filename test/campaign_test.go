package test

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var db = app.DBConnect()
var repo = repository.NewCampaignRepository(db)
var serv = service.NewCampaignService(repo)

func IfError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func TestCampaignGetAll(t *testing.T) {
	repo := repository.NewCampaignRepository(db)
	campaigns, err := repo.FindAll()
	IfError(err)
	assert.Equal(t, 3, len(campaigns))
}

func TestCampaignFindByUserId(t *testing.T) {
	campaigns, err := repo.FindByUserId(1)
	IfError(err)
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}
}

func TestGetCampaignsService(t *testing.T) {
	campaigns, err := serv.GetCampaigns(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, campaign := range campaigns {
		log.Println(campaign)
	}
}

func TestCreateCampaignService(t *testing.T) {
	repoUser := repository.NewUserRepository(db)
	servUser := service.NewUserService(repoUser)

	input := web.CreateCampaignInput{}
	input.Name = "Penggalangan Dana Startup"
	input.Summary = "Ini penggalangan dana"
	input.Description = "Deskripsi ini panjang sekali"
	input.GoalAmount = 150000000
	input.Perks = "hadiah satu, ini dua, itu tiga"

	user, _ := servUser.GetUserById(2)
	input.User = user

	campaign, err := serv.CreateCampaign(input)
	IfError(err)

	log.Println(campaign)
}
