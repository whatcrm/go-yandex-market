package models

type ConfirmPricesRequest struct {
	OfferIds []ShopSku `json:"offerIds"`
}

type GetQuarantineOffersRequest struct {
	CardStatuses *[]OfferCardStatusType `json:"cardStatuses,omitempty"`

	CategoryIds *[]int32 `json:"categoryIds,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	VendorNames *[]string `json:"vendorNames,omitempty"`
}

type GetQuarantineOffersResponse struct {
	Result *GetQuarantineOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type PriceQuarantineVerdictDTO struct {
	Params []PriceQuarantineVerdictParameterDTO `json:"params"`

	Type *PriceQuarantineVerdictType `json:"type,omitempty"`
}

type PriceQuarantineVerdictParamNameType string

type PriceQuarantineVerdictParameterDTO struct {
	Name PriceQuarantineVerdictParamNameType `json:"name"`

	Value string `json:"value"`
}

type PriceQuarantineVerdictType string
