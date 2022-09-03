package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/repository"
	"errors"
)

type TransactionServiceImpl struct {
	repo               repository.TransactionRepository
	campaignRepository repository.CampaignRepository
}

func NewTransactionService(repo repository.TransactionRepository, campaignRepository repository.CampaignRepository) TransactionService {
	return &TransactionServiceImpl{
		repo:               repo,
		campaignRepository: campaignRepository,
	}
}

func (s *TransactionServiceImpl) GetTransactionsByCampaignId(input web.GetCampaignTransactionsInput) ([]domain.Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []domain.Transaction{}, err
	}

	if input.User.Id != campaign.UserId {
		return []domain.Transaction{}, errors.New("You're not an owner of the campaign")
	}

	transactions, err := s.repo.FindByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
