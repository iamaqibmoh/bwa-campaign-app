package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{service: service}
}

func (ct *UserControllerImpl) RegisterUser(c *gin.Context) {
	user := web.RegisterUserRequest{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.UserRequestError(err, c)
		return
	}

	register, err := ct.service.Register(user)
	if err != nil {
		helper.ForwardToServiceError(err, c)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(&register, "tokentokentokentokentoken")
	apiResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", userResponseFormatter)

	c.JSON(http.StatusOK, &apiResponse)
}
