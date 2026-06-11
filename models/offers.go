package models

import "time"

const (
	AddOffersToArchiveErrorTypeOFFERHASSTOCKS AddOffersToArchiveErrorType = "OFFER_HAS_STOCKS"
	AddOffersToArchiveErrorTypeUNKNOWN        AddOffersToArchiveErrorType = "UNKNOWN"
)

const (
	DeleteOfferParameterTypeADDITIONALEXPENSES    DeleteOfferParameterType = "ADDITIONAL_EXPENSES"
	DeleteOfferParameterTypeADULT                 DeleteOfferParameterType = "ADULT"
	DeleteOfferParameterTypeAGE                   DeleteOfferParameterType = "AGE"
	DeleteOfferParameterTypeBARCODES              DeleteOfferParameterType = "BARCODES"
	DeleteOfferParameterTypeBOXCOUNT              DeleteOfferParameterType = "BOX_COUNT"
	DeleteOfferParameterTypeCERTIFICATES          DeleteOfferParameterType = "CERTIFICATES"
	DeleteOfferParameterTypeCOMMODITYCODES        DeleteOfferParameterType = "COMMODITY_CODES"
	DeleteOfferParameterTypeCONDITION             DeleteOfferParameterType = "CONDITION"
	DeleteOfferParameterTypeCUSTOMSCOMMODITYCODE  DeleteOfferParameterType = "CUSTOMS_COMMODITY_CODE"
	DeleteOfferParameterTypeDESCRIPTION           DeleteOfferParameterType = "DESCRIPTION"
	DeleteOfferParameterTypeDOWNLOADABLE          DeleteOfferParameterType = "DOWNLOADABLE"
	DeleteOfferParameterTypeGUARANTEEPERIOD       DeleteOfferParameterType = "GUARANTEE_PERIOD"
	DeleteOfferParameterTypeLIFETIME              DeleteOfferParameterType = "LIFE_TIME"
	DeleteOfferParameterTypeMANUALS               DeleteOfferParameterType = "MANUALS"
	DeleteOfferParameterTypeMANUFACTURERCOUNTRIES DeleteOfferParameterType = "MANUFACTURER_COUNTRIES"
	DeleteOfferParameterTypePARAMETERS            DeleteOfferParameterType = "PARAMETERS"
	DeleteOfferParameterTypePICTURES              DeleteOfferParameterType = "PICTURES"
	DeleteOfferParameterTypePURCHASEPRICE         DeleteOfferParameterType = "PURCHASE_PRICE"
	DeleteOfferParameterTypeSHELFLIFE             DeleteOfferParameterType = "SHELF_LIFE"
	DeleteOfferParameterTypeTAGS                  DeleteOfferParameterType = "TAGS"
	DeleteOfferParameterTypeTYPE                  DeleteOfferParameterType = "TYPE"
	DeleteOfferParameterTypeVENDORCODE            DeleteOfferParameterType = "VENDOR_CODE"
	DeleteOfferParameterTypeVIDEOS                DeleteOfferParameterType = "VIDEOS"
)

const (
	OfferConditionQualityTypeEXCELLENT    OfferConditionQualityType = "EXCELLENT"
	OfferConditionQualityTypeGOOD         OfferConditionQualityType = "GOOD"
	OfferConditionQualityTypeNOTSPECIFIED OfferConditionQualityType = "NOT_SPECIFIED"
	OfferConditionQualityTypePERFECT      OfferConditionQualityType = "PERFECT"
)

const (
	OfferConditionTypeNOTSPECIFIED   OfferConditionType = "NOT_SPECIFIED"
	OfferConditionTypePREOWNED       OfferConditionType = "PREOWNED"
	OfferConditionTypeREDUCTION      OfferConditionType = "REDUCTION"
	OfferConditionTypeREFURBISHED    OfferConditionType = "REFURBISHED"
	OfferConditionTypeRENOVATED      OfferConditionType = "RENOVATED"
	OfferConditionTypeSHOWCASESAMPLE OfferConditionType = "SHOWCASESAMPLE"
)

const (
	OfferContentErrorTypeINVALIDCATEGORY          OfferContentErrorType = "INVALID_CATEGORY"
	OfferContentErrorTypeINVALIDGROUPIDCHARACTERS OfferContentErrorType = "INVALID_GROUP_ID_CHARACTERS"
	OfferContentErrorTypeINVALIDGROUPIDLENGTH     OfferContentErrorType = "INVALID_GROUP_ID_LENGTH"
	OfferContentErrorTypeINVALIDUNITID            OfferContentErrorType = "INVALID_UNIT_ID"
	OfferContentErrorTypeNUMBERFORMAT             OfferContentErrorType = "NUMBER_FORMAT"
	OfferContentErrorTypeOFFERNOTFOUND            OfferContentErrorType = "OFFER_NOT_FOUND"
	OfferContentErrorTypeUNEXPECTEDBOOLEANVALUE   OfferContentErrorType = "UNEXPECTED_BOOLEAN_VALUE"
	OfferContentErrorTypeUNKNOWNCATEGORY          OfferContentErrorType = "UNKNOWN_CATEGORY"
	OfferContentErrorTypeUNKNOWNPARAMETER         OfferContentErrorType = "UNKNOWN_PARAMETER"
)

const (
	OfferTypeALCOHOL     OfferType = "ALCOHOL"
	OfferTypeARTISTTITLE OfferType = "ARTIST_TITLE"
	OfferTypeAUDIOBOOK   OfferType = "AUDIOBOOK"
	OfferTypeBOOK        OfferType = "BOOK"
	OfferTypeDEFAULT     OfferType = "DEFAULT"
	OfferTypeMEDICINE    OfferType = "MEDICINE"
	OfferTypeONDEMAND    OfferType = "ON_DEMAND"
)

const (
	PromoOfferParticipationStatusFilterTypeMANUALLYADDED    PromoOfferParticipationStatusFilterType = "MANUALLY_ADDED"
	PromoOfferParticipationStatusFilterTypeNOTMANUALLYADDED PromoOfferParticipationStatusFilterType = "NOT_MANUALLY_ADDED"
)

const (
	PromoOfferParticipationStatusMultiFilterTypeMANUALLYADDED    PromoOfferParticipationStatusMultiFilterType = "MANUALLY_ADDED"
	PromoOfferParticipationStatusMultiFilterTypeMINIMUMFORPROMOS PromoOfferParticipationStatusMultiFilterType = "MINIMUM_FOR_PROMOS"
	PromoOfferParticipationStatusMultiFilterTypeNOTMANUALLYADDED PromoOfferParticipationStatusMultiFilterType = "NOT_MANUALLY_ADDED"
	PromoOfferParticipationStatusMultiFilterTypeRENEWED          PromoOfferParticipationStatusMultiFilterType = "RENEWED"
	PromoOfferParticipationStatusMultiFilterTypeRENEWFAILED      PromoOfferParticipationStatusMultiFilterType = "RENEW_FAILED"
)

const (
	RejectedPromoOfferDeleteReasonTypeOFFERDOESNOTEXIST RejectedPromoOfferDeleteReasonType = "OFFER_DOES_NOT_EXIST"
)

const (
	RejectedPromoOfferUpdateReasonTypeDEADLINEFORFOCUSPROMOSEXCEEDED RejectedPromoOfferUpdateReasonType = "DEADLINE_FOR_FOCUS_PROMOS_EXCEEDED"
	RejectedPromoOfferUpdateReasonTypeEMPTYOLDPRICE                  RejectedPromoOfferUpdateReasonType = "EMPTY_OLD_PRICE"
	RejectedPromoOfferUpdateReasonTypeEMPTYPROMOPRICE                RejectedPromoOfferUpdateReasonType = "EMPTY_PROMO_PRICE"
	RejectedPromoOfferUpdateReasonTypeMAXPROMOPRICEEXCEEDED          RejectedPromoOfferUpdateReasonType = "MAX_PROMO_PRICE_EXCEEDED"
	RejectedPromoOfferUpdateReasonTypeOFFERDOESNOTEXIST              RejectedPromoOfferUpdateReasonType = "OFFER_DOES_NOT_EXIST"
	RejectedPromoOfferUpdateReasonTypeOFFERDUPLICATION               RejectedPromoOfferUpdateReasonType = "OFFER_DUPLICATION"
	RejectedPromoOfferUpdateReasonTypeOFFERNOTELIGIBLEFORPROMO       RejectedPromoOfferUpdateReasonType = "OFFER_NOT_ELIGIBLE_FOR_PROMO"
	RejectedPromoOfferUpdateReasonTypeOFFERPROMOSMAXBYTESIZEEXCEEDED RejectedPromoOfferUpdateReasonType = "OFFER_PROMOS_MAX_BYTE_SIZE_EXCEEDED"
	RejectedPromoOfferUpdateReasonTypeOLDPRICETOOBIG                 RejectedPromoOfferUpdateReasonType = "OLD_PRICE_TOO_BIG"
	RejectedPromoOfferUpdateReasonTypePRICETOOBIG                    RejectedPromoOfferUpdateReasonType = "PRICE_TOO_BIG"
	RejectedPromoOfferUpdateReasonTypePROMOPRICEBIGGERTHANMAX        RejectedPromoOfferUpdateReasonType = "PROMO_PRICE_BIGGER_THAN_MAX"
	RejectedPromoOfferUpdateReasonTypePROMOPRICESMALLERTHANMIN       RejectedPromoOfferUpdateReasonType = "PROMO_PRICE_SMALLER_THAN_MIN"
)

type AddOffersToArchiveDTO struct {
	NotArchivedOffers *[]AddOffersToArchiveErrorDTO `json:"notArchivedOffers,omitempty"`
}

type AddOffersToArchiveErrorDTO struct {
	Error AddOffersToArchiveErrorType `json:"error"`

	OfferId string `json:"offerId"`
}

type AddOffersToArchiveErrorType string

type BarcodeOfferInfoDTO struct {
	BarcodeCount int64 `json:"barcodeCount"`

	OfferId string `json:"offerId"`
}

type BaseOfferDTO struct {
	Adult *bool `json:"adult,omitempty"`

	Age *AgeDTO `json:"age,omitempty"`

	Barcodes *[]string `json:"barcodes,omitempty"`

	BoxCount *int32 `json:"boxCount,omitempty"`

	Category *string `json:"category,omitempty"`

	Certificates *[]string `json:"certificates,omitempty"`

	CommodityCodes *[]CommodityCodeDTO `json:"commodityCodes,omitempty"`

	Condition *OfferConditionDTO `json:"condition,omitempty"`

	CustomsCommodityCode *string `json:"customsCommodityCode,omitempty"`

	Description *string `json:"description,omitempty"`

	Downloadable *bool `json:"downloadable,omitempty"`

	GuaranteePeriod *TimePeriodDTO `json:"guaranteePeriod,omitempty"`

	LifeTime *TimePeriodDTO `json:"lifeTime,omitempty"`

	Manuals *[]OfferManualDTO `json:"manuals,omitempty"`

	ManufacturerCountries *[]string `json:"manufacturerCountries,omitempty"`

	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`

	Name *string `json:"name,omitempty"`

	OfferId string `json:"offerId"`

	Params *[]OfferParamDTO `json:"params,omitempty"`

	Pictures *[]Url `json:"pictures,omitempty"`

	ShelfLife *TimePeriodDTO `json:"shelfLife,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Type *OfferType `json:"type,omitempty"`

	Vendor *string `json:"vendor,omitempty"`

	VendorCode *string `json:"vendorCode,omitempty"`

	Videos *[]Url `json:"videos,omitempty"`

	WeightDimensions *OfferWeightDimensionsDTO `json:"weightDimensions,omitempty"`
}

type BaseOfferResponseDTO struct {
	Adult *bool `json:"adult,omitempty"`

	Age *AgeDTO `json:"age,omitempty"`

	Barcodes *[]string `json:"barcodes,omitempty"`

	BoxCount *int32 `json:"boxCount,omitempty"`

	Category *string `json:"category,omitempty"`

	Certificates *[]string `json:"certificates,omitempty"`

	CommodityCodes *[]CommodityCodeDTO `json:"commodityCodes,omitempty"`

	Condition *OfferConditionDTO `json:"condition,omitempty"`

	CustomsCommodityCode *string `json:"customsCommodityCode,omitempty"`

	Description *string `json:"description,omitempty"`

	Downloadable *bool `json:"downloadable,omitempty"`

	GuaranteePeriod *TimePeriodDTO `json:"guaranteePeriod,omitempty"`

	LifeTime *TimePeriodDTO `json:"lifeTime,omitempty"`

	Manuals *[]OfferManualDTO `json:"manuals,omitempty"`

	ManufacturerCountries *[]string `json:"manufacturerCountries,omitempty"`

	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`

	Name *string `json:"name,omitempty"`

	OfferId string `json:"offerId"`

	Params *[]OfferParamDTO `json:"params,omitempty"`

	Pictures *[]Url `json:"pictures,omitempty"`

	ShelfLife *TimePeriodDTO `json:"shelfLife,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Type *OfferType `json:"type,omitempty"`

	Vendor *string `json:"vendor,omitempty"`

	VendorCode *string `json:"vendorCode,omitempty"`

	Videos *[]Url `json:"videos,omitempty"`

	WeightDimensions *OfferWeightDimensionsDTO `json:"weightDimensions,omitempty"`
}

type CalculateTariffsOfferDTO struct {
	CategoryId int64 `json:"categoryId"`

	Height float32 `json:"height"`

	Length float32 `json:"length"`

	Price float32 `json:"price"`

	Quantity *int32 `json:"quantity,omitempty"`

	Weight float32 `json:"weight"`

	Width float32 `json:"width"`
}

type CalculateTariffsOfferInfoDTO struct {
	Offer CalculateTariffsOfferDTO `json:"offer"`

	Tariffs []CalculatedTariffDTO `json:"tariffs"`
}

type DeleteCampaignOffersRequest struct {
	OfferIds []ShopSku `json:"offerIds"`
}

type DeleteCampaignOffersResponse struct {
	Result *DeleteCampaignOffersDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type DeleteOfferParameterType string

type DeleteOffersDTO struct {
	NotDeletedOfferIds *[]ShopSku `json:"notDeletedOfferIds,omitempty"`
}

type DeleteOffersFromArchiveDTO struct {
	NotUnarchivedOfferIds *[]ShopSku `json:"notUnarchivedOfferIds,omitempty"`
}

type DeletePromoOffersResultDTO struct {
	RejectedOffers *[]RejectedPromoOfferDeleteDTO `json:"rejectedOffers,omitempty"`
}

type GenerateOfferBarcodesResultDTO struct {
	UnprocessedOfferIds []ShopSku `json:"unprocessedOfferIds"`
}

type GetCampaignOffersRequest struct {
	CategoryIds *[]int32 `json:"categoryIds,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	Statuses *[]OfferCampaignStatusType `json:"statuses,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	VendorNames *[]string `json:"vendorNames,omitempty"`
}

type GetCampaignOffersResponse struct {
	Result *GetCampaignOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetHiddenOffersResultDTO struct {
	HiddenOffers []HiddenOfferDTO `json:"hiddenOffers"`

	Paging *PackagingScrollingPagerDTO `json:"paging,omitempty"`
}

type GetOfferDTO struct {
	AdditionalExpenses *GetPriceDTO `json:"additionalExpenses,omitempty"`

	Adult *bool `json:"adult,omitempty"`

	Age *AgeDTO `json:"age,omitempty"`

	Archived *bool `json:"archived,omitempty"`

	Barcodes *[]string `json:"barcodes,omitempty"`

	BasicPrice *GetPriceWithDiscountDTO `json:"basicPrice,omitempty"`

	BoxCount *int32 `json:"boxCount,omitempty"`

	Campaigns *[]OfferCampaignStatusDTO `json:"campaigns,omitempty"`

	CardStatus *OfferCardStatusType `json:"cardStatus,omitempty"`

	Category *string `json:"category,omitempty"`

	Certificates *[]string `json:"certificates,omitempty"`

	CommodityCodes *[]CommodityCodeDTO `json:"commodityCodes,omitempty"`

	Condition *OfferConditionDTO `json:"condition,omitempty"`

	CustomsCommodityCode *string `json:"customsCommodityCode,omitempty"`

	Description *string `json:"description,omitempty"`

	Downloadable *bool `json:"downloadable,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	GuaranteePeriod *TimePeriodDTO `json:"guaranteePeriod,omitempty"`

	LifeTime *TimePeriodDTO `json:"lifeTime,omitempty"`

	Manuals *[]OfferManualDTO `json:"manuals,omitempty"`

	ManufacturerCountries *[]string `json:"manufacturerCountries,omitempty"`

	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`

	MediaFiles *OfferMediaFilesDTO `json:"mediaFiles,omitempty"`

	Name *string `json:"name,omitempty"`

	OfferId string `json:"offerId"`

	Params *[]OfferParamDTO `json:"params,omitempty"`

	Pictures *[]Url `json:"pictures,omitempty"`

	PurchasePrice *GetPriceDTO `json:"purchasePrice,omitempty"`

	SellingPrograms *[]OfferSellingProgramDTO `json:"sellingPrograms,omitempty"`

	ShelfLife *TimePeriodDTO `json:"shelfLife,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Type *OfferType `json:"type,omitempty"`

	Vendor *string `json:"vendor,omitempty"`

	VendorCode *string `json:"vendorCode,omitempty"`

	Videos *[]Url `json:"videos,omitempty"`

	WeightDimensions *OfferWeightDimensionsDTO `json:"weightDimensions,omitempty"`
}

type GetOfferRecommendationsRequest struct {
	CompetitivenessFilter *PriceCompetitivenessType `json:"competitivenessFilter,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`
}

type GetOfferRecommendationsResponse struct {
	Result *OfferRecommendationsResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPromoOfferDTO struct {
	AutoParticipatingDetails *PromoOfferAutoParticipatingDetailsDTO `json:"autoParticipatingDetails,omitempty"`

	OfferId string `json:"offerId"`

	Params PromoOfferParamsDTO `json:"params"`

	Status PromoOfferParticipationStatusType `json:"status"`
}

type GetPromoOffersResultDTO struct {
	Offers []GetPromoOfferDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetQuarantineOffersResultDTO struct {
	Offers []QuarantineOfferDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type HiddenOfferDTO struct {
	OfferId string `json:"offerId"`
}

type OfferCardContentStatusType string

type OfferCardDTO struct {
	AverageContentRating *int32 `json:"averageContentRating,omitempty"`

	CardStatus *OfferCardStatusType `json:"cardStatus,omitempty"`

	ContentRating *int32 `json:"contentRating,omitempty"`

	ContentRatingStatus *OfferCardContentStatusType `json:"contentRatingStatus,omitempty"`

	Errors *[]OfferErrorDTO `json:"errors,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	Mapping *GetMappingDTO `json:"mapping,omitempty"`

	OfferId string `json:"offerId"`

	ParameterValues *[]ParameterValueDTO `json:"parameterValues,omitempty"`

	Recommendations *[]OfferCardRecommendationDTO `json:"recommendations,omitempty"`

	Warnings *[]OfferErrorDTO `json:"warnings,omitempty"`
}

type OfferCardRecommendationDTO struct {
	Percent *int32 `json:"percent,omitempty"`

	RemainingRatingPoints *int32 `json:"remainingRatingPoints,omitempty"`

	Type OfferCardRecommendationType `json:"type"`
}

type OfferCardRecommendationType string

type OfferCardStatusType string

type OfferCardsContentStatusDTO struct {
	OfferCards []OfferCardDTO `json:"offerCards"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferConditionDTO struct {
	Quality *OfferConditionQualityType `json:"quality,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Type *OfferConditionType `json:"type,omitempty"`
}

type OfferConditionQualityType string

type OfferConditionType string

type OfferContentDTO struct {
	CategoryId int32 `json:"categoryId"`

	OfferId string `json:"offerId"`

	ParameterValues []ParameterValueDTO `json:"parameterValues"`
}

type OfferContentErrorDTO struct {
	Message string `json:"message"`

	ParameterId *int64 `json:"parameterId,omitempty"`

	Type OfferContentErrorType `json:"type"`
}

type OfferContentErrorType string

type OfferDefaultPriceDTO struct {
	CurrencyId *CurrencyType `json:"currencyId,omitempty"`

	DiscountBase *float32 `json:"discountBase,omitempty"`

	ExcludedFromBestsellers *bool `json:"excludedFromBestsellers,omitempty"`

	MinimumForBestseller *float32 `json:"minimumForBestseller,omitempty"`

	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	Value *float32 `json:"value,omitempty"`
}

type OfferDefaultPriceListResponseDTO struct {
	Offers []OfferDefaultPriceResponseDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferDefaultPriceResponseDTO struct {
	OfferId string `json:"offerId"`

	Price *OfferDefaultPriceDTO `json:"price,omitempty"`
}

type OfferErrorDTO struct {
	Comment *string `json:"comment,omitempty"`

	Message *string `json:"message,omitempty"`
}

type OfferForRecommendationDTO struct {
	Competitiveness *PriceCompetitivenessType `json:"competitiveness,omitempty"`

	OfferId *string `json:"offerId,omitempty"`

	Price *BasePriceDTO `json:"price,omitempty"`

	Shows *int64 `json:"shows,omitempty"`
}

type OfferManualDTO struct {
	Title *string `json:"title,omitempty"`
	Url   string  `json:"url"`
}

type OfferMediaFileDTO struct {
	Title *string `json:"title,omitempty"`

	UploadState *MediaFileUploadStateType `json:"uploadState,omitempty"`
	Url         *string                   `json:"url,omitempty"`
}

type OfferMediaFilesDTO struct {
	FirstVideoAsCover *bool `json:"firstVideoAsCover,omitempty"`

	Manuals *[]OfferMediaFileDTO `json:"manuals,omitempty"`

	Pictures *[]OfferMediaFileDTO `json:"pictures,omitempty"`

	Videos *[]OfferMediaFileDTO `json:"videos,omitempty"`
}

type OfferParamDTO struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

type OfferPriceByOfferIdsListResponseDTO struct {
	Offers []OfferPriceByOfferIdsResponseDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferPriceByOfferIdsResponseDTO struct {
	OfferId *string `json:"offerId,omitempty"`

	Price *PriceDTO `json:"price,omitempty"`

	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type OfferPriceDTO struct {
	OfferId *string `json:"offerId,omitempty"`

	Price *PriceDTO `json:"price,omitempty"`
}

type OfferPriceListResponseDTO struct {
	Offers []OfferPriceResponseDTO `json:"offers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

type OfferPriceResponseDTO struct {
	Id *string `json:"id,omitempty"`

	MarketSku *int64 `json:"marketSku,omitempty"`

	Price *PriceDTO `json:"price,omitempty"`

	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type OfferRecommendationDTO struct {
	Offer *OfferForRecommendationDTO `json:"offer,omitempty"`

	Recommendation *OfferRecommendationInfoDTO `json:"recommendation,omitempty"`
}

type OfferRecommendationInfoDTO struct {
	CompetitivenessThresholds *PriceCompetitivenessThresholdsDTO `json:"competitivenessThresholds,omitempty"`

	OfferId *string `json:"offerId,omitempty"`
}

type OfferRecommendationsResultDTO struct {
	OfferRecommendations []OfferRecommendationDTO `json:"offerRecommendations"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OfferSellingProgramDTO struct {
	SellingProgram SellingProgramType `json:"sellingProgram"`

	Status OfferSellingProgramStatusType `json:"status"`
}

type OfferSellingProgramStatusType string

type OfferType string

type OfferWeightDimensionsDTO struct {
	Height float32 `json:"height"`

	Length float32 `json:"length"`

	Weight float32 `json:"weight"`

	Width float32 `json:"width"`
}

type PromoOfferAutoParticipatingDetailsDTO struct {
	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`
}

type PromoOfferDiscountParamsDTO struct {
	MaxPromoPrice *int64 `json:"maxPromoPrice,omitempty"`

	Price *int64 `json:"price,omitempty"`

	PromoPrice *int64 `json:"promoPrice,omitempty"`
}

type PromoOfferParamsDTO struct {
	DiscountParams *PromoOfferDiscountParamsDTO `json:"discountParams,omitempty"`
}

type PromoOfferParticipationStatusFilterType string

type PromoOfferParticipationStatusMultiFilterType string

type PromoOfferParticipationStatusType string

type PromoOfferUpdateWarningCodeType string

type PromoOfferUpdateWarningDTO struct {
	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	Code PromoOfferUpdateWarningCodeType `json:"code"`
}

type QuarantineOfferDTO struct {
	CurrentPrice *BasePriceDTO `json:"currentPrice,omitempty"`

	LastValidPrice *BasePriceDTO `json:"lastValidPrice,omitempty"`

	OfferId *string `json:"offerId,omitempty"`

	Verdicts *[]PriceQuarantineVerdictDTO `json:"verdicts,omitempty"`
}

type RejectedPromoOfferDeleteDTO struct {
	OfferId string `json:"offerId"`

	Reason RejectedPromoOfferDeleteReasonType `json:"reason"`
}

type RejectedPromoOfferDeleteReasonType string

type RejectedPromoOfferUpdateDTO struct {
	OfferId string `json:"offerId"`

	Reason RejectedPromoOfferUpdateReasonType `json:"reason"`
}

type RejectedPromoOfferUpdateReasonType string

type UpdateCampaignOffersRequest struct {
	Offers []UpdateCampaignOfferDTO `json:"offers"`
}

type UpdateOfferContentResultDTO struct {
	Errors *[]OfferContentErrorDTO `json:"errors,omitempty"`

	OfferId string `json:"offerId"`

	Warnings *[]OfferContentErrorDTO `json:"warnings,omitempty"`
}

type UpdateOfferDTO struct {
	AdditionalExpenses *BasePriceDTO `json:"additionalExpenses,omitempty"`

	Adult *bool `json:"adult,omitempty"`

	Age *AgeDTO `json:"age,omitempty"`

	Barcodes *[]string `json:"barcodes,omitempty"`

	BasicPrice *PriceWithDiscountDTO `json:"basicPrice,omitempty"`

	BoxCount *int32 `json:"boxCount,omitempty"`

	Category *string `json:"category,omitempty"`

	Certificates *[]string `json:"certificates,omitempty"`

	CommodityCodes *[]CommodityCodeDTO `json:"commodityCodes,omitempty"`

	Condition *OfferConditionDTO `json:"condition,omitempty"`

	CustomsCommodityCode *string `json:"customsCommodityCode,omitempty"`

	DeleteParameters *[]DeleteOfferParameterType `json:"deleteParameters,omitempty"`

	Description *string `json:"description,omitempty"`

	Downloadable *bool `json:"downloadable,omitempty"`

	FirstVideoAsCover *bool `json:"firstVideoAsCover,omitempty"`

	GuaranteePeriod *TimePeriodDTO `json:"guaranteePeriod,omitempty"`

	LifeTime *TimePeriodDTO `json:"lifeTime,omitempty"`

	Manuals *[]OfferManualDTO `json:"manuals,omitempty"`

	ManufacturerCountries *[]string `json:"manufacturerCountries,omitempty"`

	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`

	Name *string `json:"name,omitempty"`

	OfferId string `json:"offerId"`

	ParameterValues *[]ParameterValueDTO `json:"parameterValues,omitempty"`

	Params *[]OfferParamDTO `json:"params,omitempty"`

	Pictures *[]Url `json:"pictures,omitempty"`

	PurchasePrice *BasePriceDTO `json:"purchasePrice,omitempty"`

	ShelfLife *TimePeriodDTO `json:"shelfLife,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Type *OfferType `json:"type,omitempty"`

	Vendor *string `json:"vendor,omitempty"`

	VendorCode *string `json:"vendorCode,omitempty"`

	Videos *[]Url `json:"videos,omitempty"`

	WeightDimensions *OfferWeightDimensionsDTO `json:"weightDimensions,omitempty"`
}

type UpdatePromoOfferDTO struct {
	OfferId string `json:"offerId"`

	Params *UpdatePromoOfferParamsDTO `json:"params,omitempty"`
}

type UpdatePromoOfferDiscountParamsDTO struct {
	Price *int64 `json:"price,omitempty"`

	PromoPrice *int64 `json:"promoPrice,omitempty"`
}

type UpdatePromoOfferParamsDTO struct {
	DiscountParams *UpdatePromoOfferDiscountParamsDTO `json:"discountParams,omitempty"`
}

type UpdatePromoOffersResultDTO struct {
	RejectedOffers *[]RejectedPromoOfferUpdateDTO `json:"rejectedOffers,omitempty"`

	WarningOffers *[]WarningPromoOfferUpdateDTO `json:"warningOffers,omitempty"`
}

type WarehouseOfferDTO struct {
	OfferId string `json:"offerId"`

	Stocks []WarehouseStockDTO `json:"stocks"`

	TurnoverSummary *TurnoverDTO `json:"turnoverSummary,omitempty"`

	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type WarehouseOffersDTO struct {
	Offers []WarehouseOfferDTO `json:"offers"`

	WarehouseId int64 `json:"warehouseId"`
}

type WarningPromoOfferUpdateDTO struct {
	OfferId string `json:"offerId"`

	Warnings []PromoOfferUpdateWarningDTO `json:"warnings"`
}

type GetOfferCardsContentStatusParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetOfferRecommendationsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetBusinessQuarantineOffersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetPromoOffersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetHiddenOffersParams struct {
	OfferId *[]ShopSku `form:"offer_id,omitempty" json:"offer_id,omitempty"`

	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetPricesByOfferIdsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GenerateOfferBarcodesJSONRequestBody = GenerateOfferBarcodesRequest

type GetOfferCardsContentStatusJSONRequestBody = GetOfferCardsContentStatusRequest

type UpdateOfferContentJSONRequestBody = UpdateOfferContentRequest

type AddOffersToArchiveJSONRequestBody = AddOffersToArchiveRequest

type DeleteOffersJSONRequestBody = DeleteOffersRequest

type DeleteOffersFromArchiveJSONRequestBody = DeleteOffersFromArchiveRequest

type GetOfferRecommendationsJSONRequestBody = GetOfferRecommendationsRequest

type GetBusinessQuarantineOffersJSONRequestBody = GetQuarantineOffersRequest

type GetPromoOffersJSONRequestBody = GetPromoOffersRequest

type DeletePromoOffersJSONRequestBody = DeletePromoOffersRequest

type UpdatePromoOffersJSONRequestBody = UpdatePromoOffersRequest

type AddHiddenOffersJSONRequestBody = AddHiddenOffersRequest

type DeleteHiddenOffersJSONRequestBody = DeleteHiddenOffersRequest

type GetPricesByOfferIdsJSONRequestBody = GetPricesByOfferIdsRequest
