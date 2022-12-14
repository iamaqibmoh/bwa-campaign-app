package helper

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
	"strings"
)

func UserResponseFormatter(user domain.User, token string) *web.UserResponseFormatter {
	userResp := web.UserResponseFormatter{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return &userResp
}

func CampaignResponseFormatterBasic(campaign domain.Campaign) web.CampaignResponseFormatterBasic {
	camp := web.CampaignResponseFormatterBasic{}
	camp.Id = campaign.Id
	camp.UserId = campaign.UserId
	camp.Name = campaign.Name
	camp.Summary = campaign.Summary
	camp.GoalAmount = campaign.GoalAmount
	camp.CurrentAmount = campaign.CurrentAmount
	camp.Slug = campaign.Slug
	camp.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		camp.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return camp
}

func CampaignResponseFormatter(campaign domain.Campaign) web.CampaignResponseFormatter {
	campaignResp := web.CampaignResponseFormatter{}
	campaignResp.Id = campaign.Id
	campaignResp.Name = campaign.Name
	campaignResp.Summary = campaign.Summary
	campaignResp.Description = campaign.Description
	campaignResp.GoalAmount = campaign.GoalAmount
	campaignResp.CurrentAmount = campaign.CurrentAmount
	campaignResp.UserId = campaign.UserId
	campaignResp.Slug = campaign.Slug
	campaignResp.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignResp.ImageUrl = campaign.CampaignImages[0].FileName
	}

	//Campaign Perks
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignResp.Perks = perks
	//End Campaign Perks

	//Campaign User
	user := campaign.User
	campUser := web.CampaignUserFormatter{}
	campUser.Name = user.Name
	campUser.Avatar = user.Avatar

	campaignResp.User = campUser
	//End Campaign User

	//Campaign Images
	imageFormat := []web.CampaignImageFormatter{}
	for _, campaignImage := range campaign.CampaignImages {
		image := web.CampaignImageFormatter{}
		image.ImageUrl = campaignImage.FileName
		image.IsPrimary = false
		if campaignImage.IsPrimary == 1 {
			image.IsPrimary = true
		}
		imageFormat = append(imageFormat, image)
	}

	campaignResp.Images = imageFormat
	//End Campaign Images

	return campaignResp
}

func CampaignResponseFormatters(campaigns []domain.Campaign) []web.CampaignResponseFormatter {
	campResps := []web.CampaignResponseFormatter{}
	for _, campaign := range campaigns {
		formatter := CampaignResponseFormatter(campaign)
		campResps = append(campResps, formatter)
	}
	return campResps
}

func APIResponseStruct(message string, code int, status string, data interface{}) *web.WebResponse {
	meta := web.Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	webResp := web.WebResponse{
		Meta: meta,
		Data: data,
	}

	return &webResp
}

func CampaignTransactionFormatter(transaction domain.Transaction) web.CampaignTransactionFormatter {
	formatter := web.CampaignTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func CampaignTransactionFormatters(transactions []domain.Transaction) []web.CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []web.CampaignTransactionFormatter{}
	}

	var formatters []web.CampaignTransactionFormatter

	for _, tr := range transactions {
		formatters = append(formatters, CampaignTransactionFormatter(tr))
	}
	return formatters
}
