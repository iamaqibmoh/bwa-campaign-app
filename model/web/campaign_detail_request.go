package web

type GetCampaignDetailInput struct {
	CampId int `uri:"id" binding:"required"`
}
