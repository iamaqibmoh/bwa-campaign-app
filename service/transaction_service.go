package service

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
)

type TransactionService interface {
	GetTransactionsByCampaignId(input web.GetCampaignTransactionsInput) ([]domain.Transaction, error)
}
