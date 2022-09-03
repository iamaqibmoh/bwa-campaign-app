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

	//user Dependency
	repoUser := repository.NewUserRepository(db)
	servUser := service.NewUserService(repoUser)
	authServ := app.NewAuthService()
	contrUser := controller.NewUserController(servUser, authServ)

	//campaign Dependency
	repoCamp := repository.NewCampaignRepository(db)
	servCamp := service.NewCampaignService(repoCamp)
	contrCamp := controller.NewCampaignController(servCamp)

	//Transaction Dependency
	repoTrs := repository.NewTransactionRepository(db)
	servTrs := service.NewTransactionService(repoTrs, repoCamp)
	contrTrs := controller.NewTransactionController(servTrs)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	//user Endpoint
	api.POST("/users", contrUser.Register)
	api.POST("/sessions", contrUser.Login)
	api.POST("/email-checkers", contrUser.CheckEmailAvailable)
	api.POST("/avatars", middleware.AuthMiddleware(authServ, servUser), contrUser.UploadAvatar)

	//campaign Endpoint
	api.GET("/campaigns", contrCamp.GetCampaigns)
	api.GET("/campaigns/:id", contrCamp.GetCampaignById)
	api.POST("/campaigns", middleware.AuthMiddleware(authServ, servUser), contrCamp.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authServ, servUser), contrCamp.UpdateCampaign)

	//Campaign Image Endpoint
	api.POST("/campaign-images", middleware.AuthMiddleware(authServ, servUser), contrCamp.CreateCampaignImages)

	//Campaign Transaction Endpoint
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authServ, servUser), contrTrs.GetCampaignTransactions)

	router.Run()
}
