package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
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

	return campaigns, nil
}

func (s *CampaignServiceImpl) GetCampaignById(campId int) (domain.Campaign, error) {
	campaign, err := s.repo.FindById(campId)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignServiceImpl) CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error) {
	campaign := domain.Campaign{}
	campaign.Name = input.Name
	campaign.Summary = input.Summary
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.Id

	slugText := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugText)

	save, err := s.repo.Save(campaign)
	if err != nil {
		return save, err
	}

	return save, nil
}

func (s *CampaignServiceImpl) UpdateCampaign(inputId web.GetCampaignDetailInput, inputData web.CreateCampaignInput) (domain.Campaign, error) {
	campaign, err := s.repo.FindById(inputId.CampId)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("You're not authorized to change this campaign")
	}

	campaign.Name = inputData.Name
	campaign.Summary = inputData.Summary
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks
	campaign.UserId = inputData.User.Id

	update, err := s.repo.Update(campaign)
	return update, err
}
