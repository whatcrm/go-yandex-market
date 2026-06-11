package models

const (
	OfferCampaignStatusTypeARCHIVED              OfferCampaignStatusType = "ARCHIVED"
	OfferCampaignStatusTypeCHECKING              OfferCampaignStatusType = "CHECKING"
	OfferCampaignStatusTypeCREATINGCARD          OfferCampaignStatusType = "CREATING_CARD"
	OfferCampaignStatusTypeDISABLEDAUTOMATICALLY OfferCampaignStatusType = "DISABLED_AUTOMATICALLY"
	OfferCampaignStatusTypeDISABLEDBYPARTNER     OfferCampaignStatusType = "DISABLED_BY_PARTNER"
	OfferCampaignStatusTypeNOCARD                OfferCampaignStatusType = "NO_CARD"
	OfferCampaignStatusTypeNOSTOCKS              OfferCampaignStatusType = "NO_STOCKS"
	OfferCampaignStatusTypePUBLISHED             OfferCampaignStatusType = "PUBLISHED"
	OfferCampaignStatusTypeREJECTEDBYMARKET      OfferCampaignStatusType = "REJECTED_BY_MARKET"
)

type BaseCampaignOfferDTO struct {
	Available *bool `json:"available,omitempty"`

	OfferId string `json:"offerId"`

	Quantum *QuantumDTO `json:"quantum,omitempty"`
}

type CampaignDTO struct {
	ApiAvailability *ApiAvailabilityStatusType `json:"apiAvailability,omitempty"`

	Business *BusinessDTO `json:"business,omitempty"`

	ClientId *int64 `json:"clientId,omitempty"`

	Domain *string `json:"domain,omitempty"`

	Id *int64 `json:"id,omitempty"`

	PlacementType *PlacementType `json:"placementType,omitempty"`
}

type CampaignId = int64

type CampaignQualityRatingDTO struct {
	CampaignId int64 `json:"campaignId"`

	Ratings []QualityRatingDTO `json:"ratings"`
}

type CampaignSettingsDTO struct {
	CountryRegion *int64 `json:"countryRegion,omitempty"`

	LocalRegion *CampaignSettingsLocalRegionDTO `json:"localRegion,omitempty"`

	ShopName *string `json:"shopName,omitempty"`

	ShowInContext *bool `json:"showInContext,omitempty"`

	ShowInPremium *bool `json:"showInPremium,omitempty"`

	UseOpenStat *bool `json:"useOpenStat,omitempty"`
}

type CampaignSettingsDeliveryDTO struct {
	Schedule *CampaignSettingsScheduleDTO `json:"schedule,omitempty"`
}

type CampaignSettingsLocalRegionDTO struct {
	Delivery *CampaignSettingsDeliveryDTO `json:"delivery,omitempty"`

	DeliveryOptionsSource *CampaignSettingsScheduleSourceType `json:"deliveryOptionsSource,omitempty"`

	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *RegionType `json:"type,omitempty"`
}

type CampaignSettingsScheduleDTO struct {
	AvailableOnHolidays *bool `json:"availableOnHolidays,omitempty"`

	CustomHolidays []DateDdMmYyyy `json:"customHolidays"`

	CustomWorkingDays []DateDdMmYyyy `json:"customWorkingDays"`

	Period *CampaignSettingsTimePeriodDTO `json:"period,omitempty"`

	TotalHolidays []DateDdMmYyyy `json:"totalHolidays"`

	WeeklyHolidays []int32 `json:"weeklyHolidays"`
}

type CampaignSettingsScheduleSourceType string

type CampaignSettingsTimePeriodDTO struct {
	FromDate *string `json:"fromDate,omitempty"`

	ToDate *string `json:"toDate,omitempty"`
}

type CampaignsQualityRatingDTO struct {
	CampaignRatings []CampaignQualityRatingDTO `json:"campaignRatings"`
}

type DeleteCampaignOffersDTO struct {
	NotDeletedOfferIds *[]ShopSku `json:"notDeletedOfferIds,omitempty"`
}

type GetCampaignOfferDTO struct {
	Available *bool `json:"available,omitempty"`

	BasicPrice *GetPriceWithDiscountDTO `json:"basicPrice,omitempty"`

	CampaignPrice *GetPriceWithVatDTO `json:"campaignPrice,omitempty"`

	Errors *[]OfferErrorDTO `json:"errors,omitempty"`

	OfferId string `json:"offerId"`

	Quantum *QuantumDTO `json:"quantum,omitempty"`

	Status *OfferCampaignStatusType `json:"status,omitempty"`

	Warnings *[]OfferErrorDTO `json:"warnings,omitempty"`
}

type GetCampaignOffersResultDTO struct {
	Offers []GetCampaignOfferDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetCampaignResponse struct {
	Campaign *CampaignDTO `json:"campaign,omitempty"`
}

type GetCampaignSettingsResponse struct {
	Settings *CampaignSettingsDTO `json:"settings,omitempty"`
}

type GetCampaignsResponse struct {
	Campaigns []CampaignDTO `json:"campaigns"`

	Pager *FlippingPagerDTO `json:"pager,omitempty"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferCampaignStatusDTO struct {
	CampaignId int64 `json:"campaignId"`

	Status OfferCampaignStatusType `json:"status"`
}

type OfferCampaignStatusType string

type UpdateCampaignOfferDTO struct {
	Available *bool `json:"available,omitempty"`

	OfferId string `json:"offerId"`

	Quantum *QuantumDTO `json:"quantum,omitempty"`

	Vat *VatType `json:"vat,omitempty"`
}

type GetCampaignsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	PageSize *int32 `form:"pageSize,omitempty" json:"pageSize,omitempty"`
}

type GetCampaignOffersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetCampaignQuarantineOffersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type PutBidsForCampaignJSONRequestBody = PutSkuBidsRequest

type GetCampaignOffersJSONRequestBody = GetCampaignOffersRequest

type DeleteCampaignOffersJSONRequestBody = DeleteCampaignOffersRequest

type UpdateCampaignOffersJSONRequestBody = UpdateCampaignOffersRequest

type GetCampaignQuarantineOffersJSONRequestBody = GetQuarantineOffersRequest

type ConfirmCampaignPricesJSONRequestBody = ConfirmPricesRequest
