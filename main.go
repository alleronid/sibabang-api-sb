package main

import (
	"lanaya/api/app/ayolinx"
	"lanaya/api/app/merchant"
	"lanaya/api/app/payment"
	"lanaya/api/auth"
	"lanaya/api/config"
	"lanaya/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.GetDB()

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	merchantRepository := merchant.NewRepository(db)
	paymentRepository := payment.NewRepository(db)
	merchantService := merchant.NewService(merchantRepository)
	ayolinxService := ayolinx.NewAyolinxService()
	paymentService := payment.NewService(paymentRepository, ayolinxService)
	authService := auth.NewService()

	authHandler := handler.NewUserHandler(merchantService, authService)
	paymentHandler := handler.NewPaymentHandler(paymentService, merchantService)

	api.POST("/auth/generate-token", authHandler.GenerateToken)
	api.POST("/payment/generate-qris", auth.AuthMiddleware(authService, merchantService), paymentHandler.GenerateTransaction)

	router.Run(":212")

}
