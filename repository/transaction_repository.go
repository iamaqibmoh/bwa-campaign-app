package repository

import "BWA-CAMPAIGN-APP/model/domain"

type TransactionRepository interface {
	FindByCampaignId(campId int) ([]domain.Transaction, error)
}
