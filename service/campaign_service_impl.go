package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/repository"
)

type CampaignServiceImpl struct {
	repo repository.CampaignRepository
}

func NewCampaignService(campaignRepository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{repo: campaignRepository}
}

func (s *CampaignServiceImpl) GetCampaigns(userId int) ([]domain.Campaign, error) {
	if userId == 0 {
		campaigns, err := s.repo.FindAll()
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repo.FindByUserId(userId)
	if err != nil {
		return campaigns, err
	}

	return campaigns, err
}
