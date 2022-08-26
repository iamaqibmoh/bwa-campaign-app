package test

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var db = app.DBConnect()

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
	repo := repository.NewCampaignRepository(db)
	campaigns, err := repo.FindByUserId(1)
	IfError(err)
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}
}
