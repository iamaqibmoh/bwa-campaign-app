package main

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/controller"
	"BWA-CAMPAIGN-APP/repository"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := app.DBConnect()
	repo := repository.NewUserRepository(db)
	serv := service.NewUserService(repo)
	contr := controller.NewUserController(serv)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", contr.Register)
	api.POST("/sessions", contr.Login)

	router.Run()
}
