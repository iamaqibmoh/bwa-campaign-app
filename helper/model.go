package helper

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
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
	campaignResp.UserId = campaign.UserId
	campaignResp.Name = campaign.Name
	campaignResp.Summary = campaign.Summary
	campaignResp.GoalAmount = campaign.GoalAmount
	campaignResp.CurrentAmount = campaign.CurrentAmount
	campaignResp.Slug = campaign.Slug
	campaignResp.ImageUrl = ""
	if len(campaign.CampaignImages) > 0 {
		campaignResp.ImageUrl = campaign.CampaignImages[0].FileName
	}
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
