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
