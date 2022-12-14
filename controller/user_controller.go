package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	CheckEmailAvailable(c *gin.Context)
	UploadAvatar(c *gin.Context)
}
