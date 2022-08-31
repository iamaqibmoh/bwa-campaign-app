package service

import "BWA-CAMPAIGN-APP/model/domain"

type CampaignService interface {
	GetCampaigns(userId int) ([]domain.Campaign, error)
	GetCampaignById(campId int) (domain.Campaign, error)
}
