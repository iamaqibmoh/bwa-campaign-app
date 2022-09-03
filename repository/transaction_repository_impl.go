package repository

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) FindByCampaignId(campId int) ([]domain.Transaction, error) {
	var trs []domain.Transaction
	err := r.db.Find(&trs).Where("campaign_id", campId).Preload("User").Order("id desc").Error

	if err != nil {
		return trs, err
	}

	return trs, nil
}
