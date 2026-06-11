package models

type GetGoodsStatsRequest struct {
	ShopSkus []ShopSku `json:"shopSkus"`
}

type GetGoodsStatsResponse struct {
	Result *GoodsStatsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GoodsStatsDTO struct {
	ShopSkus []GoodsStatsGoodsDTO `json:"shopSkus"`
}

type GoodsStatsGoodsDTO struct {
	CategoryId *int64 `json:"categoryId,omitempty"`

	CategoryName *string `json:"categoryName,omitempty"`

	MarketSku *int64 `json:"marketSku,omitempty"`

	Name *string `json:"name,omitempty"`

	Pictures *[]Url `json:"pictures,omitempty"`

	Price *float32 `json:"price,omitempty"`

	ShopSku *string `json:"shopSku,omitempty"`

	Tariffs *[]TariffDTO `json:"tariffs,omitempty"`

	Warehouses *[]GoodsStatsWarehouseDTO `json:"warehouses,omitempty"`

	WeightDimensions *GoodsStatsWeightDimensionsDTO `json:"weightDimensions,omitempty"`
}

type GoodsStatsWeightDimensionsDTO struct {
	Height *float32 `json:"height,omitempty"`

	Length *float32 `json:"length,omitempty"`

	Weight *float32 `json:"weight,omitempty"`

	Width *float32 `json:"width,omitempty"`
}

type GetGoodsStatsJSONRequestBody = GetGoodsStatsRequest
