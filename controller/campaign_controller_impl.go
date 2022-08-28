package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	serv service.CampaignService
}

func NewCampaignController(serv service.CampaignService) CampaignController {
	return &CampaignControllerImpl{serv: serv}
}

func (contr *CampaignControllerImpl) GetCampaigns(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := contr.serv.GetCampaigns(id)

	campaignsResp := helper.CampaignResponseFormatters(campaigns)

	if err != nil {
		response := helper.APIResponseStruct("Failed get campaign", http.StatusBadRequest, "error", campaignsResp)
		helper.WriteToResponseBody(c.Writer, &response)
	}

	response := helper.APIResponseStruct("Get campaign successfully", http.StatusOK, "success", campaignsResp)
	helper.WriteToResponseBody(c.Writer, &response)
}
