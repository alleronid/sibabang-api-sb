package main

import (
	"lanaya/api/app/ayolinx"
	"lanaya/api/app/merchant"
	"lanaya/api/app/payment"
	"lanaya/api/auth"
	"lanaya/api/config"
	"lanaya/api/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "212"
	}

	mode := os.Getenv("GIN_MODE")
	if mode != "" {
		gin.SetMode(mode)
	}

	config.ConnectDB()
	db := config.GetDB()

	router := gin.Default()
	router.Static("/public", "./public") // Serve static files from the public directory
	router.LoadHTMLGlob("public/*.html")
	api := router.Group("/api")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

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

	router.Run(":" + port)

}
