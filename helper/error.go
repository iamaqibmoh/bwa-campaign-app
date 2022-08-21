package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func FormatValidationError(err error) *[]string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return &errors
}

func UserRequestError(err error, c *gin.Context) {
	if err != nil {
		validationError := FormatValidationError(err)
		errMessage := gin.H{"errors": validationError}
		apiResponse := APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusBadRequest, &apiResponse)
		return
	}
}

func ForwardToServiceError(err error, c *gin.Context) {
	if err != nil {
		apiResponse := APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, &apiResponse)
		return
	}
}
