package handler

import (
	"IkezawaYuki/a-root-backend/config"
	"IkezawaYuki/a-root-backend/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81/webhook"
	"io"
	"net/http"
	"os"
)

type PaymentHandler interface{}

type paymentHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewPaymentHandler(customerUsecase usecase.CustomerUsecase) PaymentHandler {
	return &paymentHandler{
		customerUsecase: customerUsecase,
	}
}

func (p *paymentHandler) HandleWebhook(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	// リクエストボディを読み込む
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Error reading request body"})
		return
	}

	// StripeのWebhookシークレットキー
	endpointSecret := config.Env.StripeEndpointSecret

	// Webhookイベントの検証
	event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), endpointSecret)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid webhook signature"})
		return
	}

	// イベントタイプに応じた処理
	switch event.Type {
	case "invoice.payment_failed":

	case "invoice.payment_succeeded":
		
	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
}
