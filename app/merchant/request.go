package merchant

type HeaderInput struct {
	XApiKey     string `header:"x-api-key" binding:"required"`
	SecretToken string `header:"secret-token" binding:"required"`
}
