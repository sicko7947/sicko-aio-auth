package postgresql

import (
	"fmt"

	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
)

func RetrieveSuccess(key string) (results []*auth_service.StreamRetrieveSuccessItemsResponse_SuccessItem) {
	var successItems []*successItem

	query := fmt.Sprintf(`
	SELECT
		t1.*, t2.*
	FROM 
		"public"."successTable" t1
		LEFT JOIN "public"."productDetail" t2 ON t1."ProductId" = t2."ProductId"
	WHERE
		t1."KeyId" = '%s'
	ORDER BY
		t1."Timestamp" DESC;`, key)
	eg.SQL(query).Find(&successItems)

	for _, item := range successItems {
		results = append(results, &auth_service.StreamRetrieveSuccessItemsResponse_SuccessItem{
			Category:    item.Category,
			Region:      item.Region,
			ProductSku:  item.ProductSku,
			ProductName: item.ProductName,
			OrderNumber: item.OrderNumber,
			Email:       item.Email,
			Size:        item.Size,
			Timestamp:   item.Timestamp.Format("2006-01-02T15:04:05.000Z"),
			ImageUrl:    item.ImageUrl,
			RedirectUrl: item.RedirectUrl,
		})
	}
	return results
}
