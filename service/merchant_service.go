package service

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-core/constantscore"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-core/response"
	"github.com/diegocabrera89/ms-payment-merchant/constantsmicro"
	"github.com/diegocabrera89/ms-payment-merchant/models"
	"github.com/diegocabrera89/ms-payment-merchant/repository"
	"github.com/diegocabrera89/ms-payment-merchant/utils"
	"net/http"
)

// MerchantService represents the service for the MerchantService entity.
type MerchantService struct {
	merchantRepo *repository.MerchantRepositoryImpl
}

// NewMerchantService create a new MerchantService instance.
func NewMerchantService() *MerchantService {
	return &MerchantService{
		merchantRepo: repository.NewMerchantRepository(),
	}
}

// CreateMerchantService handles the creation of a new merchant.
func (r *MerchantService) CreateMerchantService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("CreateMerchantService", ctx, request)
	var merchant models.Merchant
	err := json.Unmarshal([]byte(request.Body), &merchant)
	if err != nil {
		logs.LogTrackingError("CreateMerchantService", "JSON Unmarshal", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.InvalidRequestBody)
	}

	utils.BuildCreateMerchant(&merchant)

	createMerchant, errorMerchantRepository := r.merchantRepo.CreateMerchantRepository(ctx, request, merchant)
	if errorMerchantRepository != nil {
		logs.LogTrackingError("CreateMerchantService", "CreateMerchantRepository", ctx, request, errorMerchantRepository)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.ErrorCreatingItem)
	}

	responseBody, err := json.Marshal(createMerchant)
	if err != nil {
		logs.LogTrackingError("CreateMerchantService", "JSON Marshal", ctx, request, err)
		return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
	}
	return response.SuccessResponse(http.StatusCreated, responseBody, constantscore.ItemCreatedSuccessfully)
}

// ValidatePublicMerchantIDService handles the creation of a new pet.
func (r *MerchantService) ValidatePublicMerchantIDService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("ValidatePublicMerchantIDService", ctx, request)
	logs.LogTrackingInfoData("ValidatePublicMerchantIDService request", request, ctx, request)
	var responseBody []byte
	publicID := request.PathParameters[constantsmicro.PublicID]
	logs.LogTrackingInfoData("ValidatePublicMerchantIDService publicID", publicID, ctx, request)
	if publicID == "" {
		logs.LogTrackingError("ValidatePublicMerchantIDService", "PathParameters publicID", ctx, request, nil)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorCreatingItem)
	}

	getPublicID, err := r.merchantRepo.ValidatePublicMerchantIdRepository(ctx, request, publicID)
	if err != nil {
		logs.LogTrackingError("ValidatePublicMerchantIDService", "ValidatePublicMerchantIdRepository", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorCreatingItem)
	}

	logs.LogTrackingInfoData("ValidatePublicMerchantIDService getPublicID", getPublicID, ctx, request)
	if getPublicID.MerchantID != "" {
		responseBody, err = json.Marshal(getPublicID)
		if err != nil {
			logs.LogTrackingError("ValidatePublicMerchantIDService", "JSON Marshal", ctx, request, err)
			return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
		}
		return response.SuccessResponse(http.StatusOK, responseBody, constantscore.ItemSuccessfullyObtained)
	}
	return response.SuccessResponse(http.StatusOK, responseBody, constantscore.DataNotFound)
}

// ValidatePrivateMerchantIDService handles the creation of a new pet.
func (r *MerchantService) ValidatePrivateMerchantIDService(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.LogTrackingInfo("ValidatePrivateMerchantIDService", ctx, request)
	logs.LogTrackingInfoData("ValidatePrivateMerchantIDService request", request, ctx, request)
	var responseBody []byte
	privateID := request.PathParameters[constantsmicro.PrivateID]
	logs.LogTrackingInfoData("ValidatePublicMerchantIDService privateID", privateID, ctx, request)
	if privateID == "" {
		logs.LogTrackingError("ValidatePublicMerchantIDService", "PathParameters privateID", ctx, request, nil)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorCreatingItem)
	}

	getPrivateID, err := r.merchantRepo.ValidatePrivateMerchantIdRepository(ctx, request, privateID)
	if err != nil {
		logs.LogTrackingError("ValidatePublicMerchantIDService", "ValidatePrivateMerchantIdRepository", ctx, request, err)
		return response.ErrorResponse(http.StatusBadRequest, constantscore.ErrorCreatingItem)
	}

	logs.LogTrackingInfoData("ValidatePublicMerchantIDService getPrivateID", getPrivateID, ctx, request)
	if getPrivateID.MerchantID != "" {
		responseBody, err = json.Marshal(getPrivateID)
		if err != nil {
			logs.LogTrackingError("ValidatePublicMerchantIDService", "JSON Marshal", ctx, request, err)
			return response.ErrorResponse(http.StatusInternalServerError, constantscore.InvalidResponseBody)
		}
		return response.SuccessResponse(http.StatusOK, responseBody, constantscore.ItemSuccessfullyObtained)
	}
	return response.SuccessResponse(http.StatusOK, responseBody, constantscore.DataNotFound)
}
