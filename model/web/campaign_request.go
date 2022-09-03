package web

import "BWA-CAMPAIGN-APP/model/domain"

type GetCampaignDetailInput struct {
	CampId int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name        string      `json:"name" binding:"required"`
	Summary     string      `json:"summary" binding:"required"`
	Description string      `json:"description" binding:"required"`
	GoalAmount  int         `json:"goal_amount" binding:"required"`
	Perks       string      `json:"perks" binding:"required"`
	User        domain.User `json:"user"`
}

type CreateCampaignImageInput struct {
	CampaignId int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       domain.User
}

type GetCampaignTransactionsInput struct {
	Id   int `uri:"id" binding:"required"`
	User domain.User
}
