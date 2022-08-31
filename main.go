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

	//user endpoint
	api.POST("/users", contr.Register)
	api.POST("/sessions", contr.Login)
	api.POST("/email-checkers", contr.CheckEmailAvailable)
	api.POST("/avatars", middleware.AuthMiddleware(authServ, serv), contr.UploadAvatar)

	//campaign endpoint
	api.GET("/campaigns", contrCamp.GetCampaigns)
	api.GET("/campaigns/:id", contrCamp.GetCampaignById)
	api.POST("/campaigns", middleware.AuthMiddleware(authServ, serv), contrCamp.CreateCampaign)

	router.Run()
}
