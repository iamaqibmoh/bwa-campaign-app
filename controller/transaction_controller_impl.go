package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
)

type TransactionControllerImpl struct {
	serv service.TransactionService
}

func NewTransactionController(serv service.TransactionService) TransactionController {
	return &TransactionControllerImpl{serv: serv}
}

func (contr *TransactionControllerImpl) GetCampaignTransactions(c *gin.Context) {
	input := web.GetCampaignTransactionsInput{}
	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	user := c.MustGet("currentUser").(domain.User)
	input.User = user

	transactions, err := contr.serv.GetTransactionsByCampaignId(input)
	if err != nil {
		helper.ServiceError("Failed to get campaign's transactions", c, err.Error(), err)
		return
	}

	response := helper.APIResponseStruct("You're got campaign's transactions successfully", 200, "success", helper.CampaignTransactionFormatters(transactions))
	c.JSON(200, response)
}
