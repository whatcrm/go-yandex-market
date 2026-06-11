package models

type AddOffersToArchiveRequest struct {
	OfferIds []ShopSku `json:"offerIds"`
}

type AddOffersToArchiveResponse struct {
	Result *AddOffersToArchiveDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type DeleteOffersFromArchiveRequest struct {
	OfferIds []ShopSku `json:"offerIds"`
}

type DeleteOffersFromArchiveResponse struct {
	Result *DeleteOffersFromArchiveDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type DeleteOffersRequest struct {
	OfferIds []ShopSku `json:"offerIds"`
}

type DeleteOffersResponse struct {
	Result *DeleteOffersDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GenerateOfferBarcodesRequest struct {
	OfferIds []ShopSku `json:"offerIds"`

	SkipIfExists *bool `json:"skipIfExists,omitempty"`
}

type GenerateOfferBarcodesResponse struct {
	Result *GenerateOfferBarcodesResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOfferMappingDTO struct {
	Mapping *GetMappingDTO `json:"mapping,omitempty"`

	Offer *GetOfferDTO `json:"offer,omitempty"`

	ShowcaseUrls *[]ShowcaseUrlDTO `json:"showcaseUrls,omitempty"`
}

type GetOfferMappingsRequest struct {
	Archived *bool `json:"archived,omitempty"`

	CardStatuses *[]OfferCardStatusType `json:"cardStatuses,omitempty"`

	CategoryIds *[]int32 `json:"categoryIds,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	VendorNames *[]string `json:"vendorNames,omitempty"`
}

type GetOfferMappingsResponse struct {
	Result *GetOfferMappingsResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOfferMappingsResultDTO struct {
	OfferMappings []GetOfferMappingDTO `json:"offerMappings"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferMappingErrorDTO struct {
	Message string `json:"message"`

	ParameterId *int64 `json:"parameterId,omitempty"`

	Type OfferMappingErrorType `json:"type"`
}

type OfferMappingErrorType string

type UpdateBusinessOfferPriceDTO struct {
	OfferId string `json:"offerId"`

	Price UpdateBusinessPricesDTO `json:"price"`
}

type UpdateOfferMappingDTO struct {
	Mapping *UpdateMappingDTO `json:"mapping,omitempty"`

	Offer UpdateOfferDTO `json:"offer"`
}

type UpdateOfferMappingResultDTO struct {
	Errors *[]OfferMappingErrorDTO `json:"errors,omitempty"`

	OfferId string `json:"offerId"`

	Warnings *[]OfferMappingErrorDTO `json:"warnings,omitempty"`
}

type UpdateOfferMappingsRequest struct {
	OfferMappings []UpdateOfferMappingDTO `json:"offerMappings"`

	OnlyPartnerMediaContent *bool `json:"onlyPartnerMediaContent,omitempty"`
}

type UpdateOfferMappingsResponse struct {
	Results *[]UpdateOfferMappingResultDTO `json:"results,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOfferMappingsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	Language *CatalogLanguageType `form:"language,omitempty" json:"language,omitempty"`
}

type UpdateOfferMappingsParams struct {
	Language *CatalogLanguageType `form:"language,omitempty" json:"language,omitempty"`
}

type GetOfferMappingsJSONRequestBody = GetOfferMappingsRequest

type UpdateOfferMappingsJSONRequestBody = UpdateOfferMappingsRequest
