package handler

import (
	"lanaya/api/app/merchant"
	"lanaya/api/auth"
	"lanaya/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	merchantService merchant.MerchantService
	authService     auth.Service
}

func NewUserHandler(merchantService merchant.MerchantService, authService auth.Service) *authHandler {
	return &authHandler{merchantService, authService}
}

func (h *authHandler) GenerateToken(c *gin.Context) {
	var input merchant.HeaderInput

	err := c.ShouldBindHeader(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Input Required", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getMerchant, err := h.merchantService.GetMerchant(input)

	if err != nil {
		response := utils.APIResponse("Authentication Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(getMerchant)

	if err != nil {
		response := utils.APIResponse("Generate Token Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := merchant.FormatAuth(getMerchant, token)
	response := utils.APIResponse("Token Successfully Generated", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
