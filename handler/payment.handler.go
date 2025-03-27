package handler

import (
	"lanaya/api/app/merchant"
	"lanaya/api/app/payment"
	"lanaya/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentHandler struct {
	paymentService payment.PaymentService
	authService    merchant.MerchantService
}

func NewPaymentHandler(paymentService payment.PaymentService, authService merchant.MerchantService) *paymentHandler {
	return &paymentHandler{paymentService, authService}
}

func (h *paymentHandler) GenerateTransaction(c *gin.Context) {
	var input payment.PaymentInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.APIResponse("Generate Transaction Failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentMerchant := c.MustGet("currentMerchant").(merchant.Merchant)
	input.Merchant = currentMerchant

	cekPayment, err := h.paymentService.GetTransaction(input.TrxId)

	if err != nil {
		response := utils.APIResponse("Failed to create payment", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if cekPayment.TrxId == input.TrxId {
		response := utils.APIResponse("Payment already exists", http.StatusConflict, "error", nil)
		c.JSON(http.StatusConflict, response)
		return
	}

	newPayment, err := h.paymentService.SavePayment(input)

	if err != nil {
		response := utils.APIResponse("Failed to create payment", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := payment.FormatPayment(newPayment)
	response := utils.APIResponse("Successfully create payment", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
