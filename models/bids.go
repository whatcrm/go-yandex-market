package models

type BidRecommendationItemDTO struct {
	Benefits *[]BenefitType `json:"benefits,omitempty"`

	Bid int32 `json:"bid"`

	ShowPercent int64 `json:"showPercent"`
}

type GetBidsInfoRequest struct {
	Skus *[]ShopSku `json:"skus,omitempty"`
}

type GetBidsInfoResponse struct {
	Result *GetBidsInfoResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetBidsInfoResponseDTO struct {
	Bids []SkuBidItemDTO `json:"bids"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetBidsRecommendationsRequest struct {
	Skus []ShopSku `json:"skus"`
}

type GetBidsRecommendationsResponse struct {
	Result *GetBidsRecommendationsResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetBidsRecommendationsResponseDTO struct {
	Recommendations []SkuBidRecommendationItemDTO `json:"recommendations"`
}

type PutSkuBidsRequest struct {
	Bids []SkuBidItemDTO `json:"bids"`
}

type SkuBidItemDTO struct {
	Bid int32 `json:"bid"`

	Sku string `json:"sku"`
}

type SkuBidRecommendationItemDTO struct {
	Bid int32 `json:"bid"`

	BidRecommendations *[]BidRecommendationItemDTO `json:"bidRecommendations,omitempty"`

	PriceRecommendations *[]PriceRecommendationItemDTO `json:"priceRecommendations,omitempty"`

	Sku string `json:"sku"`
}

type GetBidsInfoForBusinessParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type PutBidsForBusinessJSONRequestBody = PutSkuBidsRequest

type GetBidsInfoForBusinessJSONRequestBody = GetBidsInfoRequest

type GetBidsRecommendationsJSONRequestBody = GetBidsRecommendationsRequest
