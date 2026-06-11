package models

import "time"

const (
	PriceCompetitivenessTypeAVERAGE PriceCompetitivenessType = "AVERAGE"
	PriceCompetitivenessTypeLOW     PriceCompetitivenessType = "LOW"
	PriceCompetitivenessTypeOPTIMAL PriceCompetitivenessType = "OPTIMAL"
)

type ApiLockedErrorResponse = ApiErrorResponse

type BasePriceDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	Value float32 `json:"value"`
}

type DeliveryOptionPriceDTO = CurrencyValueDTO

type DeliveryPriceDTO struct {
	Payment *CurrencyValueDTO `json:"payment,omitempty"`

	Subsidy *CurrencyValueDTO `json:"subsidy,omitempty"`

	Vat *OrderVatType `json:"vat,omitempty"`
}

type EmptyApiResponse = ApiResponse

type GetDefaultPricesRequest struct {
	Archived *bool `json:"archived,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`
}

type GetDefaultPricesResponse struct {
	Result *OfferDefaultPriceListResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPriceDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	UpdatedAt time.Time `json:"updatedAt"`

	Value float32 `json:"value"`
}

type GetPriceWithDiscountDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	DiscountBase *DiscountBase `json:"discountBase,omitempty"`

	UpdatedAt time.Time `json:"updatedAt"`

	Value float32 `json:"value"`
}

type GetPriceWithVatDTO struct {
	CurrencyId *CurrencyType `json:"currencyId,omitempty"`

	DiscountBase *float32 `json:"discountBase,omitempty"`

	UpdatedAt time.Time `json:"updatedAt"`

	Value *float32 `json:"value,omitempty"`

	Vat *int32 `json:"vat,omitempty"`
}

type GetPricesByOfferIdsRequest struct {
	OfferIds *[]ShopSku `json:"offerIds,omitempty"`
}

type GetPricesByOfferIdsResponse struct {
	Result *OfferPriceByOfferIdsListResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPricesResponse struct {
	Result *OfferPriceListResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type ItemPriceDTO struct {
	Cashback *CurrencyValueDTO `json:"cashback,omitempty"`

	Payment *CurrencyValueDTO `json:"payment,omitempty"`

	Subsidy *CurrencyValueDTO `json:"subsidy,omitempty"`

	Vat *OrderVatType `json:"vat,omitempty"`
}

type PriceCompetitivenessThresholdsDTO struct {
	AveragePrice *BasePriceDTO `json:"averagePrice,omitempty"`

	OptimalPrice *BasePriceDTO `json:"optimalPrice,omitempty"`
}

type PriceCompetitivenessType string

type PriceDTO struct {
	CurrencyId *CurrencyType `json:"currencyId,omitempty"`

	DiscountBase *float32 `json:"discountBase,omitempty"`

	Value *float32 `json:"value,omitempty"`

	Vat *int32 `json:"vat,omitempty"`
}

type PriceRecommendationItemDTO struct {
	CampaignId int64 `json:"campaignId"`

	Price float32 `json:"price"`
}

type PriceWithDiscountDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	DiscountBase *DiscountBase `json:"discountBase,omitempty"`

	Value float32 `json:"value"`
}

type UpdateBusinessPricesDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	DiscountBase *DiscountBase `json:"discountBase,omitempty"`

	MinimumForBestseller *float32 `json:"minimumForBestseller,omitempty"`

	Value float32 `json:"value"`
}

type UpdateBusinessPricesRequest struct {
	Offers []UpdateBusinessOfferPriceDTO `json:"offers"`
}

type UpdatePricesRequest struct {
	Offers []OfferPriceDTO `json:"offers"`
}

type GetDefaultPricesParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetPricesParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	Archived *bool `form:"archived,omitempty" json:"archived,omitempty"`
}

type GetDefaultPricesJSONRequestBody = GetDefaultPricesRequest

type UpdateBusinessPricesJSONRequestBody = UpdateBusinessPricesRequest

type ConfirmBusinessPricesJSONRequestBody = ConfirmPricesRequest

type UpdatePricesJSONRequestBody = UpdatePricesRequest
