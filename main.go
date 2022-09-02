package main

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/controller"
	"BWA-CAMPAIGN-APP/middleware"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := app.DBConnect()

	//user dependency
	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)
	authServ := app.NewAuthService()
	contr := controller.NewUserController(serv, authServ)

	//campaign dependency
	repoCamp := repository.NewCampaignRepository(db)
	servCamp := service.NewCampaignService(repoCamp)
	contrCamp := controller.NewCampaignController(servCamp)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//user Endpoint
	api.POST("/users", contr.Register)
	api.POST("/sessions", contr.Login)
	api.POST("/email-checkers", contr.CheckEmailAvailable)
	api.POST("/avatars", middleware.AuthMiddleware(authServ, serv), contr.UploadAvatar)

	//campaign Endpoint
	api.GET("/campaigns", contrCamp.GetCampaigns)
	api.GET("/campaigns/:id", contrCamp.GetCampaignById)
	api.POST("/campaigns", middleware.AuthMiddleware(authServ, serv), contrCamp.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authServ, serv), contrCamp.UpdateCampaign)

	//Campaign Image Endpoint
	api.POST("/campaign-images", middleware.AuthMiddleware(authServ, serv), contrCamp.CreateCampaignImages)

	router.Run()
}
