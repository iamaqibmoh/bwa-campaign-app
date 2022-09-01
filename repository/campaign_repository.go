package repository

import "BWA-CAMPAIGN-APP/model/domain"

type CampaignRepository interface {
	FindAll() ([]domain.Campaign, error)
	FindByUserId(userId int) ([]domain.Campaign, error)
	FindById(campId int) (domain.Campaign, error)
	Save(campaign domain.Campaign) (domain.Campaign, error)
	Update(campaign domain.Campaign) (domain.Campaign, error)
}
