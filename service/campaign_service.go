package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
)

type CampaignService interface {
	GetCampaigns(userId int) ([]domain.Campaign, error)
	GetCampaignById(campId int) (domain.Campaign, error)
	CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error)
	UpdateCampaign(inputId web.GetCampaignDetailInput, inputData web.CreateCampaignInput) (domain.Campaign, error)
}
