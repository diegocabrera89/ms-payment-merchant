package utils

import (
	"github.com/diegocabrera89/ms-payment-merchant/constantsmicro"
	"github.com/diegocabrera89/ms-payment-merchant/models"
	"github.com/google/uuid"
	"time"
)

// BuildCreateMerchant build merchant object.
func BuildCreateMerchant(merchant *models.Merchant) {
	merchant.MerchantID = uuid.New().String()    // Generate a unique ID for the client
	merchant.CreatedAt = time.Now().UTC().Unix() //Date in UTC
	merchant.StatusMerchant = constantsmicro.StatusMerchantEnable
}
