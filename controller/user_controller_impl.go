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
		helper.UserRequestError(err, c)
		return
	}

	register, err := ct.userService.Register(user)
	if err != nil {
		helper.UserServiceError(err, c)
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
		helper.UserRequestError(err, c)
		return
	}

	user, err := ct.userService.Login(login)
	if err != nil {
		helper.UserServiceError(err, c)
		return
	}

	userResponseFormatter := helper.UserResponseFormatter(user, "tokentokentokentokentoken")
	response := helper.APIResponse("Login Successfully", 200, "success", &userResponseFormatter)
	c.JSON(200, &response)

	//user memasukkan input (email & password)
	//input ditangkap handler
	//mapping dari input user ke input struct
	//input struct passing ke service
	//di service mencari dg bantuan repository user dengan email x
	//mencocokkan password
}
