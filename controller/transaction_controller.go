package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	GetCampaignTransactions(c *gin.Context)
}
