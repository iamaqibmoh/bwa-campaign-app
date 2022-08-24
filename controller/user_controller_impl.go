package controller

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/model/web"
	"BWA-CAMPAIGN-APP/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
	authService app.AuthService
}

func NewUserController(service service.UserService, authService app.AuthService) UserController {
	return &UserControllerImpl{userService: service, authService: authService}
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

	token, err := ct.authService.GenerateToken(register.Id)
	if err != nil {
		helper.UserServiceError("Token generate is error", c, err)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(register, token)
	apiResponse := helper.APIResponseStruct("Account has been registered", http.StatusOK, "success", userResponseFormatter)

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

	token, err := ct.authService.GenerateToken(user.Id)
	if err != nil {
		helper.UserServiceError("Token generate is error", c, err)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(user, token)
	response := helper.APIResponseStruct("Login Successfully", 200, "success", &userResponseFormatter)
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

	response := helper.APIResponseStruct(metaMessage, http.StatusOK, "success", data)
	c.JSON(200, &response)
}

func (ct *UserControllerImpl) UploadAvatar(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		helper.APIResponseStruct("Error upload your avatar", http.StatusOK, "error", data)
		return
	}

	path := fmt.Sprintf("images/%d-%s", 1, fileHeader.Filename)

	err = c.SaveUploadedFile(fileHeader, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		helper.APIResponseStruct("Error upload your avatar", http.StatusOK, "error", data)
		return
	}

	_, err = ct.userService.UpdateAvatar(1, path)
	if err != nil {
		helper.UserServiceError("Error save your avatar", c, err)
		return
	}

	avatar := helper.APIResponseStruct("Avatar successfully uploaded", 200, "success", gin.H{
		"is_uploaded": true,
	})
	c.JSON(200, &avatar)
}
