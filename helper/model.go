package helper

import (
	"BWA-CAMPAIGN-APP/model/domain"
	"BWA-CAMPAIGN-APP/model/web"
)

func UserResponseFormatter(user *domain.User, token string) *web.UserResponseFormatter {
	userResp := web.UserResponseFormatter{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return &userResp
}

func APIResponse(message string, code int, status string, data interface{}) *web.WebResponse {
	meta := web.Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	webResp := web.WebResponse{
		Meta: meta,
		Data: data,
	}

	return &webResp
}
