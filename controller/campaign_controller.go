package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	GetCampaigns(c *gin.Context)
}
