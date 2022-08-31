package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/web"
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
	helper.WriteToResponseBody(c.Writer, response)
}

func (contr *CampaignControllerImpl) GetCampaignById(c *gin.Context) {
	input := web.GetCampaignDetailInput{}
	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	campaign, err := contr.serv.GetCampaignById(input.CampId)
	if err != nil {
		helper.ServiceError("Failed get campaign by id", c, "error", err)
		return
	}

	apiResponseStruct := helper.APIResponseStruct("Success get campaign by id", 200, "success", helper.CampaignResponseFormatter(campaign))
	helper.WriteToResponseBody(c.Writer, apiResponseStruct)
}
