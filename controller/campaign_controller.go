package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	GetCampaigns(c *gin.Context)
	GetCampaignById(c *gin.Context)
	CreateCampaign(c *gin.Context)
}
