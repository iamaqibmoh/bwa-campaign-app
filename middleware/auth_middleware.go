package middleware

import (
	"BWA-CAMPAIGN-APP/app"
	"BWA-CAMPAIGN-APP/helper"
	"BWA-CAMPAIGN-APP/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(authService app.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.Contains(header, "Bearer") {
			response := helper.APIResponseStruct("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var token string
		arrayToken := strings.Split(header, " ")
		token = arrayToken[1]

		validateToken, err := authService.ValidateToken(token)
		if err != nil {
			response := helper.APIResponseStruct("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			response := helper.APIResponseStruct("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user_id := int(claims["user_id"].(float64))

		user, err := userService.GetUserById(user_id)
		if err != nil {
			response := helper.APIResponseStruct("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
