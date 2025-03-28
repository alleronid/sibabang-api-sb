package auth

import (
	"lanaya/api/app/merchant"
	"lanaya/api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(authService Service, merchantService merchant.MerchantService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		clientKey := c.GetHeader("x-api-key")

		if clientKey == "" {
			response := utils.APIResponse("Client key is required", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, err := merchantService.GetMerchantByClientKey(clientKey)

		if err != nil {
			response := utils.APIResponse("Client key is unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if authHeader == "" {
			response := utils.APIResponse("Token b2b is required", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claimMerchant, ok := claim["merchant"].(map[string]interface{})

		if !ok {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		merchantIDFloat, ok := claimMerchant["MerchantId"].(float64)

		if !ok {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		merchantID := int(merchantIDFloat)

		merchant, err := merchantService.GetMerchantByID(merchantID)

		if err != nil {
			response := utils.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentMerchant", merchant)
	}
}
