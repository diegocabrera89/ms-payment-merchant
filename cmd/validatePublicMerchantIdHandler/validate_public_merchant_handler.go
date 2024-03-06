package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/middleware/metadata"
	"github.com/diegocabrera89/ms-payment-core/response"
	"github.com/diegocabrera89/ms-payment-merchant/constantsmicro"
	"github.com/diegocabrera89/ms-payment-merchant/service"
	"net/http"
)

// MerchantHandler handles HTTP requests related to the Merchant entity.
type MerchantHandler struct {
	merchantService *service.MerchantService
}

// NewMerchantHandler create a new NewMerchantHandler instance.
func NewMerchantHandler() *MerchantHandler {
	return &MerchantHandler{
		merchantService: service.NewMerchantService(),
	}
}

// ValidatePublicMerchantIDHandler handler for ValidatePublicMerchantIDHandler get merchant.
func (h *MerchantHandler) ValidatePublicMerchantIDHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("ValidatePublicMerchantIDHandler", ctx, request)
	createMerchantHandler, errorMerchantHandler := h.merchantService.ValidatePublicMerchantIDService(ctx, request)
	if errorMerchantHandler != nil {
		logs.LogTrackingError("ValidatePublicMerchantIDHandler", "ValidatePublicMerchantIDService", ctx, request, errorMerchantHandler)
		return response.ErrorResponse(http.StatusInternalServerError, constantsmicro.ErrorGettingMerchant)
	}
	return createMerchantHandler, nil
}

func main() {
	// Create an instance of PetHandler in the main function.
	merchantHandler := NewMerchantHandler()

	// Wrap the handler function with logging middleware.
	handlerWithLogging := metadata.MiddlewareMetadata(merchantHandler.ValidatePublicMerchantIDHandler)

	// Start the Lambda handler with the handler function wrapped in the middleware.
	lambda.Start(handlerWithLogging)
}
