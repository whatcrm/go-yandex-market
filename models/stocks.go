package models

import "time"

type GetWarehouseStocksRequest struct {
	Archived *bool `json:"archived,omitempty"`

	HasStocks *bool `json:"hasStocks,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	StocksWarehouseId *int64 `json:"stocksWarehouseId,omitempty"`

	WithTurnover *bool `json:"withTurnover,omitempty"`
}

type GetWarehouseStocksResponse struct {
	Result *GetWarehouseStocksDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type UpdateStockDTO struct {
	Items []UpdateStockItemDTO `json:"items"`

	Sku string `json:"sku"`
}

type UpdateStockItemDTO struct {
	Count int64 `json:"count"`

	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type UpdateStocksRequest struct {
	Skus []UpdateStockDTO `json:"skus"`
}

type GetStocksParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetStocksJSONRequestBody = GetWarehouseStocksRequest

type UpdateStocksJSONRequestBody = UpdateStocksRequest
