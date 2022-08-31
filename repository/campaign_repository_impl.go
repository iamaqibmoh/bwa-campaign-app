package repository

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"gorm.io/gorm"
)

type campaignRepositoryImpl struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepositoryImpl{db: db}
}

func (r *campaignRepositoryImpl) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepositoryImpl) FindByUserId(userId int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign

	err := r.db.Where("user_id=?", userId).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepositoryImpl) FindById(campId int) (domain.Campaign, error) {
	campaign := domain.Campaign{}
	err := r.db.Preload("CampaignImages").Preload("User").Where("id=?", campId).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
