package controller

import (
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{userService: service}
}

func (ct *UserControllerImpl) Register(c *gin.Context) {
	user := web.RegisterUserRequest{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.UserRequestError(c, err)
		return
	}

	register, err := ct.userService.Register(user)
	if err != nil {
		helper.UserServiceError("Register user failed", c, err)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(register, "tokentokentokentokentoken")
	apiResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", userResponseFormatter)

	c.JSON(http.StatusOK, &apiResponse)
}

func (ct *UserControllerImpl) Login(c *gin.Context) {
	login := web.LoginUserRequest{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		helper.UserRequestError(c, err)
		return
	}

	user, err := ct.userService.Login(login)
	if err != nil {
		helper.UserServiceError("Login failed", c, err)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(user, "tokentokentokentokentoken")
	response := helper.APIResponse("Login Successfully", 200, "success", &userResponseFormatter)
	c.JSON(200, &response)
}

func (ct *UserControllerImpl) CheckEmailAvailable(c *gin.Context) {
	email := web.CheckEmailInput{}
	err := c.ShouldBindJSON(&email)
	if err != nil {
		helper.UserRequestError(c, err)
		return
	}
	isEmailAvailable, err := ct.userService.IsEmailAvailable(email)
	if err != nil {
		helper.UserServiceError("Your email not available", c, err)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email available to use"
	}
	data := gin.H{"is_available": isEmailAvailable}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(200, &response)
}
