package repository

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/diegocabrera89/ms-payment-core/dynamodbcore"
	"github.com/diegocabrera89/ms-payment-core/helpers"
	"github.com/diegocabrera89/ms-payment-core/logs"
	"github.com/diegocabrera89/ms-payment-merchant/constantsmicro"
	"github.com/diegocabrera89/ms-payment-merchant/models"
	"os"
)

// MerchantRepositoryImpl implements the MerchantRepository interface of the ms-payment-core package.
type MerchantRepositoryImpl struct {
	CoreRepository dynamodbcore.CoreRepository
}

// NewMerchantRepository create a new MerchantRepository instance.
func NewMerchantRepository() *MerchantRepositoryImpl {
	merchantTable := os.Getenv(constantsmicro.MerchantTable)
	region := os.Getenv(constantsmicro.Region)

	coreRepository, _ := dynamodbcore.NewDynamoDBRepository(merchantTable, region)

	return &MerchantRepositoryImpl{
		CoreRepository: coreRepository,
	}
}

// CreateMerchantRepository put item in DynamoDB.
func (r *MerchantRepositoryImpl) CreateMerchantRepository(ctx context.Context, request events.APIGatewayProxyRequest, merchant models.Merchant) (models.Merchant, error) {
	logs.LogTrackingInfo("CreateMerchantRepository", ctx, request)
	item, errorMarshallItem := helpers.MarshallItem(merchant)
	if errorMarshallItem != nil {
		logs.LogTrackingError("CreateMerchantRepository", "MarshallItem", ctx, request, errorMarshallItem)
		return models.Merchant{}, errorMarshallItem
	}

	errorPutItemCore := r.CoreRepository.PutItemCore(ctx, request, item)
	if errorPutItemCore != nil {
		return models.Merchant{}, errorPutItemCore
	}
	return merchant, nil
}

// ValidatePublicMerchantIdRepository get item in DynamoDB.
func (r *MerchantRepositoryImpl) ValidatePublicMerchantIdRepository(ctx context.Context, request events.APIGatewayProxyRequest, publicID string) (models.Merchant, error) {
	logs.LogTrackingInfo("ValidatePublicMerchantIdRepository", ctx, request)

	responseValidateMerchant, errorValidateMerchant := r.CoreRepository.GetItemByFieldCore(ctx, request, constantsmicro.PublicID, publicID, constantsmicro.PublicIDIndex, constantsmicro.Status, constantsmicro.StatusEnable)
	if errorValidateMerchant != nil {
		logs.LogTrackingError("ValidatePublicMerchantIdRepository", "GetItemByFieldCore", ctx, request, errorValidateMerchant)
		return models.Merchant{}, errorValidateMerchant
	}

	logs.LogTrackingInfoData("ValidatePublicMerchantIdRepository", responseValidateMerchant, ctx, request)

	// Verificar si hay al menos un elemento en la respuesta
	if responseValidateMerchant.Count == 0 {
		return models.Merchant{}, fmt.Errorf("No items found")
	}

	// Crear un slice de instancias de models.Merchant
	merchants := make([]models.Merchant, len(responseValidateMerchant.Items))
	logs.LogTrackingInfoData("ValidatePublicMerchantIdRepository merchants", merchants, ctx, request)

	// Deserializar los maps en las instancias de models.Merchant
	for i, item := range responseValidateMerchant.Items {
		var m models.Merchant
		err := helpers.UnmarshalMapToType(item, &m)
		if err != nil {
			logs.LogTrackingError("ValidatePublicMerchantIdRepository", "UnmarshalMapToType", ctx, request, err)
			return models.Merchant{}, err
		}
		merchants[i] = m
		logs.LogTrackingInfoData("ValidatePublicMerchantIdRepository merchants[i]", merchants[i], ctx, request)
	}

	logs.LogTrackingInfoData("ValidatePublicMerchantIdRepository merchants", merchants, ctx, request)
	// Si solo esperas un elemento, devuélvelo
	if len(merchants) == 1 {
		return merchants[0], nil
	}

	return models.Merchant{}, fmt.Errorf("More than one item found")
}

// ValidatePrivateMerchantIdRepository get item in DynamoDB.
func (r *MerchantRepositoryImpl) ValidatePrivateMerchantIdRepository(ctx context.Context, request events.APIGatewayProxyRequest, privateID string) (models.Merchant, error) {
	logs.LogTrackingInfo("ValidatePrivateMerchantIdRepository", ctx, request)

	responseValidateMerchant, errorValidateMerchant := r.CoreRepository.GetItemByFieldCore(ctx, request, constantsmicro.PrivateID, privateID, constantsmicro.PrivateIDIndex, constantsmicro.Status, constantsmicro.StatusEnable)
	if errorValidateMerchant != nil {
		logs.LogTrackingError("ValidatePrivateMerchantIdRepository", "GetItemByFieldCore", ctx, request, errorValidateMerchant)
		return models.Merchant{}, errorValidateMerchant
	}

	logs.LogTrackingInfoData("ValidatePrivateMerchantIdRepository", responseValidateMerchant, ctx, request)

	// Verificar si hay al menos un elemento en la respuesta
	if responseValidateMerchant.Count == 0 {
		return models.Merchant{}, fmt.Errorf("No items found")
	}

	// Crear un slice de instancias de models.Merchant
	merchants := make([]models.Merchant, len(responseValidateMerchant.Items))
	logs.LogTrackingInfoData("ValidatePrivateMerchantIdRepository merchants", merchants, ctx, request)

	// Deserializar los maps en las instancias de models.Merchant
	for i, item := range responseValidateMerchant.Items {
		var m models.Merchant
		err := helpers.UnmarshalMapToType(item, &m)
		if err != nil {
			logs.LogTrackingError("ValidatePrivateMerchantIdRepository", "UnmarshalMapToType", ctx, request, err)
			return models.Merchant{}, err
		}
		merchants[i] = m
		logs.LogTrackingInfoData("ValidatePrivateMerchantIdRepository merchants[i]", merchants[i], ctx, request)
	}

	logs.LogTrackingInfoData("ValidatePrivateMerchantIdRepository merchants", merchants, ctx, request)
	// Si solo esperas un elemento, devuélvelo
	if len(merchants) == 1 {
		return merchants[0], nil
	}

	return models.Merchant{}, fmt.Errorf("More than one item found")
}
