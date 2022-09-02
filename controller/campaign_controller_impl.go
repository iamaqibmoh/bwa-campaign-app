package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/service"
	"fmt"
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

func (contr *CampaignControllerImpl) CreateCampaign(c *gin.Context) {
	input := web.CreateCampaignInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	currentUser := c.MustGet("currentUser").(domain.User)
	input.User = currentUser

	campaign, err := contr.serv.CreateCampaign(input)
	if err != nil {
		helper.ServiceError("Failed create campaign", c, err.Error(), err)
		return
	}

	response := helper.APIResponseStruct("Your campaign was successfully created", 200, "success", helper.CampaignResponseFormatterBasic(campaign))
	helper.WriteToResponseBody(c.Writer, response)
}

func (contr *CampaignControllerImpl) UpdateCampaign(c *gin.Context) {
	inputId := web.GetCampaignDetailInput{}
	err := c.ShouldBindUri(&inputId)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	inputData := web.CreateCampaignInput{}
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	user := c.MustGet("currentUser").(domain.User)
	inputData.User = user

	campaign, err := contr.serv.UpdateCampaign(inputId, inputData)
	if err != nil {
		helper.ServiceError("Failed to update campaign", c, err.Error(), err)
		return
	}

	response := helper.APIResponseStruct("Campaign is successfully updated", 200, "success", helper.CampaignResponseFormatterBasic(campaign))
	c.JSON(200, response)
}

func (contr *CampaignControllerImpl) CreateCampaignImages(c *gin.Context) {
	var input web.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		helper.RequestError(c, err)
		return
	}

	user := c.MustGet("currentUser").(domain.User)
	input.User = user

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponseStruct("Error upload your campaign image", http.StatusBadRequest, "error", data)
		c.JSON(400, &response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", user.Id, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponseStruct("Error upload your campaign image", http.StatusBadRequest, "error", data)
		c.JSON(400, &response)
		return
	}

	_, err = contr.serv.CreateCampaignImages(input, path)
	if err != nil {
		helper.ServiceError("Error upload your campaign image", c, err.Error(), err)
		return
	}

	data := gin.H{"is_primary": true}
	response := helper.APIResponseStruct("Campaign image is successfully uploaded", 200, "success", data)
	c.JSON(200, response)
}
