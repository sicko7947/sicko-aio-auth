package postgresql

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func InsertSuccess(successItem *models.SuccessItem) (err error) {

	t, err := time.Parse("2006-01-02T15:04:05.000Z", successItem.Setup.GetTimestamp())
	if err != nil {
		fmt.Println(err)
	}
	productId := uuid.NewString()
	successEntry := &successTable{
		KeyId:           successItem.KeyId,
		EntryId:         uuid.NewString(),
		ProductId:       productId,
		Category:        successItem.Setup.Category,
		Region:          successItem.Setup.GetRegion(),
		TaskType:        successItem.Setup.GetTaskType(),
		MonitorMode:     successItem.Setup.GetMonitorMode(),
		Timestamp:       t,
		UsePsychoCookie: successItem.Setup.GetUsePsychoCookie(),
	}

	_, err = eg.Insert(successEntry)
	if err != nil {
		return err
	}

	productEntry := &productDetail{
		ProductId:          productId,
		MerchGroup:         successItem.Product.GetMerchGroup(),
		ProductSku:         successItem.Product.GetProductSku(),
		ProductName:        successItem.Product.GetProductName(),
		ProductDescription: successItem.Product.GetProductDescription(),
		Size:               successItem.Product.GetSize(),
		Price:              successItem.Product.GetPrice(),
		Quantity:           successItem.Product.GetQuantity(),
		OrderNumber:        successItem.Product.GetOrderNumber(),
		ProfileName:        successItem.Product.GetProfileName(),
		Email:              successItem.Product.GetEmail(),
		Account:            successItem.Product.GetAccount(),
		GiftCards:          successItem.Product.GetGiftCards(),
		DiscountCode:       successItem.Product.GetDiscountCode(),
		ImageUrl:           successItem.Product.GetImageUrl(),
		RedirectUrl:        successItem.Product.GetRedirectUrl(),
	}
	_, err = eg.Insert(productEntry)
	if err != nil {
		return err
	}

	return nil
}
