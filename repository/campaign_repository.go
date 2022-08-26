package repository

import "BWA-CAMPAIGN-APP/model/domain"

type CampaignRepository interface {
	FindAll() ([]domain.Campaign, error)
	FindByUserId(userId int) ([]domain.Campaign, error)
}
