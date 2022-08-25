package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func formatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

func UserRequestError(c *gin.Context, err error) {
	if err != nil {
		validationError := formatValidationError(err)
		errMessage := gin.H{"errors": validationError}
		apiResponse := APIResponseStruct("Ups, error create your request", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusBadRequest, &apiResponse)
	}
}

func UserServiceError(message string, c *gin.Context, data string, err error) {
	if err != nil {
		apiResponse := APIResponseStruct(message, http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, &apiResponse)
	}
}
