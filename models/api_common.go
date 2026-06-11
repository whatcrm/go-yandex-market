package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (e AddOffersToArchiveErrorType) Valid() bool {
	switch e {
	case AddOffersToArchiveErrorTypeOFFERHASSTOCKS:
		return true
	case AddOffersToArchiveErrorTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e AffectedOrderQualityRatingComponentType) Valid() bool {
	switch e {
	case AffectedOrderQualityRatingComponentTypeDBSCANCELLATIONRATE:
		return true
	case AffectedOrderQualityRatingComponentTypeDBSLATEDELIVERYRATE:
		return true
	case AffectedOrderQualityRatingComponentTypeFBSCANCELLATIONRATE:
		return true
	case AffectedOrderQualityRatingComponentTypeFBSLATESHIPRATE:
		return true
	default:
		return false
	}
}

const (
	AgeUnitTypeMONTH AgeUnitType = "MONTH"
	AgeUnitTypeYEAR  AgeUnitType = "YEAR"
)

func (e AgeUnitType) Valid() bool {
	switch e {
	case AgeUnitTypeMONTH:
		return true
	case AgeUnitTypeYEAR:
		return true
	default:
		return false
	}
}

const (
	ApiAvailabilityStatusTypeAVAILABLE                  ApiAvailabilityStatusType = "AVAILABLE"
	ApiAvailabilityStatusTypeDISABLEDBYINACTIVITY       ApiAvailabilityStatusType = "DISABLED_BY_INACTIVITY"
	ApiAvailabilityStatusTypeDISABLEDBYNOACTIVECONTRACT ApiAvailabilityStatusType = "DISABLED_BY_NO_ACTIVE_CONTRACT"
	ApiAvailabilityStatusTypeDISABLEDBYNOPLACEMENTTYPE  ApiAvailabilityStatusType = "DISABLED_BY_NO_PLACEMENT_TYPE"
	ApiAvailabilityStatusTypeMANUALLYDISABLED           ApiAvailabilityStatusType = "MANUALLY_DISABLED"
)

func (e ApiAvailabilityStatusType) Valid() bool {
	switch e {
	case ApiAvailabilityStatusTypeAVAILABLE:
		return true
	case ApiAvailabilityStatusTypeDISABLEDBYINACTIVITY:
		return true
	case ApiAvailabilityStatusTypeDISABLEDBYNOACTIVECONTRACT:
		return true
	case ApiAvailabilityStatusTypeDISABLEDBYNOPLACEMENTTYPE:
		return true
	case ApiAvailabilityStatusTypeMANUALLYDISABLED:
		return true
	default:
		return false
	}
}

const (
	ALLMETHODS                          ApiKeyScopeType = "ALL_METHODS"
	ALLMETHODSREADONLY                  ApiKeyScopeType = "ALL_METHODS_READ_ONLY"
	COMMUNICATION                       ApiKeyScopeType = "COMMUNICATION"
	FINANCEANDACCOUNTING                ApiKeyScopeType = "FINANCE_AND_ACCOUNTING"
	INVENTORYANDORDERPROCESSING         ApiKeyScopeType = "INVENTORY_AND_ORDER_PROCESSING"
	INVENTORYANDORDERPROCESSINGREADONLY ApiKeyScopeType = "INVENTORY_AND_ORDER_PROCESSING_READ_ONLY"
	OFFERSANDCARDSMANAGEMENT            ApiKeyScopeType = "OFFERS_AND_CARDS_MANAGEMENT"
	OFFERSANDCARDSMANAGEMENTREADONLY    ApiKeyScopeType = "OFFERS_AND_CARDS_MANAGEMENT_READ_ONLY"
	PRICING                             ApiKeyScopeType = "PRICING"
	PRICINGREADONLY                     ApiKeyScopeType = "PRICING_READ_ONLY"
	PROMOTION                           ApiKeyScopeType = "PROMOTION"
	PROMOTIONREADONLY                   ApiKeyScopeType = "PROMOTION_READ_ONLY"
	SETTINGSMANAGEMENT                  ApiKeyScopeType = "SETTINGS_MANAGEMENT"
	SUPPLIESMANAGEMENTREADONLY          ApiKeyScopeType = "SUPPLIES_MANAGEMENT_READ_ONLY"
)

func (e ApiKeyScopeType) Valid() bool {
	switch e {
	case ALLMETHODS:
		return true
	case ALLMETHODSREADONLY:
		return true
	case COMMUNICATION:
		return true
	case FINANCEANDACCOUNTING:
		return true
	case INVENTORYANDORDERPROCESSING:
		return true
	case INVENTORYANDORDERPROCESSINGREADONLY:
		return true
	case OFFERSANDCARDSMANAGEMENT:
		return true
	case OFFERSANDCARDSMANAGEMENTREADONLY:
		return true
	case PRICING:
		return true
	case PRICINGREADONLY:
		return true
	case PROMOTION:
		return true
	case PROMOTIONREADONLY:
		return true
	case SETTINGSMANAGEMENT:
		return true
	case SUPPLIESMANAGEMENTREADONLY:
		return true
	default:
		return false
	}
}

const (
	ApiResponseStatusTypeERROR ApiResponseStatusType = "ERROR"
	ApiResponseStatusTypeOK    ApiResponseStatusType = "OK"
)

func (e ApiResponseStatusType) Valid() bool {
	switch e {
	case ApiResponseStatusTypeERROR:
		return true
	case ApiResponseStatusTypeOK:
		return true
	default:
		return false
	}
}

const (
	F3020       BarcodeFormatType = "F_30_20"
	F4325       BarcodeFormatType = "F_43_25"
	F4325SINGLE BarcodeFormatType = "F_43_25_SINGLE"
	F5840       BarcodeFormatType = "F_58_40"
)

func (e BarcodeFormatType) Valid() bool {
	switch e {
	case F3020:
		return true
	case F4325:
		return true
	case F4325SINGLE:
		return true
	case F5840:
		return true
	default:
		return false
	}
}

const (
	BESTS           BenefitType = "BESTS"
	MARKETSUBSIDY10 BenefitType = "MARKET_SUBSIDY_10"
	MARKETSUBSIDY14 BenefitType = "MARKET_SUBSIDY_1_4"
	MARKETSUBSIDY59 BenefitType = "MARKET_SUBSIDY_5_9"
	SPLIT0012       BenefitType = "SPLIT_0_0_12"
	SPLIT004        BenefitType = "SPLIT_0_0_4"
	SPLIT006        BenefitType = "SPLIT_0_0_6"
)

func (e BenefitType) Valid() bool {
	switch e {
	case BESTS:
		return true
	case MARKETSUBSIDY10:
		return true
	case MARKETSUBSIDY14:
		return true
	case MARKETSUBSIDY59:
		return true
	case SPLIT0012:
		return true
	case SPLIT004:
		return true
	case SPLIT006:
		return true
	default:
		return false
	}
}

const (
	LIGHT  BusinessSubscriptionLevelType = "LIGHT"
	MEDIUM BusinessSubscriptionLevelType = "MEDIUM"
	NONE   BusinessSubscriptionLevelType = "NONE"
)

func (e BusinessSubscriptionLevelType) Valid() bool {
	switch e {
	case LIGHT:
		return true
	case MEDIUM:
		return true
	case NONE:
		return true
	default:
		return false
	}
}

const (
	MARKETYANDEXGO BusinessTraitType = "MARKET_YANDEX_GO"
)

func (e BusinessTraitType) Valid() bool {
	switch e {
	case MARKETYANDEXGO:
		return true
	default:
		return false
	}
}

func (e CalculatedTariffType) Valid() bool {
	switch e {
	case CalculatedTariffTypeAGENCYCOMMISSION:
		return true
	case CalculatedTariffTypeCROSSREGIONALDELIVERY:
		return true
	case CalculatedTariffTypeDELIVERYTOCUSTOMER:
		return true
	case CalculatedTariffTypeEXPRESSDELIVERY:
		return true
	case CalculatedTariffTypeFEE:
		return true
	case CalculatedTariffTypeMIDDLEMILE:
		return true
	case CalculatedTariffTypePAYMENTTRANSFER:
		return true
	case CalculatedTariffTypeSORTING:
		return true
	default:
		return false
	}
}

const (
	WEB CampaignSettingsScheduleSourceType = "WEB"
	YML CampaignSettingsScheduleSourceType = "YML"
)

func (e CampaignSettingsScheduleSourceType) Valid() bool {
	switch e {
	case WEB:
		return true
	case YML:
		return true
	default:
		return false
	}
}

const (
	CatalogLanguageTypeRU CatalogLanguageType = "RU"
	CatalogLanguageTypeUZ CatalogLanguageType = "UZ"
)

func (e CatalogLanguageType) Valid() bool {
	switch e {
	case CatalogLanguageTypeRU:
		return true
	case CatalogLanguageTypeUZ:
		return true
	default:
		return false
	}
}

func (e CategoryErrorType) Valid() bool {
	switch e {
	case CategoryErrorTypeCATEGORYISNOTLEAF:
		return true
	case CategoryErrorTypeUNKNOWNCATEGORY:
		return true
	default:
		return false
	}
}

const (
	ChannelTypeAPPPERFORMANCE          ChannelType = "APP_PERFORMANCE"
	ChannelTypeBANNERPICKUPPOINT       ChannelType = "BANNER_PICKUP_POINT"
	ChannelTypeBLOGGERPERFORMANCE      ChannelType = "BLOGGER_PERFORMANCE"
	ChannelTypeCAROUSELRETAILPAGE      ChannelType = "CAROUSEL_RETAIL_PAGE"
	ChannelTypeCPA                     ChannelType = "CPA"
	ChannelTypeDIGITALCHANNELBANNER    ChannelType = "DIGITAL_CHANNEL_BANNER"
	ChannelTypeMAINPAGECAROUSEL        ChannelType = "MAIN_PAGE_CAROUSEL"
	ChannelTypeMAINPAGECAROUSELWEB     ChannelType = "MAIN_PAGE_CAROUSEL_WEB"
	ChannelTypeOTHER                   ChannelType = "OTHER"
	ChannelTypePARTNERSMAINBANNER      ChannelType = "PARTNERS_MAIN_BANNER"
	ChannelTypePOPUPAPPLICATION        ChannelType = "POPUP_APPLICATION"
	ChannelTypePOSTTELEGRAM            ChannelType = "POST_TELEGRAM"
	ChannelTypePRODUCTRETAILPAGE       ChannelType = "PRODUCT_RETAIL_PAGE"
	ChannelTypePRODUCTSEPARATELANDING  ChannelType = "PRODUCT_SEPARATE_LANDING"
	ChannelTypePUSH                    ChannelType = "PUSH"
	ChannelTypeSTRETCHMAIN             ChannelType = "STRETCH_MAIN"
	ChannelTypeSUPERSHELFCATEGORY      ChannelType = "SUPER_SHELF_CATEGORY"
	ChannelTypeWEBPERFORMANCEDIRECT    ChannelType = "WEB_PERFORMANCE_DIRECT"
	ChannelTypeYANDEXECOSYSTEMCHANNELS ChannelType = "YANDEX_ECOSYSTEM_CHANNELS"
)

func (e ChannelType) Valid() bool {
	switch e {
	case ChannelTypeAPPPERFORMANCE:
		return true
	case ChannelTypeBANNERPICKUPPOINT:
		return true
	case ChannelTypeBLOGGERPERFORMANCE:
		return true
	case ChannelTypeCAROUSELRETAILPAGE:
		return true
	case ChannelTypeCPA:
		return true
	case ChannelTypeDIGITALCHANNELBANNER:
		return true
	case ChannelTypeMAINPAGECAROUSEL:
		return true
	case ChannelTypeMAINPAGECAROUSELWEB:
		return true
	case ChannelTypeOTHER:
		return true
	case ChannelTypePARTNERSMAINBANNER:
		return true
	case ChannelTypePOPUPAPPLICATION:
		return true
	case ChannelTypePOSTTELEGRAM:
		return true
	case ChannelTypePRODUCTRETAILPAGE:
		return true
	case ChannelTypePRODUCTSEPARATELANDING:
		return true
	case ChannelTypePUSH:
		return true
	case ChannelTypeSTRETCHMAIN:
		return true
	case ChannelTypeSUPERSHELFCATEGORY:
		return true
	case ChannelTypeWEBPERFORMANCEDIRECT:
		return true
	case ChannelTypeYANDEXECOSYSTEMCHANNELS:
		return true
	default:
		return false
	}
}

func (e ChatContextIdentifiableType) Valid() bool {
	switch e {
	case ChatContextIdentifiableTypeORDER:
		return true
	case ChatContextIdentifiableTypeRETURN:
		return true
	default:
		return false
	}
}

func (e ChatContextType) Valid() bool {
	switch e {
	case ChatContextTypeDIRECT:
		return true
	case ChatContextTypeORDER:
		return true
	case ChatContextTypeRETURN:
		return true
	default:
		return false
	}
}

func (e ChatMessageSenderType) Valid() bool {
	switch e {
	case ChatMessageSenderTypeCUSTOMER:
		return true
	case ChatMessageSenderTypeMARKET:
		return true
	case ChatMessageSenderTypePARTNER:
		return true
	case ChatMessageSenderTypeSUPPORT:
		return true
	default:
		return false
	}
}

func (e ChatStatusType) Valid() bool {
	switch e {
	case ChatStatusTypeFINISHED:
		return true
	case ChatStatusTypeNEW:
		return true
	case ChatStatusTypeWAITINGFORARBITER:
		return true
	case ChatStatusTypeWAITINGFORCUSTOMER:
		return true
	case ChatStatusTypeWAITINGFORMARKET:
		return true
	case ChatStatusTypeWAITINGFORPARTNER:
		return true
	default:
		return false
	}
}

const (
	ARBITRAGE ChatType = "ARBITRAGE"
	CHAT      ChatType = "CHAT"
)

func (e ChatType) Valid() bool {
	switch e {
	case ARBITRAGE:
		return true
	case CHAT:
		return true
	default:
		return false
	}
}

const (
	CisStatusTypeFAILED          CisStatusType = "FAILED"
	CisStatusTypeINPROGRESS      CisStatusType = "IN_PROGRESS"
	CisStatusTypeINVALID         CisStatusType = "INVALID"
	CisStatusTypeNOTONVALIDATION CisStatusType = "NOT_ON_VALIDATION"
	CisStatusTypeOK              CisStatusType = "OK"
)

func (e CisStatusType) Valid() bool {
	switch e {
	case CisStatusTypeFAILED:
		return true
	case CisStatusTypeINPROGRESS:
		return true
	case CisStatusTypeINVALID:
		return true
	case CisStatusTypeNOTONVALIDATION:
		return true
	case CisStatusTypeOK:
		return true
	default:
		return false
	}
}

const (
	CISGTINNOTFOUND                    CisSubstatusType = "CIS_GTIN_NOT_FOUND"
	CISNOTFOUNDINGISMT                 CisSubstatusType = "CIS_NOT_FOUND_IN_GIS_MT"
	CISSERIALNUMBERNOTFOUND            CisSubstatusType = "CIS_SERIAL_NUMBER_NOT_FOUND"
	CISVALIDATIONERROR                 CisSubstatusType = "CIS_VALIDATION_ERROR"
	CRYPTOTAILFORMATMISMATCHCISTYPE    CisSubstatusType = "CRYPTO_TAIL_FORMAT_MISMATCH_CIS_TYPE"
	EXPIREDITEM                        CisSubstatusType = "EXPIRED_ITEM"
	INVALIDCRYPTOKEY                   CisSubstatusType = "INVALID_CRYPTO_KEY"
	INVALIDCRYPTOTAIL                  CisSubstatusType = "INVALID_CRYPTO_TAIL"
	INVALIDSYMBOLSFOUND                CisSubstatusType = "INVALID_SYMBOLS_FOUND"
	ITEMSOLD                           CisSubstatusType = "ITEM_SOLD"
	NOTPLACEDONMARKET                  CisSubstatusType = "NOT_PLACED_ON_MARKET"
	NOTPRINTEDONPACKAGE                CisSubstatusType = "NOT_PRINTED_ON_PACKAGE"
	SALEBLOCKEDBYOGB                   CisSubstatusType = "SALE_BLOCKED_BY_OGB"
	UNSUPPORTEDAIFOUND                 CisSubstatusType = "UNSUPPORTED_AI_FOUND"
	VERIFICATIONFAILEDINEMITTERCOUNTRY CisSubstatusType = "VERIFICATION_FAILED_IN_EMITTER_COUNTRY"
	WRONGOWNERINN                      CisSubstatusType = "WRONG_OWNER_INN"
)

func (e CisSubstatusType) Valid() bool {
	switch e {
	case CISGTINNOTFOUND:
		return true
	case CISNOTFOUNDINGISMT:
		return true
	case CISSERIALNUMBERNOTFOUND:
		return true
	case CISVALIDATIONERROR:
		return true
	case CRYPTOTAILFORMATMISMATCHCISTYPE:
		return true
	case EXPIREDITEM:
		return true
	case INVALIDCRYPTOKEY:
		return true
	case INVALIDCRYPTOTAIL:
		return true
	case INVALIDSYMBOLSFOUND:
		return true
	case ITEMSOLD:
		return true
	case NOTPLACEDONMARKET:
		return true
	case NOTPRINTEDONPACKAGE:
		return true
	case SALEBLOCKEDBYOGB:
		return true
	case UNSUPPORTEDAIFOUND:
		return true
	case VERIFICATIONFAILEDINEMITTERCOUNTRY:
		return true
	case WRONGOWNERINN:
		return true
	default:
		return false
	}
}

const (
	INCOME    ClosureDocumentsContractType = "INCOME"
	MARKETING ClosureDocumentsContractType = "MARKETING"
	OUTCOME   ClosureDocumentsContractType = "OUTCOME"
)

func (e ClosureDocumentsContractType) Valid() bool {
	switch e {
	case INCOME:
		return true
	case MARKETING:
		return true
	case OUTCOME:
		return true
	default:
		return false
	}
}

const (
	CommodityCodeTypeCUSTOMSCOMMODITYCODE CommodityCodeType = "CUSTOMS_COMMODITY_CODE"
	CommodityCodeTypeIKPUCODE             CommodityCodeType = "IKPU_CODE"
)

func (e CommodityCodeType) Valid() bool {
	switch e {
	case CommodityCodeTypeCUSTOMSCOMMODITYCODE:
		return true
	case CommodityCodeTypeIKPUCODE:
		return true
	default:
		return false
	}
}

func (e CreateOrderPackageType) Valid() bool {
	switch e {
	case CreateOrderPackageTypeBRAND:
		return true
	case CreateOrderPackageTypeWHITELABEL:
		return true
	default:
		return false
	}
}

const (
	CurrencyTypeAED CurrencyType = "AED"
	CurrencyTypeAFN CurrencyType = "AFN"
	CurrencyTypeALL CurrencyType = "ALL"
	CurrencyTypeAMD CurrencyType = "AMD"
	CurrencyTypeAOA CurrencyType = "AOA"
	CurrencyTypeARS CurrencyType = "ARS"
	CurrencyTypeAUD CurrencyType = "AUD"
	CurrencyTypeAZN CurrencyType = "AZN"
	CurrencyTypeBDT CurrencyType = "BDT"
	CurrencyTypeBGN CurrencyType = "BGN"
	CurrencyTypeBHD CurrencyType = "BHD"
	CurrencyTypeBIF CurrencyType = "BIF"
	CurrencyTypeBND CurrencyType = "BND"
	CurrencyTypeBOB CurrencyType = "BOB"
	CurrencyTypeBRL CurrencyType = "BRL"
	CurrencyTypeBWP CurrencyType = "BWP"
	CurrencyTypeBYN CurrencyType = "BYN"
	CurrencyTypeBYR CurrencyType = "BYR"
	CurrencyTypeCAD CurrencyType = "CAD"
	CurrencyTypeCDF CurrencyType = "CDF"
	CurrencyTypeCHF CurrencyType = "CHF"
	CurrencyTypeCLP CurrencyType = "CLP"
	CurrencyTypeCNY CurrencyType = "CNY"
	CurrencyTypeCOP CurrencyType = "COP"
	CurrencyTypeCRC CurrencyType = "CRC"
	CurrencyTypeCUP CurrencyType = "CUP"
	CurrencyTypeCZK CurrencyType = "CZK"
	CurrencyTypeDJF CurrencyType = "DJF"
	CurrencyTypeDKK CurrencyType = "DKK"
	CurrencyTypeDZD CurrencyType = "DZD"
	CurrencyTypeEEK CurrencyType = "EEK"
	CurrencyTypeEGP CurrencyType = "EGP"
	CurrencyTypeETB CurrencyType = "ETB"
	CurrencyTypeEUR CurrencyType = "EUR"
	CurrencyTypeGBP CurrencyType = "GBP"
	CurrencyTypeGEL CurrencyType = "GEL"
	CurrencyTypeGHS CurrencyType = "GHS"
	CurrencyTypeGMD CurrencyType = "GMD"
	CurrencyTypeGNF CurrencyType = "GNF"
	CurrencyTypeHKD CurrencyType = "HKD"
	CurrencyTypeHRK CurrencyType = "HRK"
	CurrencyTypeHUF CurrencyType = "HUF"
	CurrencyTypeIDR CurrencyType = "IDR"
	CurrencyTypeILS CurrencyType = "ILS"
	CurrencyTypeINR CurrencyType = "INR"
	CurrencyTypeIQD CurrencyType = "IQD"
	CurrencyTypeIRR CurrencyType = "IRR"
	CurrencyTypeISK CurrencyType = "ISK"
	CurrencyTypeJOD CurrencyType = "JOD"
	CurrencyTypeJPY CurrencyType = "JPY"
	CurrencyTypeKES CurrencyType = "KES"
	CurrencyTypeKGS CurrencyType = "KGS"
	CurrencyTypeKHR CurrencyType = "KHR"
	CurrencyTypeKPW CurrencyType = "KPW"
	CurrencyTypeKRW CurrencyType = "KRW"
	CurrencyTypeKWD CurrencyType = "KWD"
	CurrencyTypeKZT CurrencyType = "KZT"
	CurrencyTypeLAK CurrencyType = "LAK"
	CurrencyTypeLBP CurrencyType = "LBP"
	CurrencyTypeLKR CurrencyType = "LKR"
	CurrencyTypeLTL CurrencyType = "LTL"
	CurrencyTypeLVL CurrencyType = "LVL"
	CurrencyTypeLYD CurrencyType = "LYD"
	CurrencyTypeMAD CurrencyType = "MAD"
	CurrencyTypeMDL CurrencyType = "MDL"
	CurrencyTypeMGA CurrencyType = "MGA"
	CurrencyTypeMKD CurrencyType = "MKD"
	CurrencyTypeMNT CurrencyType = "MNT"
	CurrencyTypeMRO CurrencyType = "MRO"
	CurrencyTypeMUR CurrencyType = "MUR"
	CurrencyTypeMWK CurrencyType = "MWK"
	CurrencyTypeMXN CurrencyType = "MXN"
	CurrencyTypeMYR CurrencyType = "MYR"
	CurrencyTypeMZN CurrencyType = "MZN"
	CurrencyTypeNAD CurrencyType = "NAD"
	CurrencyTypeNGN CurrencyType = "NGN"
	CurrencyTypeNIO CurrencyType = "NIO"
	CurrencyTypeNOK CurrencyType = "NOK"
	CurrencyTypeNPR CurrencyType = "NPR"
	CurrencyTypeNZD CurrencyType = "NZD"
	CurrencyTypeOMR CurrencyType = "OMR"
	CurrencyTypePEN CurrencyType = "PEN"
	CurrencyTypePHP CurrencyType = "PHP"
	CurrencyTypePKR CurrencyType = "PKR"
	CurrencyTypePLN CurrencyType = "PLN"
	CurrencyTypePYG CurrencyType = "PYG"
	CurrencyTypeQAR CurrencyType = "QAR"
	CurrencyTypeRON CurrencyType = "RON"
	CurrencyTypeRSD CurrencyType = "RSD"
	CurrencyTypeRUR CurrencyType = "RUR"
	CurrencyTypeSAR CurrencyType = "SAR"
	CurrencyTypeSCR CurrencyType = "SCR"
	CurrencyTypeSDG CurrencyType = "SDG"
	CurrencyTypeSEK CurrencyType = "SEK"
	CurrencyTypeSGD CurrencyType = "SGD"
	CurrencyTypeSKK CurrencyType = "SKK"
	CurrencyTypeSLL CurrencyType = "SLL"
	CurrencyTypeSOS CurrencyType = "SOS"
	CurrencyTypeSRD CurrencyType = "SRD"
	CurrencyTypeSYP CurrencyType = "SYP"
	CurrencyTypeSZL CurrencyType = "SZL"
	CurrencyTypeTHB CurrencyType = "THB"
	CurrencyTypeTJS CurrencyType = "TJS"
	CurrencyTypeTL  CurrencyType = "TL"
	CurrencyTypeTMM CurrencyType = "TMM"
	CurrencyTypeTND CurrencyType = "TND"
	CurrencyTypeTRY CurrencyType = "TRY"
	CurrencyTypeTWD CurrencyType = "TWD"
	CurrencyTypeTZS CurrencyType = "TZS"
	CurrencyTypeUAH CurrencyType = "UAH"
	CurrencyTypeUE  CurrencyType = "UE"
	CurrencyTypeUGX CurrencyType = "UGX"
	CurrencyTypeUSD CurrencyType = "USD"
	CurrencyTypeUYU CurrencyType = "UYU"
	CurrencyTypeUZS CurrencyType = "UZS"
	CurrencyTypeVEF CurrencyType = "VEF"
	CurrencyTypeVND CurrencyType = "VND"
	CurrencyTypeXAF CurrencyType = "XAF"
	CurrencyTypeXDR CurrencyType = "XDR"
	CurrencyTypeXOF CurrencyType = "XOF"
	CurrencyTypeYER CurrencyType = "YER"
	CurrencyTypeZAR CurrencyType = "ZAR"
	CurrencyTypeZMK CurrencyType = "ZMK"
)

func (e CurrencyType) Valid() bool {
	switch e {
	case CurrencyTypeAED:
		return true
	case CurrencyTypeAFN:
		return true
	case CurrencyTypeALL:
		return true
	case CurrencyTypeAMD:
		return true
	case CurrencyTypeAOA:
		return true
	case CurrencyTypeARS:
		return true
	case CurrencyTypeAUD:
		return true
	case CurrencyTypeAZN:
		return true
	case CurrencyTypeBDT:
		return true
	case CurrencyTypeBGN:
		return true
	case CurrencyTypeBHD:
		return true
	case CurrencyTypeBIF:
		return true
	case CurrencyTypeBND:
		return true
	case CurrencyTypeBOB:
		return true
	case CurrencyTypeBRL:
		return true
	case CurrencyTypeBWP:
		return true
	case CurrencyTypeBYN:
		return true
	case CurrencyTypeBYR:
		return true
	case CurrencyTypeCAD:
		return true
	case CurrencyTypeCDF:
		return true
	case CurrencyTypeCHF:
		return true
	case CurrencyTypeCLP:
		return true
	case CurrencyTypeCNY:
		return true
	case CurrencyTypeCOP:
		return true
	case CurrencyTypeCRC:
		return true
	case CurrencyTypeCUP:
		return true
	case CurrencyTypeCZK:
		return true
	case CurrencyTypeDJF:
		return true
	case CurrencyTypeDKK:
		return true
	case CurrencyTypeDZD:
		return true
	case CurrencyTypeEEK:
		return true
	case CurrencyTypeEGP:
		return true
	case CurrencyTypeETB:
		return true
	case CurrencyTypeEUR:
		return true
	case CurrencyTypeGBP:
		return true
	case CurrencyTypeGEL:
		return true
	case CurrencyTypeGHS:
		return true
	case CurrencyTypeGMD:
		return true
	case CurrencyTypeGNF:
		return true
	case CurrencyTypeHKD:
		return true
	case CurrencyTypeHRK:
		return true
	case CurrencyTypeHUF:
		return true
	case CurrencyTypeIDR:
		return true
	case CurrencyTypeILS:
		return true
	case CurrencyTypeINR:
		return true
	case CurrencyTypeIQD:
		return true
	case CurrencyTypeIRR:
		return true
	case CurrencyTypeISK:
		return true
	case CurrencyTypeJOD:
		return true
	case CurrencyTypeJPY:
		return true
	case CurrencyTypeKES:
		return true
	case CurrencyTypeKGS:
		return true
	case CurrencyTypeKHR:
		return true
	case CurrencyTypeKPW:
		return true
	case CurrencyTypeKRW:
		return true
	case CurrencyTypeKWD:
		return true
	case CurrencyTypeKZT:
		return true
	case CurrencyTypeLAK:
		return true
	case CurrencyTypeLBP:
		return true
	case CurrencyTypeLKR:
		return true
	case CurrencyTypeLTL:
		return true
	case CurrencyTypeLVL:
		return true
	case CurrencyTypeLYD:
		return true
	case CurrencyTypeMAD:
		return true
	case CurrencyTypeMDL:
		return true
	case CurrencyTypeMGA:
		return true
	case CurrencyTypeMKD:
		return true
	case CurrencyTypeMNT:
		return true
	case CurrencyTypeMRO:
		return true
	case CurrencyTypeMUR:
		return true
	case CurrencyTypeMWK:
		return true
	case CurrencyTypeMXN:
		return true
	case CurrencyTypeMYR:
		return true
	case CurrencyTypeMZN:
		return true
	case CurrencyTypeNAD:
		return true
	case CurrencyTypeNGN:
		return true
	case CurrencyTypeNIO:
		return true
	case CurrencyTypeNOK:
		return true
	case CurrencyTypeNPR:
		return true
	case CurrencyTypeNZD:
		return true
	case CurrencyTypeOMR:
		return true
	case CurrencyTypePEN:
		return true
	case CurrencyTypePHP:
		return true
	case CurrencyTypePKR:
		return true
	case CurrencyTypePLN:
		return true
	case CurrencyTypePYG:
		return true
	case CurrencyTypeQAR:
		return true
	case CurrencyTypeRON:
		return true
	case CurrencyTypeRSD:
		return true
	case CurrencyTypeRUR:
		return true
	case CurrencyTypeSAR:
		return true
	case CurrencyTypeSCR:
		return true
	case CurrencyTypeSDG:
		return true
	case CurrencyTypeSEK:
		return true
	case CurrencyTypeSGD:
		return true
	case CurrencyTypeSKK:
		return true
	case CurrencyTypeSLL:
		return true
	case CurrencyTypeSOS:
		return true
	case CurrencyTypeSRD:
		return true
	case CurrencyTypeSYP:
		return true
	case CurrencyTypeSZL:
		return true
	case CurrencyTypeTHB:
		return true
	case CurrencyTypeTJS:
		return true
	case CurrencyTypeTL:
		return true
	case CurrencyTypeTMM:
		return true
	case CurrencyTypeTND:
		return true
	case CurrencyTypeTRY:
		return true
	case CurrencyTypeTWD:
		return true
	case CurrencyTypeTZS:
		return true
	case CurrencyTypeUAH:
		return true
	case CurrencyTypeUE:
		return true
	case CurrencyTypeUGX:
		return true
	case CurrencyTypeUSD:
		return true
	case CurrencyTypeUYU:
		return true
	case CurrencyTypeUZS:
		return true
	case CurrencyTypeVEF:
		return true
	case CurrencyTypeVND:
		return true
	case CurrencyTypeXAF:
		return true
	case CurrencyTypeXDR:
		return true
	case CurrencyTypeXOF:
		return true
	case CurrencyTypeYER:
		return true
	case CurrencyTypeZAR:
		return true
	case CurrencyTypeZMK:
		return true
	default:
		return false
	}
}

const (
	FRIDAY    DayOfWeekType = "FRIDAY"
	MONDAY    DayOfWeekType = "MONDAY"
	SATURDAY  DayOfWeekType = "SATURDAY"
	SUNDAY    DayOfWeekType = "SUNDAY"
	THURSDAY  DayOfWeekType = "THURSDAY"
	TUESDAY   DayOfWeekType = "TUESDAY"
	WEDNESDAY DayOfWeekType = "WEDNESDAY"
)

func (e DayOfWeekType) Valid() bool {
	switch e {
	case FRIDAY:
		return true
	case MONDAY:
		return true
	case SATURDAY:
		return true
	case SUNDAY:
		return true
	case THURSDAY:
		return true
	case TUESDAY:
		return true
	case WEDNESDAY:
		return true
	default:
		return false
	}
}

func (e DeleteOfferParameterType) Valid() bool {
	switch e {
	case DeleteOfferParameterTypeADDITIONALEXPENSES:
		return true
	case DeleteOfferParameterTypeADULT:
		return true
	case DeleteOfferParameterTypeAGE:
		return true
	case DeleteOfferParameterTypeBARCODES:
		return true
	case DeleteOfferParameterTypeBOXCOUNT:
		return true
	case DeleteOfferParameterTypeCERTIFICATES:
		return true
	case DeleteOfferParameterTypeCOMMODITYCODES:
		return true
	case DeleteOfferParameterTypeCONDITION:
		return true
	case DeleteOfferParameterTypeCUSTOMSCOMMODITYCODE:
		return true
	case DeleteOfferParameterTypeDESCRIPTION:
		return true
	case DeleteOfferParameterTypeDOWNLOADABLE:
		return true
	case DeleteOfferParameterTypeGUARANTEEPERIOD:
		return true
	case DeleteOfferParameterTypeLIFETIME:
		return true
	case DeleteOfferParameterTypeMANUALS:
		return true
	case DeleteOfferParameterTypeMANUFACTURERCOUNTRIES:
		return true
	case DeleteOfferParameterTypePARAMETERS:
		return true
	case DeleteOfferParameterTypePICTURES:
		return true
	case DeleteOfferParameterTypePURCHASEPRICE:
		return true
	case DeleteOfferParameterTypeSHELFLIFE:
		return true
	case DeleteOfferParameterTypeTAGS:
		return true
	case DeleteOfferParameterTypeTYPE:
		return true
	case DeleteOfferParameterTypeVENDORCODE:
		return true
	case DeleteOfferParameterTypeVIDEOS:
		return true
	default:
		return false
	}
}

const (
	DeliveryPaymentTypePREPAID DeliveryPaymentType = "PREPAID"
)

func (e DeliveryPaymentType) Valid() bool {
	switch e {
	case DeliveryPaymentTypePREPAID:
		return true
	default:
		return false
	}
}

const (
	EacVerificationStatusTypeACCEPTED   EacVerificationStatusType = "ACCEPTED"
	EacVerificationStatusTypeNEEDUPDATE EacVerificationStatusType = "NEED_UPDATE"
	EacVerificationStatusTypeREJECTED   EacVerificationStatusType = "REJECTED"
)

func (e EacVerificationStatusType) Valid() bool {
	switch e {
	case EacVerificationStatusTypeACCEPTED:
		return true
	case EacVerificationStatusTypeNEEDUPDATE:
		return true
	case EacVerificationStatusTypeREJECTED:
		return true
	default:
		return false
	}
}

func (e ExternalReturnDecisionReasonType) Valid() bool {
	switch e {
	case ExternalReturnDecisionReasonTypeBADQUALITY:
		return true
	case ExternalReturnDecisionReasonTypeDOESNOTFIT:
		return true
	case ExternalReturnDecisionReasonTypeWRONGITEM:
		return true
	default:
		return false
	}
}

func (e ExternalReturnDecisionSubreasonType) Valid() bool {
	switch e {
	case ExternalReturnDecisionSubreasonTypeBADPACKAGE:
		return true
	case ExternalReturnDecisionSubreasonTypeDAMAGED:
		return true
	case ExternalReturnDecisionSubreasonTypeDELIVEREDTOOLONG:
		return true
	case ExternalReturnDecisionSubreasonTypeDIDNOTMATCHDESCRIPTION:
		return true
	case ExternalReturnDecisionSubreasonTypeINCOMPLETENESS:
		return true
	case ExternalReturnDecisionSubreasonTypeNOTWORKING:
		return true
	case ExternalReturnDecisionSubreasonTypeUSERCHANGEDMIND:
		return true
	case ExternalReturnDecisionSubreasonTypeUSERDIDNOTLIKE:
		return true
	case ExternalReturnDecisionSubreasonTypeWRONGCOLOR:
		return true
	case ExternalReturnDecisionSubreasonTypeWRONGITEM:
		return true
	default:
		return false
	}
}

const (
	FeedbackReactionStatusTypeALL          FeedbackReactionStatusType = "ALL"
	FeedbackReactionStatusTypeNEEDREACTION FeedbackReactionStatusType = "NEED_REACTION"
)

func (e FeedbackReactionStatusType) Valid() bool {
	switch e {
	case FeedbackReactionStatusTypeALL:
		return true
	case FeedbackReactionStatusTypeNEEDREACTION:
		return true
	default:
		return false
	}
}

func (e GoodsFeedbackCommentAuthorType) Valid() bool {
	switch e {
	case GoodsFeedbackCommentAuthorTypeBUSINESS:
		return true
	case GoodsFeedbackCommentAuthorTypeUSER:
		return true
	default:
		return false
	}
}

func (e GoodsFeedbackCommentStatusType) Valid() bool {
	switch e {
	case GoodsFeedbackCommentStatusTypeBANNED:
		return true
	case GoodsFeedbackCommentStatusTypeDELETED:
		return true
	case GoodsFeedbackCommentStatusTypePUBLISHED:
		return true
	case GoodsFeedbackCommentStatusTypeUNMODERATED:
		return true
	default:
		return false
	}
}

func (e KeyIndicatorsReportDetalizationLevelType) Valid() bool {
	switch e {
	case KeyIndicatorsReportDetalizationLevelTypeMONTH:
		return true
	case KeyIndicatorsReportDetalizationLevelTypeWEEK:
		return true
	default:
		return false
	}
}

const (
	SORTBYGIVENORDER     LabelsSortingType = "SORT_BY_GIVEN_ORDER"
	SORTBYORDERCREATEDAT LabelsSortingType = "SORT_BY_ORDER_CREATED_AT"
)

func (e LabelsSortingType) Valid() bool {
	switch e {
	case SORTBYGIVENORDER:
		return true
	case SORTBYORDERCREATEDAT:
		return true
	default:
		return false
	}
}

const (
	LanguageTypeEN LanguageType = "EN"
	LanguageTypeRU LanguageType = "RU"
)

func (e LanguageType) Valid() bool {
	switch e {
	case LanguageTypeEN:
		return true
	case LanguageTypeRU:
		return true
	default:
		return false
	}
}

const (
	LicenseCheckStatusTypeDONTWANT   LicenseCheckStatusType = "DONT_WANT"
	LicenseCheckStatusTypeFAIL       LicenseCheckStatusType = "FAIL"
	LicenseCheckStatusTypeFAILMANUAL LicenseCheckStatusType = "FAIL_MANUAL"
	LicenseCheckStatusTypeNEW        LicenseCheckStatusType = "NEW"
	LicenseCheckStatusTypeREVOKE     LicenseCheckStatusType = "REVOKE"
	LicenseCheckStatusTypeSUCCESS    LicenseCheckStatusType = "SUCCESS"
)

func (e LicenseCheckStatusType) Valid() bool {
	switch e {
	case LicenseCheckStatusTypeDONTWANT:
		return true
	case LicenseCheckStatusTypeFAIL:
		return true
	case LicenseCheckStatusTypeFAILMANUAL:
		return true
	case LicenseCheckStatusTypeNEW:
		return true
	case LicenseCheckStatusTypeREVOKE:
		return true
	case LicenseCheckStatusTypeSUCCESS:
		return true
	default:
		return false
	}
}

const (
	LicenseTypeALCOHOL LicenseType = "ALCOHOL"
	LicenseTypeUNKNOWN LicenseType = "UNKNOWN"
)

func (e LicenseType) Valid() bool {
	switch e {
	case LicenseTypeALCOHOL:
		return true
	case LicenseTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e LogisticPointBrandType) Valid() bool {
	switch e {
	case LogisticPointBrandTypeMARKET:
		return true
	default:
		return false
	}
}

const (
	RETURNALLOWED LogisticPointFeatureType = "RETURN_ALLOWED"
)

func (e LogisticPointFeatureType) Valid() bool {
	switch e {
	case RETURNALLOWED:
		return true
	default:
		return false
	}
}

func (e LogisticPointType) Valid() bool {
	switch e {
	case LogisticPointTypePICKUPMIXED:
		return true
	case LogisticPointTypePICKUPPOINT:
		return true
	case LogisticPointTypePICKUPPOSTOFFICE:
		return true
	case LogisticPointTypePICKUPRETAIL:
		return true
	case LogisticPointTypePICKUPTERMINAL:
		return true
	case LogisticPointTypeWAREHOUSE:
		return true
	default:
		return false
	}
}

const (
	MechanicsTypeBLUEFLASH       MechanicsType = "BLUE_FLASH"
	MechanicsTypeDIRECTDISCOUNT  MechanicsType = "DIRECT_DISCOUNT"
	MechanicsTypeMARKETPROMOCODE MechanicsType = "MARKET_PROMOCODE"
)

func (e MechanicsType) Valid() bool {
	switch e {
	case MechanicsTypeBLUEFLASH:
		return true
	case MechanicsTypeDIRECTDISCOUNT:
		return true
	case MechanicsTypeMARKETPROMOCODE:
		return true
	default:
		return false
	}
}

const (
	MediaFileUploadStateTypeFAILED    MediaFileUploadStateType = "FAILED"
	MediaFileUploadStateTypeUPLOADED  MediaFileUploadStateType = "UPLOADED"
	MediaFileUploadStateTypeUPLOADING MediaFileUploadStateType = "UPLOADING"
)

func (e MediaFileUploadStateType) Valid() bool {
	switch e {
	case MediaFileUploadStateTypeFAILED:
		return true
	case MediaFileUploadStateTypeUPLOADED:
		return true
	case MediaFileUploadStateTypeUPLOADING:
		return true
	default:
		return false
	}
}

func (e OfferCampaignStatusType) Valid() bool {
	switch e {
	case OfferCampaignStatusTypeARCHIVED:
		return true
	case OfferCampaignStatusTypeCHECKING:
		return true
	case OfferCampaignStatusTypeCREATINGCARD:
		return true
	case OfferCampaignStatusTypeDISABLEDAUTOMATICALLY:
		return true
	case OfferCampaignStatusTypeDISABLEDBYPARTNER:
		return true
	case OfferCampaignStatusTypeNOCARD:
		return true
	case OfferCampaignStatusTypeNOSTOCKS:
		return true
	case OfferCampaignStatusTypePUBLISHED:
		return true
	case OfferCampaignStatusTypeREJECTEDBYMARKET:
		return true
	default:
		return false
	}
}

const (
	ACTUAL   OfferCardContentStatusType = "ACTUAL"
	UPDATING OfferCardContentStatusType = "UPDATING"
)

func (e OfferCardContentStatusType) Valid() bool {
	switch e {
	case ACTUAL:
		return true
	case UPDATING:
		return true
	default:
		return false
	}
}

const (
	ADDITIONAL         OfferCardRecommendationType = "ADDITIONAL"
	AVERAGEPICTURESIZE OfferCardRecommendationType = "AVERAGE_PICTURE_SIZE"
	AVERAGEVIDEOSIZE   OfferCardRecommendationType = "AVERAGE_VIDEO_SIZE"
	DESCRIPTIONLENGTH  OfferCardRecommendationType = "DESCRIPTION_LENGTH"
	DISTINCTIVE        OfferCardRecommendationType = "DISTINCTIVE"
	FILTERABLE         OfferCardRecommendationType = "FILTERABLE"
	FIRSTPICTURESIZE   OfferCardRecommendationType = "FIRST_PICTURE_SIZE"
	FIRSTVIDEOLENGTH   OfferCardRecommendationType = "FIRST_VIDEO_LENGTH"
	FIRSTVIDEOSIZE     OfferCardRecommendationType = "FIRST_VIDEO_SIZE"
	HASBARCODE         OfferCardRecommendationType = "HAS_BARCODE"
	HASDESCRIPTION     OfferCardRecommendationType = "HAS_DESCRIPTION"
	HASVIDEO           OfferCardRecommendationType = "HAS_VIDEO"
	MAIN               OfferCardRecommendationType = "MAIN"
	PICTURECOUNT       OfferCardRecommendationType = "PICTURE_COUNT"
	RECOGNIZEDVENDOR   OfferCardRecommendationType = "RECOGNIZED_VENDOR"
	TITLELENGTH        OfferCardRecommendationType = "TITLE_LENGTH"
	VIDEOCOUNT         OfferCardRecommendationType = "VIDEO_COUNT"
)

func (e OfferCardRecommendationType) Valid() bool {
	switch e {
	case ADDITIONAL:
		return true
	case AVERAGEPICTURESIZE:
		return true
	case AVERAGEVIDEOSIZE:
		return true
	case DESCRIPTIONLENGTH:
		return true
	case DISTINCTIVE:
		return true
	case FILTERABLE:
		return true
	case FIRSTPICTURESIZE:
		return true
	case FIRSTVIDEOLENGTH:
		return true
	case FIRSTVIDEOSIZE:
		return true
	case HASBARCODE:
		return true
	case HASDESCRIPTION:
		return true
	case HASVIDEO:
		return true
	case MAIN:
		return true
	case PICTURECOUNT:
		return true
	case RECOGNIZEDVENDOR:
		return true
	case TITLELENGTH:
		return true
	case VIDEOCOUNT:
		return true
	default:
		return false
	}
}

const (
	HASCARDCANNOTUPDATE        OfferCardStatusType = "HAS_CARD_CAN_NOT_UPDATE"
	HASCARDCANUPDATE           OfferCardStatusType = "HAS_CARD_CAN_UPDATE"
	HASCARDCANUPDATEERRORS     OfferCardStatusType = "HAS_CARD_CAN_UPDATE_ERRORS"
	HASCARDCANUPDATEPROCESSING OfferCardStatusType = "HAS_CARD_CAN_UPDATE_PROCESSING"
	NOCARDADDTOCAMPAIGN        OfferCardStatusType = "NO_CARD_ADD_TO_CAMPAIGN"
	NOCARDERRORS               OfferCardStatusType = "NO_CARD_ERRORS"
	NOCARDMARKETWILLCREATE     OfferCardStatusType = "NO_CARD_MARKET_WILL_CREATE"
	NOCARDNEEDCONTENT          OfferCardStatusType = "NO_CARD_NEED_CONTENT"
	NOCARDPROCESSING           OfferCardStatusType = "NO_CARD_PROCESSING"
)

func (e OfferCardStatusType) Valid() bool {
	switch e {
	case HASCARDCANNOTUPDATE:
		return true
	case HASCARDCANUPDATE:
		return true
	case HASCARDCANUPDATEERRORS:
		return true
	case HASCARDCANUPDATEPROCESSING:
		return true
	case NOCARDADDTOCAMPAIGN:
		return true
	case NOCARDERRORS:
		return true
	case NOCARDMARKETWILLCREATE:
		return true
	case NOCARDNEEDCONTENT:
		return true
	case NOCARDPROCESSING:
		return true
	default:
		return false
	}
}

func (e OfferConditionQualityType) Valid() bool {
	switch e {
	case OfferConditionQualityTypeEXCELLENT:
		return true
	case OfferConditionQualityTypeGOOD:
		return true
	case OfferConditionQualityTypeNOTSPECIFIED:
		return true
	case OfferConditionQualityTypePERFECT:
		return true
	default:
		return false
	}
}

func (e OfferConditionType) Valid() bool {
	switch e {
	case OfferConditionTypeNOTSPECIFIED:
		return true
	case OfferConditionTypePREOWNED:
		return true
	case OfferConditionTypeREDUCTION:
		return true
	case OfferConditionTypeREFURBISHED:
		return true
	case OfferConditionTypeRENOVATED:
		return true
	case OfferConditionTypeSHOWCASESAMPLE:
		return true
	default:
		return false
	}
}

func (e OfferContentErrorType) Valid() bool {
	switch e {
	case OfferContentErrorTypeINVALIDCATEGORY:
		return true
	case OfferContentErrorTypeINVALIDGROUPIDCHARACTERS:
		return true
	case OfferContentErrorTypeINVALIDGROUPIDLENGTH:
		return true
	case OfferContentErrorTypeINVALIDUNITID:
		return true
	case OfferContentErrorTypeNUMBERFORMAT:
		return true
	case OfferContentErrorTypeOFFERNOTFOUND:
		return true
	case OfferContentErrorTypeUNEXPECTEDBOOLEANVALUE:
		return true
	case OfferContentErrorTypeUNKNOWNCATEGORY:
		return true
	case OfferContentErrorTypeUNKNOWNPARAMETER:
		return true
	default:
		return false
	}
}

const (
	EMPTYMARKETCATEGORY      OfferMappingErrorType = "EMPTY_MARKET_CATEGORY"
	INVALIDCATEGORY          OfferMappingErrorType = "INVALID_CATEGORY"
	INVALIDCOMMODITYCODE     OfferMappingErrorType = "INVALID_COMMODITY_CODE"
	INVALIDGROUPIDCHARACTERS OfferMappingErrorType = "INVALID_GROUP_ID_CHARACTERS"
	INVALIDGROUPIDLENGTH     OfferMappingErrorType = "INVALID_GROUP_ID_LENGTH"
	INVALIDPICKERURL         OfferMappingErrorType = "INVALID_PICKER_URL"
	INVALIDUNITID            OfferMappingErrorType = "INVALID_UNIT_ID"
	LOCKEDDIMENSIONS         OfferMappingErrorType = "LOCKED_DIMENSIONS"
	NUMBERFORMAT             OfferMappingErrorType = "NUMBER_FORMAT"
	UNEXPECTEDBOOLEANVALUE   OfferMappingErrorType = "UNEXPECTED_BOOLEAN_VALUE"
	UNKNOWNCATEGORY          OfferMappingErrorType = "UNKNOWN_CATEGORY"
	UNKNOWNPARAMETER         OfferMappingErrorType = "UNKNOWN_PARAMETER"
)

func (e OfferMappingErrorType) Valid() bool {
	switch e {
	case EMPTYMARKETCATEGORY:
		return true
	case INVALIDCATEGORY:
		return true
	case INVALIDCOMMODITYCODE:
		return true
	case INVALIDGROUPIDCHARACTERS:
		return true
	case INVALIDGROUPIDLENGTH:
		return true
	case INVALIDPICKERURL:
		return true
	case INVALIDUNITID:
		return true
	case LOCKEDDIMENSIONS:
		return true
	case NUMBERFORMAT:
		return true
	case UNEXPECTEDBOOLEANVALUE:
		return true
	case UNKNOWNCATEGORY:
		return true
	case UNKNOWNPARAMETER:
		return true
	default:
		return false
	}
}

const (
	FINE   OfferSellingProgramStatusType = "FINE"
	REJECT OfferSellingProgramStatusType = "REJECT"
)

func (e OfferSellingProgramStatusType) Valid() bool {
	switch e {
	case FINE:
		return true
	case REJECT:
		return true
	default:
		return false
	}
}

func (e OfferType) Valid() bool {
	switch e {
	case OfferTypeALCOHOL:
		return true
	case OfferTypeARTISTTITLE:
		return true
	case OfferTypeAUDIOBOOK:
		return true
	case OfferTypeBOOK:
		return true
	case OfferTypeDEFAULT:
		return true
	case OfferTypeMEDICINE:
		return true
	case OfferTypeONDEMAND:
		return true
	default:
		return false
	}
}

func (e OperationStatusType) Valid() bool {
	switch e {
	case OperationStatusTypeDONE:
		return true
	case OperationStatusTypeFAILED:
		return true
	case OperationStatusTypeINPROGRESS:
		return true
	default:
		return false
	}
}

const (
	ORDERDELIVERYINTERVALUPDATE OperationType = "ORDER_DELIVERY_INTERVAL_UPDATE"
	ORDERRECIPIENTUPDATE        OperationType = "ORDER_RECIPIENT_UPDATE"
	ORDERSTATUSUPDATE           OperationType = "ORDER_STATUS_UPDATE"
	RETURNCANCELLATION          OperationType = "RETURN_CANCELLATION"
)

func (e OperationType) Valid() bool {
	switch e {
	case ORDERDELIVERYINTERVALUPDATE:
		return true
	case ORDERRECIPIENTUPDATE:
		return true
	case ORDERSTATUSUPDATE:
		return true
	case RETURNCANCELLATION:
		return true
	default:
		return false
	}
}

func (e OrderBuyerType) Valid() bool {
	switch e {
	case OrderBuyerTypeBUSINESS:
		return true
	case OrderBuyerTypePERSON:
		return true
	default:
		return false
	}
}

const (
	ORDERDELIVERED  OrderCancellationReasonType = "ORDER_DELIVERED"
	ORDERINDELIVERY OrderCancellationReasonType = "ORDER_IN_DELIVERY"
)

func (e OrderCancellationReasonType) Valid() bool {
	switch e {
	case ORDERDELIVERED:
		return true
	case ORDERINDELIVERY:
		return true
	default:
		return false
	}
}

const (
	PARTNERMOVEDDELIVERYDATES OrderDeliveryDateReasonType = "PARTNER_MOVED_DELIVERY_DATES"
	USERMOVEDDELIVERYDATES    OrderDeliveryDateReasonType = "USER_MOVED_DELIVERY_DATES"
)

func (e OrderDeliveryDateReasonType) Valid() bool {
	switch e {
	case PARTNERMOVEDDELIVERYDATES:
		return true
	case USERMOVEDDELIVERYDATES:
		return true
	default:
		return false
	}
}

func (e OrderDeliveryDispatchType) Valid() bool {
	switch e {
	case OrderDeliveryDispatchTypeBUYER:
		return true
	case OrderDeliveryDispatchTypeMARKETBRANDEDOUTLET:
		return true
	case OrderDeliveryDispatchTypeSHOPOUTLET:
		return true
	case OrderDeliveryDispatchTypeUNKNOWN:
		return true
	default:
		return false
	}
}

const (
	CHECKINGBYMERCHANT OrderDeliveryEacType = "CHECKING_BY_MERCHANT"
	COURIERTOMERCHANT  OrderDeliveryEacType = "COURIER_TO_MERCHANT"
	MERCHANTTOCOURIER  OrderDeliveryEacType = "MERCHANT_TO_COURIER"
)

func (e OrderDeliveryEacType) Valid() bool {
	switch e {
	case CHECKINGBYMERCHANT:
		return true
	case COURIERTOMERCHANT:
		return true
	case MERCHANTTOCOURIER:
		return true
	default:
		return false
	}
}

func (e OrderDeliveryPartnerType) Valid() bool {
	switch e {
	case OrderDeliveryPartnerTypeSHOP:
		return true
	case OrderDeliveryPartnerTypeUNKNOWN:
		return true
	case OrderDeliveryPartnerTypeYANDEXMARKET:
		return true
	default:
		return false
	}
}

func (e OrderDeliveryType) Valid() bool {
	switch e {
	case OrderDeliveryTypeDELIVERY:
		return true
	case OrderDeliveryTypeDIGITAL:
		return true
	case OrderDeliveryTypePICKUP:
		return true
	case OrderDeliveryTypePOST:
		return true
	case OrderDeliveryTypeUNKNOWN:
		return true
	default:
		return false
	}
}

const (
	NOTREADY OrderDocumentStatusType = "NOT_READY"
	READY    OrderDocumentStatusType = "READY"
)

func (e OrderDocumentStatusType) Valid() bool {
	switch e {
	case NOTREADY:
		return true
	case READY:
		return true
	default:
		return false
	}
}

const (
	CIS         OrderItemInstanceType = "CIS"
	CISOPTIONAL OrderItemInstanceType = "CIS_OPTIONAL"
	GTD         OrderItemInstanceType = "GTD"
	RNPT        OrderItemInstanceType = "RNPT"
	UIN         OrderItemInstanceType = "UIN"
)

func (e OrderItemInstanceType) Valid() bool {
	switch e {
	case CIS:
		return true
	case CISOPTIONAL:
		return true
	case GTD:
		return true
	case RNPT:
		return true
	case UIN:
		return true
	default:
		return false
	}
}

func (e OrderItemStatusType) Valid() bool {
	switch e {
	case OrderItemStatusTypeREJECTED:
		return true
	case OrderItemStatusTypeRETURNED:
		return true
	default:
		return false
	}
}

func (e OrderItemSubsidyType) Valid() bool {
	switch e {
	case OrderItemSubsidyTypeSUBSIDY:
		return true
	case OrderItemSubsidyTypeYANDEXCASHBACK:
		return true
	default:
		return false
	}
}

const (
	SAFETAG OrderItemTagType = "SAFE_TAG"
	ULTIMA  OrderItemTagType = "ULTIMA"
)

func (e OrderItemTagType) Valid() bool {
	switch e {
	case SAFETAG:
		return true
	case ULTIMA:
		return true
	default:
		return false
	}
}

const (
	PARTNERREQUESTEDREMOVE OrderItemsModificationRequestReasonType = "PARTNER_REQUESTED_REMOVE"
	USERREQUESTEDREMOVE    OrderItemsModificationRequestReasonType = "USER_REQUESTED_REMOVE"
)

func (e OrderItemsModificationRequestReasonType) Valid() bool {
	switch e {
	case PARTNERREQUESTEDREMOVE:
		return true
	case USERREQUESTEDREMOVE:
		return true
	default:
		return false
	}
}

func (e OrderLiftType) Valid() bool {
	switch e {
	case OrderLiftTypeCARGOELEVATOR:
		return true
	case OrderLiftTypeELEVATOR:
		return true
	case OrderLiftTypeFREE:
		return true
	case OrderLiftTypeMANUAL:
		return true
	case OrderLiftTypeNOTNEEDED:
		return true
	case OrderLiftTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e OrderPaymentMethodType) Valid() bool {
	switch e {
	case OrderPaymentMethodTypeAPPLEPAY:
		return true
	case OrderPaymentMethodTypeB2BACCOUNTPOSTPAYMENT:
		return true
	case OrderPaymentMethodTypeB2BACCOUNTPREPAYMENT:
		return true
	case OrderPaymentMethodTypeBNPLBANKONDELIVERY:
		return true
	case OrderPaymentMethodTypeBNPLONDELIVERY:
		return true
	case OrderPaymentMethodTypeBNPLTBYB:
		return true
	case OrderPaymentMethodTypeBOUNDCARDONDELIVERY:
		return true
	case OrderPaymentMethodTypeCARDONDELIVERY:
		return true
	case OrderPaymentMethodTypeCASHONDELIVERY:
		return true
	case OrderPaymentMethodTypeCREDIT:
		return true
	case OrderPaymentMethodTypeEXTERNALCERTIFICATE:
		return true
	case OrderPaymentMethodTypeGOOGLEPAY:
		return true
	case OrderPaymentMethodTypeMICROCREDIT:
		return true
	case OrderPaymentMethodTypeSBP:
		return true
	case OrderPaymentMethodTypeTINKOFFCREDIT:
		return true
	case OrderPaymentMethodTypeTINKOFFINSTALLMENTS:
		return true
	case OrderPaymentMethodTypeUNKNOWN:
		return true
	case OrderPaymentMethodTypeYANDEX:
		return true
	default:
		return false
	}
}

func (e OrderPaymentType) Valid() bool {
	switch e {
	case OrderPaymentTypePOSTPAID:
		return true
	case OrderPaymentTypePREPAID:
		return true
	case OrderPaymentTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e OrderPromoType) Valid() bool {
	switch e {
	case OrderPromoTypeBLUEFLASH:
		return true
	case OrderPromoTypeBLUESET:
		return true
	case OrderPromoTypeCASHBACK:
		return true
	case OrderPromoTypeCHEAPESTASGIFT:
		return true
	case OrderPromoTypeDCOEXTRADISCOUNT:
		return true
	case OrderPromoTypeDIRECTDISCOUNT:
		return true
	case OrderPromoTypeDISCOUNTBYPAYMENTTYPE:
		return true
	case OrderPromoTypeGENERICBUNDLE:
		return true
	case OrderPromoTypeMARKETBLUE:
		return true
	case OrderPromoTypeMARKETCOIN:
		return true
	case OrderPromoTypeMARKETCOUPON:
		return true
	case OrderPromoTypeMARKETPROMOCODE:
		return true
	case OrderPromoTypePERCENTDISCOUNT:
		return true
	case OrderPromoTypePRICEDROPASYOUSHOP:
		return true
	case OrderPromoTypeSECRETSALE:
		return true
	case OrderPromoTypeSPREADDISCOUNTCOUNT:
		return true
	case OrderPromoTypeSPREADDISCOUNTRECEIPT:
		return true
	case OrderPromoTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e OrderSourcePlatformType) Valid() bool {
	switch e {
	case OrderSourcePlatformTypeMARKET:
		return true
	case OrderSourcePlatformTypeOTHER:
		return true
	case OrderSourcePlatformTypeOZON:
		return true
	case OrderSourcePlatformTypeWILDBERRIES:
		return true
	default:
		return false
	}
}

func (e OrderStatsStatusType) Valid() bool {
	switch e {
	case OrderStatsStatusTypeCANCELLEDBEFOREPROCESSING:
		return true
	case OrderStatsStatusTypeCANCELLEDINDELIVERY:
		return true
	case OrderStatsStatusTypeCANCELLEDINPROCESSING:
		return true
	case OrderStatsStatusTypeDELIVERED:
		return true
	case OrderStatsStatusTypeDELIVERY:
		return true
	case OrderStatsStatusTypeLOST:
		return true
	case OrderStatsStatusTypePARTIALLYDELIVERED:
		return true
	case OrderStatsStatusTypePARTIALLYRETURNED:
		return true
	case OrderStatsStatusTypePENDING:
		return true
	case OrderStatsStatusTypePICKUP:
		return true
	case OrderStatsStatusTypePROCESSING:
		return true
	case OrderStatsStatusTypeRESERVED:
		return true
	case OrderStatsStatusTypeRETURNED:
		return true
	case OrderStatsStatusTypeUNKNOWN:
		return true
	case OrderStatsStatusTypeUNPAID:
		return true
	default:
		return false
	}
}

func (e OrderStatusType) Valid() bool {
	switch e {
	case OrderStatusTypeCANCELLED:
		return true
	case OrderStatusTypeDELIVERED:
		return true
	case OrderStatusTypeDELIVERY:
		return true
	case OrderStatusTypePARTIALLYRETURNED:
		return true
	case OrderStatusTypePENDING:
		return true
	case OrderStatusTypePICKUP:
		return true
	case OrderStatusTypePLACING:
		return true
	case OrderStatusTypePROCESSING:
		return true
	case OrderStatusTypeRESERVED:
		return true
	case OrderStatusTypeRETURNED:
		return true
	case OrderStatusTypeUNKNOWN:
		return true
	case OrderStatusTypeUNPAID:
		return true
	default:
		return false
	}
}

func (e OrderSubsidyType) Valid() bool {
	switch e {
	case OrderSubsidyTypeDELIVERY:
		return true
	case OrderSubsidyTypeSUBSIDY:
		return true
	case OrderSubsidyTypeYANDEXCASHBACK:
		return true
	default:
		return false
	}
}

func (e OrderSubstatusType) Valid() bool {
	switch e {
	case OrderSubstatusTypeANTIFRAUD:
		return true
	case OrderSubstatusTypeASPARTOFMULTIORDER:
		return true
	case OrderSubstatusTypeASYNCPROCESSING:
		return true
	case OrderSubstatusTypeAWAITAUTODELIVERYDATES:
		return true
	case OrderSubstatusTypeAWAITCASHIER:
		return true
	case OrderSubstatusTypeAWAITCONFIRMATION:
		return true
	case OrderSubstatusTypeAWAITCUSTOMPRICECONFIRMATION:
		return true
	case OrderSubstatusTypeAWAITDELIVERYDATES:
		return true
	case OrderSubstatusTypeAWAITDELIVERYDATESCONFIRMATION:
		return true
	case OrderSubstatusTypeAWAITLAVKARESERVATION:
		return true
	case OrderSubstatusTypeAWAITPAYMENT:
		return true
	case OrderSubstatusTypeAWAITPAYMENTAFTERDELIVERY:
		return true
	case OrderSubstatusTypeAWAITSERVICEABLECONFIRMATION:
		return true
	case OrderSubstatusTypeAWAITUSERPERSONALDATA:
		return true
	case OrderSubstatusTypeAWAITUSERSTEAMFASTURL:
		return true
	case OrderSubstatusTypeBANKREJECTCREDITOFFER:
		return true
	case OrderSubstatusTypeBROKENITEM:
		return true
	case OrderSubstatusTypeCANCELLEDCOURIERNOTFOUND:
		return true
	case OrderSubstatusTypeCOURIERARRIVEDTOSENDER:
		return true
	case OrderSubstatusTypeCOURIERFOUND:
		return true
	case OrderSubstatusTypeCOURIERINTRANSITTOSENDER:
		return true
	case OrderSubstatusTypeCOURIERNOTCOMEFORORDER:
		return true
	case OrderSubstatusTypeCOURIERNOTDELIVERORDER:
		return true
	case OrderSubstatusTypeCOURIERNOTFOUND:
		return true
	case OrderSubstatusTypeCOURIERRECEIVED:
		return true
	case OrderSubstatusTypeCOURIERRETURNEDORDER:
		return true
	case OrderSubstatusTypeCOURIERRETURNSORDER:
		return true
	case OrderSubstatusTypeCOURIERSEARCH:
		return true
	case OrderSubstatusTypeCOURIERSEARCHNOTSTARTED:
		return true
	case OrderSubstatusTypeCREDITOFFERFAILED:
		return true
	case OrderSubstatusTypeCUSTOM:
		return true
	case OrderSubstatusTypeCUSTOMERREJECTCREDITOFFER:
		return true
	case OrderSubstatusTypeCUSTOMSFAILEDMARKET:
		return true
	case OrderSubstatusTypeCUSTOMSFAILEDUSERADDITIONALDATANOTPROVIDED:
		return true
	case OrderSubstatusTypeCUSTOMSFAILEDUSERCOMMERCIALITEMS:
		return true
	case OrderSubstatusTypeCUSTOMSFAILEDUSERDUTYNOTPAID:
		return true
	case OrderSubstatusTypeCUSTOMSFAILEDUSERINVALIDPERSONALDATA:
		return true
	case OrderSubstatusTypeCUSTOMSPROBLEMS:
		return true
	case OrderSubstatusTypeDAMAGEDBOX:
		return true
	case OrderSubstatusTypeDEFERREDPAYMENT:
		return true
	case OrderSubstatusTypeDELIVEREDUSERNOTRECEIVED:
		return true
	case OrderSubstatusTypeDELIVEREDUSERRECEIVED:
		return true
	case OrderSubstatusTypeDELIVERYCUSTOMSARRIVED:
		return true
	case OrderSubstatusTypeDELIVERYCUSTOMSCLEARED:
		return true
	case OrderSubstatusTypeDELIVERYNOTMANAGEDREGION:
		return true
	case OrderSubstatusTypeDELIVERYPROBLEMS:
		return true
	case OrderSubstatusTypeDELIVERYSERVICEDELIVERED:
		return true
	case OrderSubstatusTypeDELIVERYSERVICEFAILED:
		return true
	case OrderSubstatusTypeDELIVERYSERVICELOST:
		return true
	case OrderSubstatusTypeDELIVERYSERVICENOTRECEIVED:
		return true
	case OrderSubstatusTypeDELIVERYSERVICERECEIVED:
		return true
	case OrderSubstatusTypeDELIVERYSERVICEUNDELIVERED:
		return true
	case OrderSubstatusTypeDELIVERYTOSTORESTARTED:
		return true
	case OrderSubstatusTypeDELIVERYUSERNOTRECEIVED:
		return true
	case OrderSubstatusTypeDROPOFFCLOSED:
		return true
	case OrderSubstatusTypeDROPOFFLOST:
		return true
	case OrderSubstatusTypeFIRSTMILEDELIVERYSERVICERECEIVED:
		return true
	case OrderSubstatusTypeFULLNOTRANSOM:
		return true
	case OrderSubstatusTypeINAPPROPRIATEWEIGHTSIZE:
		return true
	case OrderSubstatusTypeINCOMPLETECONTACTINFORMATION:
		return true
	case OrderSubstatusTypeINCOMPLETEMULTIORDER:
		return true
	case OrderSubstatusTypeINCORRECTPERSONALDATA:
		return true
	case OrderSubstatusTypeLASTMILECOURIERSEARCH:
		return true
	case OrderSubstatusTypeLASTMILESTARTED:
		return true
	case OrderSubstatusTypeLATECONTACT:
		return true
	case OrderSubstatusTypeLEGALINFOCHANGED:
		return true
	case OrderSubstatusTypeLOST:
		return true
	case OrderSubstatusTypeMISSINGITEM:
		return true
	case OrderSubstatusTypeNOPERSONALDATAEXPIRED:
		return true
	case OrderSubstatusTypePACKAGING:
		return true
	case OrderSubstatusTypePENDINGCANCELLED:
		return true
	case OrderSubstatusTypePENDINGEXPIRED:
		return true
	case OrderSubstatusTypePICKUPEXPIRED:
		return true
	case OrderSubstatusTypePICKUPPOINTCLOSED:
		return true
	case OrderSubstatusTypePICKUPSERVICERECEIVED:
		return true
	case OrderSubstatusTypePICKUPUSERRECEIVED:
		return true
	case OrderSubstatusTypePOSTPAIDBUDGETRESERVATIONFAILED:
		return true
	case OrderSubstatusTypePOSTPAIDFAILED:
		return true
	case OrderSubstatusTypePREORDER:
		return true
	case OrderSubstatusTypePRESCRIPTIONMISMATCH:
		return true
	case OrderSubstatusTypePROCESSINGEXPIRED:
		return true
	case OrderSubstatusTypeREADYFORLASTMILE:
		return true
	case OrderSubstatusTypeREADYFORPICKUP:
		return true
	case OrderSubstatusTypeREADYTOSHIP:
		return true
	case OrderSubstatusTypeREPLACINGORDER:
		return true
	case OrderSubstatusTypeRESERVATIONEXPIRED:
		return true
	case OrderSubstatusTypeRESERVATIONFAILED:
		return true
	case OrderSubstatusTypeSERVICEFAULT:
		return true
	case OrderSubstatusTypeSHIPPED:
		return true
	case OrderSubstatusTypeSHIPPEDTOWRONGDELIVERYSERVICE:
		return true
	case OrderSubstatusTypeSHOPFAILED:
		return true
	case OrderSubstatusTypeSHOPPENDINGCANCELLED:
		return true
	case OrderSubstatusTypeSORTINGCENTERLOST:
		return true
	case OrderSubstatusTypeSTARTED:
		return true
	case OrderSubstatusTypeTECHNICALERROR:
		return true
	case OrderSubstatusTypeTOOLONGDELIVERY:
		return true
	case OrderSubstatusTypeTOOMANYDELIVERYDATECHANGES:
		return true
	case OrderSubstatusTypeUNKNOWN:
		return true
	case OrderSubstatusTypeUSERBOUGHTCHEAPER:
		return true
	case OrderSubstatusTypeUSERCHANGEDMIND:
		return true
	case OrderSubstatusTypeUSERFORGOTTOUSEBONUS:
		return true
	case OrderSubstatusTypeUSERFRAUD:
		return true
	case OrderSubstatusTypeUSERHASNOTIMETOPICKUPORDER:
		return true
	case OrderSubstatusTypeUSERIDENTIFICATIONMISMATCH:
		return true
	case OrderSubstatusTypeUSERNOTPAID:
		return true
	case OrderSubstatusTypeUSERPLACEDOTHERORDER:
		return true
	case OrderSubstatusTypeUSERRECEIVED:
		return true
	case OrderSubstatusTypeUSERRECEIVEDTECHNICALERROR:
		return true
	case OrderSubstatusTypeUSERREFUSEDDELIVERY:
		return true
	case OrderSubstatusTypeUSERREFUSEDPRODUCT:
		return true
	case OrderSubstatusTypeUSERREFUSEDQUALITY:
		return true
	case OrderSubstatusTypeUSERUNREACHABLE:
		return true
	case OrderSubstatusTypeUSERWANTEDANOTHERPAYMENTMETHOD:
		return true
	case OrderSubstatusTypeUSERWANTSTOCHANGEADDRESS:
		return true
	case OrderSubstatusTypeUSERWANTSTOCHANGEDELIVERYDATE:
		return true
	case OrderSubstatusTypeWAITINGBANKDECISION:
		return true
	case OrderSubstatusTypeWAITINGFORSTOCKS:
		return true
	case OrderSubstatusTypeWAITINGPOSTPAIDBUDGETRESERVATION:
		return true
	case OrderSubstatusTypeWAITINGTINKOFFDECISION:
		return true
	case OrderSubstatusTypeWAITINGUSERDELIVERYINPUT:
		return true
	case OrderSubstatusTypeWAITINGUSERINPUT:
		return true
	case OrderSubstatusTypeWAREHOUSEFAILEDTOSHIP:
		return true
	case OrderSubstatusTypeWRONGITEM:
		return true
	case OrderSubstatusTypeWRONGITEMDELIVERED:
		return true
	default:
		return false
	}
}

func (e OrderTaxSystemType) Valid() bool {
	switch e {
	case OrderTaxSystemTypeAUSN:
		return true
	case OrderTaxSystemTypeAUSNMINUSCOST:
		return true
	case OrderTaxSystemTypeECHN:
		return true
	case OrderTaxSystemTypeENVD:
		return true
	case OrderTaxSystemTypeNPD:
		return true
	case OrderTaxSystemTypeOSN:
		return true
	case OrderTaxSystemTypePSN:
		return true
	case OrderTaxSystemTypeUNKNOWNVALUE:
		return true
	case OrderTaxSystemTypeUSN:
		return true
	case OrderTaxSystemTypeUSNMINUSCOST:
		return true
	default:
		return false
	}
}

func (e OrderUpdateStatusType) Valid() bool {
	switch e {
	case OrderUpdateStatusTypeERROR:
		return true
	case OrderUpdateStatusTypeOK:
		return true
	default:
		return false
	}
}

func (e OrderVatType) Valid() bool {
	switch e {
	case OrderVatTypeNOVAT:
		return true
	case OrderVatTypeUNKNOWNVALUE:
		return true
	case OrderVatTypeVAT0:
		return true
	case OrderVatTypeVAT05:
		return true
	case OrderVatTypeVAT07:
		return true
	case OrderVatTypeVAT10:
		return true
	case OrderVatTypeVAT10110:
		return true
	case OrderVatTypeVAT12:
		return true
	case OrderVatTypeVAT18:
		return true
	case OrderVatTypeVAT18118:
		return true
	case OrderVatTypeVAT20:
		return true
	case OrderVatTypeVAT20120:
		return true
	case OrderVatTypeVAT22:
		return true
	default:
		return false
	}
}

func (e OrdersStatsCommissionType) Valid() bool {
	switch e {
	case OrdersStatsCommissionTypeAGENCY:
		return true
	case OrdersStatsCommissionTypeAUCTIONPROMOTION:
		return true
	case OrdersStatsCommissionTypeCROSSREGIONALDELIVERY:
		return true
	case OrdersStatsCommissionTypeDELIVERYTOCUSTOMER:
		return true
	case OrdersStatsCommissionTypeEXPRESSDELIVERYTOCUSTOMER:
		return true
	case OrdersStatsCommissionTypeFEE:
		return true
	case OrdersStatsCommissionTypeFULFILLMENT:
		return true
	case OrdersStatsCommissionTypeILLIQUIDGOODSSALE:
		return true
	case OrdersStatsCommissionTypeINSTALLMENT:
		return true
	case OrdersStatsCommissionTypeINTAKESORTING:
		return true
	case OrdersStatsCommissionTypeLOYALTYPARTICIPATIONFEE:
		return true
	case OrdersStatsCommissionTypePAYMENTTRANSFER:
		return true
	case OrdersStatsCommissionTypeRETURNEDORDERSSTORAGE:
		return true
	case OrdersStatsCommissionTypeRETURNPROCESSING:
		return true
	case OrdersStatsCommissionTypeSORTING:
		return true
	default:
		return false
	}
}

func (e OrdersStatsItemStatusType) Valid() bool {
	switch e {
	case OrdersStatsItemStatusTypeREJECTED:
		return true
	case OrdersStatsItemStatusTypeRETURNED:
		return true
	default:
		return false
	}
}

func (e OrdersStatsOrderPaymentType) Valid() bool {
	switch e {
	case OrdersStatsOrderPaymentTypePOSTPAID:
		return true
	case OrdersStatsOrderPaymentTypePREPAID:
		return true
	case OrdersStatsOrderPaymentTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e OrdersStatsPaymentSourceType) Valid() bool {
	switch e {
	case OrdersStatsPaymentSourceTypeBUYER:
		return true
	case OrdersStatsPaymentSourceTypeCASHBACK:
		return true
	case OrdersStatsPaymentSourceTypeMARKETCESSION:
		return true
	case OrdersStatsPaymentSourceTypeMARKETPLACE:
		return true
	case OrdersStatsPaymentSourceTypeSPLIT:
		return true
	default:
		return false
	}
}

const (
	PAYMENT OrdersStatsPaymentType = "PAYMENT"
	REFUND  OrdersStatsPaymentType = "REFUND"
)

func (e OrdersStatsPaymentType) Valid() bool {
	switch e {
	case PAYMENT:
		return true
	case REFUND:
		return true
	default:
		return false
	}
}

func (e OrdersStatsPriceType) Valid() bool {
	switch e {
	case OrdersStatsPriceTypeBUYER:
		return true
	case OrdersStatsPriceTypeCASHBACK:
		return true
	case OrdersStatsPriceTypeMARKETPLACE:
		return true
	default:
		return false
	}
}

func (e OrdersStatsStockType) Valid() bool {
	switch e {
	case OrdersStatsStockTypeDEFECT:
		return true
	case OrdersStatsStockTypeEXPIRED:
		return true
	case OrdersStatsStockTypeFIT:
		return true
	default:
		return false
	}
}

const (
	ACCRUAL   OrdersStatsSubsidyOperationType = "ACCRUAL"
	DEDUCTION OrdersStatsSubsidyOperationType = "DEDUCTION"
)

func (e OrdersStatsSubsidyOperationType) Valid() bool {
	switch e {
	case ACCRUAL:
		return true
	case DEDUCTION:
		return true
	default:
		return false
	}
}

const (
	DELIVERY       OrdersStatsSubsidyType = "DELIVERY"
	SUBSIDY        OrdersStatsSubsidyType = "SUBSIDY"
	YANDEXCASHBACK OrdersStatsSubsidyType = "YANDEX_CASHBACK"
)

func (e OrdersStatsSubsidyType) Valid() bool {
	switch e {
	case DELIVERY:
		return true
	case SUBSIDY:
		return true
	case YANDEXCASHBACK:
		return true
	default:
		return false
	}
}

func (e OutletStatusType) Valid() bool {
	switch e {
	case OutletStatusTypeATMODERATION:
		return true
	case OutletStatusTypeFAILED:
		return true
	case OutletStatusTypeMODERATED:
		return true
	case OutletStatusTypeNONMODERATED:
		return true
	case OutletStatusTypeUNKNOWN:
		return true
	default:
		return false
	}
}

const (
	DEPOT      OutletType = "DEPOT"
	MIXED      OutletType = "MIXED"
	NOTDEFINED OutletType = "NOT_DEFINED"
	RETAIL     OutletType = "RETAIL"
)

func (e OutletType) Valid() bool {
	switch e {
	case DEPOT:
		return true
	case MIXED:
		return true
	case NOTDEFINED:
		return true
	case RETAIL:
		return true
	default:
		return false
	}
}

func (e OutletVisibilityType) Valid() bool {
	switch e {
	case OutletVisibilityTypeHIDDEN:
		return true
	case OutletVisibilityTypeUNKNOWN:
		return true
	case OutletVisibilityTypeVISIBLE:
		return true
	default:
		return false
	}
}

const (
	PageFormatTypeA4             PageFormatType = "A4"
	PageFormatTypeA7             PageFormatType = "A7"
	PageFormatTypeA9             PageFormatType = "A9"
	PageFormatTypeA9HORIZONTALLY PageFormatType = "A9_HORIZONTALLY"
)

func (e PageFormatType) Valid() bool {
	switch e {
	case PageFormatTypeA4:
		return true
	case PageFormatTypeA7:
		return true
	case PageFormatTypeA9:
		return true
	case PageFormatTypeA9HORIZONTALLY:
		return true
	default:
		return false
	}
}

const (
	BOOLEAN ParameterType = "BOOLEAN"
	ENUM    ParameterType = "ENUM"
	NUMERIC ParameterType = "NUMERIC"
	TEXT    ParameterType = "TEXT"
)

func (e ParameterType) Valid() bool {
	switch e {
	case BOOLEAN:
		return true
	case ENUM:
		return true
	case NUMERIC:
		return true
	case TEXT:
		return true
	default:
		return false
	}
}

const (
	BIWEEKLY PaymentFrequencyType = "BIWEEKLY"
	DAILY    PaymentFrequencyType = "DAILY"
	MONTHLY  PaymentFrequencyType = "MONTHLY"
	WEEKLY   PaymentFrequencyType = "WEEKLY"
)

func (e PaymentFrequencyType) Valid() bool {
	switch e {
	case BIWEEKLY:
		return true
	case DAILY:
		return true
	case MONTHLY:
		return true
	case WEEKLY:
		return true
	default:
		return false
	}
}

const (
	PlacementTypeDBS  PlacementType = "DBS"
	PlacementTypeFBS  PlacementType = "FBS"
	PlacementTypeFBY  PlacementType = "FBY"
	PlacementTypeLAAS PlacementType = "LAAS"
)

func (e PlacementType) Valid() bool {
	switch e {
	case PlacementTypeDBS:
		return true
	case PlacementTypeFBS:
		return true
	case PlacementTypeFBY:
		return true
	case PlacementTypeLAAS:
		return true
	default:
		return false
	}
}

func (e PriceCompetitivenessType) Valid() bool {
	switch e {
	case PriceCompetitivenessTypeAVERAGE:
		return true
	case PriceCompetitivenessTypeLOW:
		return true
	case PriceCompetitivenessTypeOPTIMAL:
		return true
	default:
		return false
	}
}

const (
	CURRENCY       PriceQuarantineVerdictParamNameType = "CURRENCY"
	CURRENTPRICE   PriceQuarantineVerdictParamNameType = "CURRENT_PRICE"
	LASTVALIDPRICE PriceQuarantineVerdictParamNameType = "LAST_VALID_PRICE"
	MINPRICE       PriceQuarantineVerdictParamNameType = "MIN_PRICE"
)

func (e PriceQuarantineVerdictParamNameType) Valid() bool {
	switch e {
	case CURRENCY:
		return true
	case CURRENTPRICE:
		return true
	case LASTVALIDPRICE:
		return true
	case MINPRICE:
		return true
	default:
		return false
	}
}

const (
	LOWPRICE      PriceQuarantineVerdictType = "LOW_PRICE"
	LOWPRICEPROMO PriceQuarantineVerdictType = "LOW_PRICE_PROMO"
	PRICECHANGE   PriceQuarantineVerdictType = "PRICE_CHANGE"
)

func (e PriceQuarantineVerdictType) Valid() bool {
	switch e {
	case LOWPRICE:
		return true
	case LOWPRICEPROMO:
		return true
	case PRICECHANGE:
		return true
	default:
		return false
	}
}

func (e PromoOfferParticipationStatusFilterType) Valid() bool {
	switch e {
	case PromoOfferParticipationStatusFilterTypeMANUALLYADDED:
		return true
	case PromoOfferParticipationStatusFilterTypeNOTMANUALLYADDED:
		return true
	default:
		return false
	}
}

func (e PromoOfferParticipationStatusMultiFilterType) Valid() bool {
	switch e {
	case PromoOfferParticipationStatusMultiFilterTypeMANUALLYADDED:
		return true
	case PromoOfferParticipationStatusMultiFilterTypeMINIMUMFORPROMOS:
		return true
	case PromoOfferParticipationStatusMultiFilterTypeNOTMANUALLYADDED:
		return true
	case PromoOfferParticipationStatusMultiFilterTypeRENEWED:
		return true
	case PromoOfferParticipationStatusMultiFilterTypeRENEWFAILED:
		return true
	default:
		return false
	}
}

const (
	AUTO             PromoOfferParticipationStatusType = "AUTO"
	MANUAL           PromoOfferParticipationStatusType = "MANUAL"
	MINIMUMFORPROMOS PromoOfferParticipationStatusType = "MINIMUM_FOR_PROMOS"
	NOTPARTICIPATING PromoOfferParticipationStatusType = "NOT_PARTICIPATING"
	PARTIALLYAUTO    PromoOfferParticipationStatusType = "PARTIALLY_AUTO"
	RENEWED          PromoOfferParticipationStatusType = "RENEWED"
	RENEWFAILED      PromoOfferParticipationStatusType = "RENEW_FAILED"
)

func (e PromoOfferParticipationStatusType) Valid() bool {
	switch e {
	case AUTO:
		return true
	case MANUAL:
		return true
	case MINIMUMFORPROMOS:
		return true
	case NOTPARTICIPATING:
		return true
	case PARTIALLYAUTO:
		return true
	case RENEWED:
		return true
	case RENEWFAILED:
		return true
	default:
		return false
	}
}

const (
	CATALOGPRICEISLOWERTHANPROMO PromoOfferUpdateWarningCodeType = "CATALOG_PRICE_IS_LOWER_THAN_PROMO"
	DEEPDISCOUNTOFFER            PromoOfferUpdateWarningCodeType = "DEEP_DISCOUNT_OFFER"
	SHOPOFFERNOTELIGIBLEFORPROMO PromoOfferUpdateWarningCodeType = "SHOP_OFFER_NOT_ELIGIBLE_FOR_PROMO"
	SHOPPRICESARELOWERTHANPROMO  PromoOfferUpdateWarningCodeType = "SHOP_PRICES_ARE_LOWER_THAN_PROMO"
)

func (e PromoOfferUpdateWarningCodeType) Valid() bool {
	switch e {
	case CATALOGPRICEISLOWERTHANPROMO:
		return true
	case DEEPDISCOUNTOFFER:
		return true
	case SHOPOFFERNOTELIGIBLEFORPROMO:
		return true
	case SHOPPRICESARELOWERTHANPROMO:
		return true
	default:
		return false
	}
}

const (
	PARTICIPATED     PromoParticipationType = "PARTICIPATED"
	PARTICIPATINGNOW PromoParticipationType = "PARTICIPATING_NOW"
)

func (e PromoParticipationType) Valid() bool {
	switch e {
	case PARTICIPATED:
		return true
	case PARTICIPATINGNOW:
		return true
	default:
		return false
	}
}

func (e QualityRatingComponentType) Valid() bool {
	switch e {
	case QualityRatingComponentTypeDBSCANCELLATIONRATE:
		return true
	case QualityRatingComponentTypeDBSLATEDELIVERYRATE:
		return true
	case QualityRatingComponentTypeFBSCANCELLATIONRATE:
		return true
	case QualityRatingComponentTypeFBSLATESHIPRATE:
		return true
	case QualityRatingComponentTypeFBYCANCELLATIONRATE:
		return true
	case QualityRatingComponentTypeFBYDELIVERYDIFFRATE:
		return true
	case QualityRatingComponentTypeFBYLATEDELIVERYRATE:
		return true
	case QualityRatingComponentTypeFBYLATEEDITINGRATE:
		return true
	default:
		return false
	}
}

const (
	CREATEDATASC  QuestionSortOrderType = "CREATED_AT_ASC"
	CREATEDATDESC QuestionSortOrderType = "CREATED_AT_DESC"
)

func (e QuestionSortOrderType) Valid() bool {
	switch e {
	case CREATEDATASC:
		return true
	case CREATEDATDESC:
		return true
	default:
		return false
	}
}

func (e QuestionsTextContentAuthorType) Valid() bool {
	switch e {
	case QuestionsTextContentAuthorTypeBRAND:
		return true
	case QuestionsTextContentAuthorTypeBUSINESS:
		return true
	case QuestionsTextContentAuthorTypeUSER:
		return true
	case QuestionsTextContentAuthorTypeVENDOR:
		return true
	default:
		return false
	}
}

func (e QuestionsTextContentModerationStatusType) Valid() bool {
	switch e {
	case QuestionsTextContentModerationStatusTypeBANNED:
		return true
	case QuestionsTextContentModerationStatusTypeDELETED:
		return true
	case QuestionsTextContentModerationStatusTypePUBLISHED:
		return true
	case QuestionsTextContentModerationStatusTypeUNMODERATED:
		return true
	default:
		return false
	}
}

const (
	CREATE QuestionsTextEntityOperationType = "CREATE"
	DELETE QuestionsTextEntityOperationType = "DELETE"
	UPDATE QuestionsTextEntityOperationType = "UPDATE"
)

func (e QuestionsTextEntityOperationType) Valid() bool {
	switch e {
	case CREATE:
		return true
	case DELETE:
		return true
	case UPDATE:
		return true
	default:
		return false
	}
}

const (
	ANSWER   QuestionsTextEntityType = "ANSWER"
	COMMENT  QuestionsTextEntityType = "COMMENT"
	QUESTION QuestionsTextEntityType = "QUESTION"
)

func (e QuestionsTextEntityType) Valid() bool {
	switch e {
	case ANSWER:
		return true
	case COMMENT:
		return true
	case QUESTION:
		return true
	default:
		return false
	}
}

const (
	DELIVERYSERVICE RecipientType = "DELIVERY_SERVICE"
	POST            RecipientType = "POST"
	SHOP            RecipientType = "SHOP"
)

func (e RecipientType) Valid() bool {
	switch e {
	case DELIVERYSERVICE:
		return true
	case POST:
		return true
	case SHOP:
		return true
	default:
		return false
	}
}

const (
	RefundStatusTypeCANCELLED                    RefundStatusType = "CANCELLED"
	RefundStatusTypeCOMPLETEWITHOUTREFUND        RefundStatusType = "COMPLETE_WITHOUT_REFUND"
	RefundStatusTypeDECISIONMADE                 RefundStatusType = "DECISION_MADE"
	RefundStatusTypeFAILED                       RefundStatusType = "FAILED"
	RefundStatusTypePREMODERATIONDECISIONMADE    RefundStatusType = "PREMODERATION_DECISION_MADE"
	RefundStatusTypePREMODERATIONDECISIONWAITING RefundStatusType = "PREMODERATION_DECISION_WAITING"
	RefundStatusTypePREMODERATIONDISPUTE         RefundStatusType = "PREMODERATION_DISPUTE"
	RefundStatusTypePREMODERATIONSELECTDELIVERY  RefundStatusType = "PREMODERATION_SELECT_DELIVERY"
	RefundStatusTypeREFUNDED                     RefundStatusType = "REFUNDED"
	RefundStatusTypeREFUNDEDBYSHOP               RefundStatusType = "REFUNDED_BY_SHOP"
	RefundStatusTypeREFUNDEDWITHBONUSES          RefundStatusType = "REFUNDED_WITH_BONUSES"
	RefundStatusTypeREFUNDINPROGRESS             RefundStatusType = "REFUND_IN_PROGRESS"
	RefundStatusTypeREJECTED                     RefundStatusType = "REJECTED"
	RefundStatusTypeSTARTEDBYUSER                RefundStatusType = "STARTED_BY_USER"
	RefundStatusTypeUNKNOWN                      RefundStatusType = "UNKNOWN"
	RefundStatusTypeWAITINGFORDECISION           RefundStatusType = "WAITING_FOR_DECISION"
)

func (e RefundStatusType) Valid() bool {
	switch e {
	case RefundStatusTypeCANCELLED:
		return true
	case RefundStatusTypeCOMPLETEWITHOUTREFUND:
		return true
	case RefundStatusTypeDECISIONMADE:
		return true
	case RefundStatusTypeFAILED:
		return true
	case RefundStatusTypePREMODERATIONDECISIONMADE:
		return true
	case RefundStatusTypePREMODERATIONDECISIONWAITING:
		return true
	case RefundStatusTypePREMODERATIONDISPUTE:
		return true
	case RefundStatusTypePREMODERATIONSELECTDELIVERY:
		return true
	case RefundStatusTypeREFUNDED:
		return true
	case RefundStatusTypeREFUNDEDBYSHOP:
		return true
	case RefundStatusTypeREFUNDEDWITHBONUSES:
		return true
	case RefundStatusTypeREFUNDINPROGRESS:
		return true
	case RefundStatusTypeREJECTED:
		return true
	case RefundStatusTypeSTARTEDBYUSER:
		return true
	case RefundStatusTypeUNKNOWN:
		return true
	case RefundStatusTypeWAITINGFORDECISION:
		return true
	default:
		return false
	}
}

func (e RegionType) Valid() bool {
	switch e {
	case RegionTypeCITY:
		return true
	case RegionTypeCITYDISTRICT:
		return true
	case RegionTypeCONTINENT:
		return true
	case RegionTypeCOUNTRY:
		return true
	case RegionTypeCOUNTRYDISTRICT:
		return true
	case RegionTypeOTHER:
		return true
	case RegionTypeREGION:
		return true
	case RegionTypeREPUBLIC:
		return true
	case RegionTypeREPUBLICAREA:
		return true
	case RegionTypeSUBWAYSTATION:
		return true
	case RegionTypeVILLAGE:
		return true
	default:
		return false
	}
}

func (e RejectedPromoOfferDeleteReasonType) Valid() bool {
	switch e {
	case RejectedPromoOfferDeleteReasonTypeOFFERDOESNOTEXIST:
		return true
	default:
		return false
	}
}

func (e RejectedPromoOfferUpdateReasonType) Valid() bool {
	switch e {
	case RejectedPromoOfferUpdateReasonTypeDEADLINEFORFOCUSPROMOSEXCEEDED:
		return true
	case RejectedPromoOfferUpdateReasonTypeEMPTYOLDPRICE:
		return true
	case RejectedPromoOfferUpdateReasonTypeEMPTYPROMOPRICE:
		return true
	case RejectedPromoOfferUpdateReasonTypeMAXPROMOPRICEEXCEEDED:
		return true
	case RejectedPromoOfferUpdateReasonTypeOFFERDOESNOTEXIST:
		return true
	case RejectedPromoOfferUpdateReasonTypeOFFERDUPLICATION:
		return true
	case RejectedPromoOfferUpdateReasonTypeOFFERNOTELIGIBLEFORPROMO:
		return true
	case RejectedPromoOfferUpdateReasonTypeOFFERPROMOSMAXBYTESIZEEXCEEDED:
		return true
	case RejectedPromoOfferUpdateReasonTypeOLDPRICETOOBIG:
		return true
	case RejectedPromoOfferUpdateReasonTypePRICETOOBIG:
		return true
	case RejectedPromoOfferUpdateReasonTypePROMOPRICEBIGGERTHANMAX:
		return true
	case RejectedPromoOfferUpdateReasonTypePROMOPRICESMALLERTHANMIN:
		return true
	default:
		return false
	}
}

const (
	CSV  ReportFormatType = "CSV"
	FILE ReportFormatType = "FILE"
	JSON ReportFormatType = "JSON"
)

func (e ReportFormatType) Valid() bool {
	switch e {
	case CSV:
		return true
	case FILE:
		return true
	case JSON:
		return true
	default:
		return false
	}
}

const (
	EN ReportLanguageType = "EN"
	RU ReportLanguageType = "RU"
)

func (e ReportLanguageType) Valid() bool {
	switch e {
	case EN:
		return true
	case RU:
		return true
	default:
		return false
	}
}

func (e ReportStatusType) Valid() bool {
	switch e {
	case ReportStatusTypeDONE:
		return true
	case ReportStatusTypeFAILED:
		return true
	case ReportStatusTypePENDING:
		return true
	case ReportStatusTypePROCESSING:
		return true
	default:
		return false
	}
}

const (
	NODATA           ReportSubStatusType = "NO_DATA"
	RESOURCENOTFOUND ReportSubStatusType = "RESOURCE_NOT_FOUND"
	TOOLARGE         ReportSubStatusType = "TOO_LARGE"
)

func (e ReportSubStatusType) Valid() bool {
	switch e {
	case NODATA:
		return true
	case RESOURCENOTFOUND:
		return true
	case TOOLARGE:
		return true
	default:
		return false
	}
}

func (e ReturnDecisionReasonType) Valid() bool {
	switch e {
	case ReturnDecisionReasonTypeBADQUALITY:
		return true
	case ReturnDecisionReasonTypeCONTENTFAIL:
		return true
	case ReturnDecisionReasonTypeDAMAGEDELIVERY:
		return true
	case ReturnDecisionReasonTypeDELIVERYFAIL:
		return true
	case ReturnDecisionReasonTypeDOESNOTFIT:
		return true
	case ReturnDecisionReasonTypeLOYALTYFAIL:
		return true
	case ReturnDecisionReasonTypeUNKNOWN:
		return true
	case ReturnDecisionReasonTypeWRONGITEM:
		return true
	default:
		return false
	}
}

func (e ReturnDecisionSubreasonType) Valid() bool {
	switch e {
	case ReturnDecisionSubreasonTypeBADFLOWERS:
		return true
	case ReturnDecisionSubreasonTypeBADPACKAGE:
		return true
	case ReturnDecisionSubreasonTypeBROKEN:
		return true
	case ReturnDecisionSubreasonTypeDAMAGED:
		return true
	case ReturnDecisionSubreasonTypeDELIVEREDTOOLONG:
		return true
	case ReturnDecisionSubreasonTypeDIDNOTMATCHDESCRIPTION:
		return true
	case ReturnDecisionSubreasonTypeINCOMPLETE:
		return true
	case ReturnDecisionSubreasonTypeINCOMPLETENESS:
		return true
	case ReturnDecisionSubreasonTypeITEMWASUSED:
		return true
	case ReturnDecisionSubreasonTypeNOTWORKING:
		return true
	case ReturnDecisionSubreasonTypePARCELMISSING:
		return true
	case ReturnDecisionSubreasonTypeUNKNOWN:
		return true
	case ReturnDecisionSubreasonTypeUSERCHANGEDMIND:
		return true
	case ReturnDecisionSubreasonTypeUSERDIDNOTLIKE:
		return true
	case ReturnDecisionSubreasonTypeWRAPPINGDAMAGED:
		return true
	case ReturnDecisionSubreasonTypeWRONGAMOUNTDELIVERED:
		return true
	case ReturnDecisionSubreasonTypeWRONGCOLOR:
		return true
	case ReturnDecisionSubreasonTypeWRONGITEM:
		return true
	case ReturnDecisionSubreasonTypeWRONGORDER:
		return true
	default:
		return false
	}
}

func (e ReturnDecisionType) Valid() bool {
	switch e {
	case ReturnDecisionTypeDECLINEREFUND:
		return true
	case ReturnDecisionTypeFASTREFUNDMONEY:
		return true
	case ReturnDecisionTypeOTHERDECISION:
		return true
	case ReturnDecisionTypePARTIALMONEYREFUND:
		return true
	case ReturnDecisionTypeREFUNDMONEY:
		return true
	case ReturnDecisionTypeREFUNDMONEYINCLUDINGSHIPMENT:
		return true
	case ReturnDecisionTypeREPAIR:
		return true
	case ReturnDecisionTypeREPLACE:
		return true
	case ReturnDecisionTypeSENDTOEXAMINATION:
		return true
	case ReturnDecisionTypeUNKNOWN:
		return true
	default:
		return false
	}
}

func (e ReturnInstanceStatusType) Valid() bool {
	switch e {
	case ReturnInstanceStatusTypeCANCELLED:
		return true
	case ReturnInstanceStatusTypeCREATED:
		return true
	case ReturnInstanceStatusTypeEXPROPRIATED:
		return true
	case ReturnInstanceStatusTypeINTRANSIT:
		return true
	case ReturnInstanceStatusTypeLOST:
		return true
	case ReturnInstanceStatusTypeNOTINDEMAND:
		return true
	case ReturnInstanceStatusTypePICKED:
		return true
	case ReturnInstanceStatusTypePREPAREDFORUTILIZATION:
		return true
	case ReturnInstanceStatusTypeREADYFORPICKUP:
		return true
	case ReturnInstanceStatusTypeRECEIVED:
		return true
	case ReturnInstanceStatusTypeRECEIVEDONFULFILLMENT:
		return true
	case ReturnInstanceStatusTypeUTILIZED:
		return true
	default:
		return false
	}
}

func (e ReturnInstanceStockType) Valid() bool {
	switch e {
	case ReturnInstanceStockTypeANOMALY:
		return true
	case ReturnInstanceStockTypeDEFECT:
		return true
	case ReturnInstanceStockTypeDEMO:
		return true
	case ReturnInstanceStockTypeEXPIRED:
		return true
	case ReturnInstanceStockTypeFIRMWARE:
		return true
	case ReturnInstanceStockTypeFIT:
		return true
	case ReturnInstanceStockTypeINCORRECTCIS:
		return true
	case ReturnInstanceStockTypeINCORRECTIMEI:
		return true
	case ReturnInstanceStockTypeINCORRECTSERIALNUMBER:
		return true
	case ReturnInstanceStockTypeMARKDOWN:
		return true
	case ReturnInstanceStockTypeMISGRADING:
		return true
	case ReturnInstanceStockTypeNONCOMPLIENT:
		return true
	case ReturnInstanceStockTypeNOTACCEPTABLE:
		return true
	case ReturnInstanceStockTypePARTMISSING:
		return true
	case ReturnInstanceStockTypeREPAIR:
		return true
	case ReturnInstanceStockTypeSERVICE:
		return true
	case ReturnInstanceStockTypeSURPLUS:
		return true
	case ReturnInstanceStockTypeUNDEFINED:
		return true
	case ReturnInstanceStockTypeUNKNOWN:
		return true
	default:
		return false
	}
}

const (
	CONFIGURATIONORPACKAGINGCOMPROMISED ReturnRequestDecisionReasonType = "CONFIGURATION_OR_PACKAGING_COMPROMISED"
	DEVICEACTIVATED                     ReturnRequestDecisionReasonType = "DEVICE_ACTIVATED"
	ISSUEWITHTHEPRODUCTWASNOTCONFIRMED  ReturnRequestDecisionReasonType = "ISSUE_WITH_THE_PRODUCT_WAS_NOT_CONFIRMED"
	MECHANICALDAMAGE                    ReturnRequestDecisionReasonType = "MECHANICAL_DAMAGE"
	PRODUCTAPPEARANCECOMPROMISED        ReturnRequestDecisionReasonType = "PRODUCT_APPEARANCE_COMPROMISED"
	WARRANTYPERIODHASEXPIRED            ReturnRequestDecisionReasonType = "WARRANTY_PERIOD_HAS_EXPIRED"
	WARRANTYTERMSVIOLATED               ReturnRequestDecisionReasonType = "WARRANTY_TERMS_VIOLATED"
)

func (e ReturnRequestDecisionReasonType) Valid() bool {
	switch e {
	case CONFIGURATIONORPACKAGINGCOMPROMISED:
		return true
	case DEVICEACTIVATED:
		return true
	case ISSUEWITHTHEPRODUCTWASNOTCONFIRMED:
		return true
	case MECHANICALDAMAGE:
		return true
	case PRODUCTAPPEARANCECOMPROMISED:
		return true
	case WARRANTYPERIODHASEXPIRED:
		return true
	case WARRANTYTERMSVIOLATED:
		return true
	default:
		return false
	}
}

const (
	DECLINEREFUND                ReturnRequestDecisionType = "DECLINE_REFUND"
	FASTREFUNDMONEY              ReturnRequestDecisionType = "FAST_REFUND_MONEY"
	OTHERDECISION                ReturnRequestDecisionType = "OTHER_DECISION"
	PARTIALMONEYREFUND           ReturnRequestDecisionType = "PARTIAL_MONEY_REFUND"
	REFUNDMONEY                  ReturnRequestDecisionType = "REFUND_MONEY"
	REFUNDMONEYINCLUDINGSHIPMENT ReturnRequestDecisionType = "REFUND_MONEY_INCLUDING_SHIPMENT"
	REPAIR                       ReturnRequestDecisionType = "REPAIR"
	REPLACE                      ReturnRequestDecisionType = "REPLACE"
	SENDTOEXAMINATION            ReturnRequestDecisionType = "SEND_TO_EXAMINATION"
)

func (e ReturnRequestDecisionType) Valid() bool {
	switch e {
	case DECLINEREFUND:
		return true
	case FASTREFUNDMONEY:
		return true
	case OTHERDECISION:
		return true
	case PARTIALMONEYREFUND:
		return true
	case REFUNDMONEY:
		return true
	case REFUNDMONEYINCLUDINGSHIPMENT:
		return true
	case REPAIR:
		return true
	case REPLACE:
		return true
	case SENDTOEXAMINATION:
		return true
	default:
		return false
	}
}

func (e ReturnShipmentStatusType) Valid() bool {
	switch e {
	case ReturnShipmentStatusTypeCANCELLED:
		return true
	case ReturnShipmentStatusTypeCREATED:
		return true
	case ReturnShipmentStatusTypeEXPIRED:
		return true
	case ReturnShipmentStatusTypeFULFILMENTRECEIVED:
		return true
	case ReturnShipmentStatusTypeINTRANSIT:
		return true
	case ReturnShipmentStatusTypeLOST:
		return true
	case ReturnShipmentStatusTypeNOTINDEMAND:
		return true
	case ReturnShipmentStatusTypePICKED:
		return true
	case ReturnShipmentStatusTypePREPAREDFORUTILIZATION:
		return true
	case ReturnShipmentStatusTypeREADYFOREXPROPRIATION:
		return true
	case ReturnShipmentStatusTypeREADYFORPICKUP:
		return true
	case ReturnShipmentStatusTypeRECEIVED:
		return true
	case ReturnShipmentStatusTypeRECEIVEDFOREXPROPRIATION:
		return true
	case ReturnShipmentStatusTypeUNKNOWN:
		return true
	case ReturnShipmentStatusTypeUTILIZED:
		return true
	default:
		return false
	}
}

const (
	RETURN     ReturnType = "RETURN"
	UNREDEEMED ReturnType = "UNREDEEMED"
)

func (e ReturnType) Valid() bool {
	switch e {
	case RETURN:
		return true
	case UNREDEEMED:
		return true
	default:
		return false
	}
}

const (
	SellingProgramTypeDBS     SellingProgramType = "DBS"
	SellingProgramTypeEXPRESS SellingProgramType = "EXPRESS"
	SellingProgramTypeFBS     SellingProgramType = "FBS"
	SellingProgramTypeFBY     SellingProgramType = "FBY"
	SellingProgramTypeLAAS    SellingProgramType = "LAAS"
)

func (e SellingProgramType) Valid() bool {
	switch e {
	case SellingProgramTypeDBS:
		return true
	case SellingProgramTypeEXPRESS:
		return true
	case SellingProgramTypeFBS:
		return true
	case SellingProgramTypeFBY:
		return true
	case SellingProgramTypeLAAS:
		return true
	default:
		return false
	}
}

const (
	CHANGEPALLETSCOUNT            ShipmentActionType = "CHANGE_PALLETS_COUNT"
	CONFIRM                       ShipmentActionType = "CONFIRM"
	DOWNLOADACT                   ShipmentActionType = "DOWNLOAD_ACT"
	DOWNLOADDISCREPANCYACT        ShipmentActionType = "DOWNLOAD_DISCREPANCY_ACT"
	DOWNLOADINBOUNDACT            ShipmentActionType = "DOWNLOAD_INBOUND_ACT"
	DOWNLOADTRANSPORTATIONWAYBILL ShipmentActionType = "DOWNLOAD_TRANSPORTATION_WAYBILL"
)

func (e ShipmentActionType) Valid() bool {
	switch e {
	case CHANGEPALLETSCOUNT:
		return true
	case CONFIRM:
		return true
	case DOWNLOADACT:
		return true
	case DOWNLOADDISCREPANCYACT:
		return true
	case DOWNLOADINBOUNDACT:
		return true
	case DOWNLOADTRANSPORTATIONWAYBILL:
		return true
	default:
		return false
	}
}

func (e ShipmentPalletLabelPageFormatType) Valid() bool {
	switch e {
	case ShipmentPalletLabelPageFormatTypeA4:
		return true
	case ShipmentPalletLabelPageFormatTypeA8:
		return true
	default:
		return false
	}
}

func (e ShipmentStatusType) Valid() bool {
	switch e {
	case ShipmentStatusTypeACCEPTED:
		return true
	case ShipmentStatusTypeACCEPTEDWITHDISCREPANCIES:
		return true
	case ShipmentStatusTypeERROR:
		return true
	case ShipmentStatusTypeFINISHED:
		return true
	case ShipmentStatusTypeOUTBOUNDCONFIRMED:
		return true
	case ShipmentStatusTypeOUTBOUNDCREATED:
		return true
	case ShipmentStatusTypeOUTBOUNDREADYFORCONFIRMATION:
		return true
	case ShipmentStatusTypeOUTBOUNDSIGNED:
		return true
	default:
		return false
	}
}

func (e ShipmentType) Valid() bool {
	switch e {
	case ShipmentTypeIMPORT:
		return true
	case ShipmentTypeWITHDRAW:
		return true
	default:
		return false
	}
}

const (
	B2B ShowcaseType = "B2B"
	B2C ShowcaseType = "B2C"
)

func (e ShowcaseType) Valid() bool {
	switch e {
	case B2B:
		return true
	case B2C:
		return true
	default:
		return false
	}
}

const (
	CATEGORIES ShowsSalesGroupingType = "CATEGORIES"
	OFFERS     ShowsSalesGroupingType = "OFFERS"
)

func (e ShowsSalesGroupingType) Valid() bool {
	switch e {
	case CATEGORIES:
		return true
	case OFFERS:
		return true
	default:
		return false
	}
}

const (
	ASC  SortOrderType = "ASC"
	DESC SortOrderType = "DESC"
)

func (e SortOrderType) Valid() bool {
	switch e {
	case ASC:
		return true
	case DESC:
		return true
	default:
		return false
	}
}

const (
	CLICKS StatisticsAttributionType = "CLICKS"
	SHOWS  StatisticsAttributionType = "SHOWS"
)

func (e StatisticsAttributionType) Valid() bool {
	switch e {
	case CLICKS:
		return true
	case SHOWS:
		return true
	default:
		return false
	}
}

func (e SupplyRequestDocumentType) Valid() bool {
	switch e {
	case SupplyRequestDocumentTypeACTOFDISCREPANCY:
		return true
	case SupplyRequestDocumentTypeACTOFRECEPTIONTRANSFER:
		return true
	case SupplyRequestDocumentTypeACTOFWITHDRAW:
		return true
	case SupplyRequestDocumentTypeACTOFWITHDRAWFROMSTORAGE:
		return true
	case SupplyRequestDocumentTypeADDITIONALSUPPLY:
		return true
	case SupplyRequestDocumentTypeADDITIONALSUPPLYACCEPTABLEGOODS:
		return true
	case SupplyRequestDocumentTypeADDITIONALSUPPLYUNACCEPTABLEGOODS:
		return true
	case SupplyRequestDocumentTypeANOMALYCONTAINERSWITHDRAWACT:
		return true
	case SupplyRequestDocumentTypeCARGOUNITS:
		return true
	case SupplyRequestDocumentTypeCISFACT:
		return true
	case SupplyRequestDocumentTypeIDENTIFIERS:
		return true
	case SupplyRequestDocumentTypeINBOUNDUTD:
		return true
	case SupplyRequestDocumentTypeITEMSWITHCISES:
		return true
	case SupplyRequestDocumentTypeOUTBOUNDUTD:
		return true
	case SupplyRequestDocumentTypeREPORTOFWITHDRAWWITHCISES:
		return true
	case SupplyRequestDocumentTypeRNPTFACT:
		return true
	case SupplyRequestDocumentTypeSECONDARYACCEPTANCECISES:
		return true
	case SupplyRequestDocumentTypeSECONDARYRECEPTIONACT:
		return true
	case SupplyRequestDocumentTypeSUPPLY:
		return true
	case SupplyRequestDocumentTypeTRANSFER:
		return true
	case SupplyRequestDocumentTypeVALIDATIONERRORS:
		return true
	case SupplyRequestDocumentTypeVIRTUALDISTRIBUTIONCENTERSUPPLY:
		return true
	case SupplyRequestDocumentTypeWITHDRAW:
		return true
	default:
		return false
	}
}

func (e SupplyRequestLocationType) Valid() bool {
	switch e {
	case SupplyRequestLocationTypeFULFILLMENT:
		return true
	case SupplyRequestLocationTypePICKUPPOINT:
		return true
	case SupplyRequestLocationTypeXDOC:
		return true
	default:
		return false
	}
}

func (e SupplyRequestReferenceType) Valid() bool {
	switch e {
	case SupplyRequestReferenceTypeADDITIONALSUPPLY:
		return true
	case SupplyRequestReferenceTypeUTILIZATION:
		return true
	case SupplyRequestReferenceTypeVIRTUALDISTRIBUTION:
		return true
	case SupplyRequestReferenceTypeWITHDRAW:
		return true
	default:
		return false
	}
}

func (e SupplyRequestSortAttributeType) Valid() bool {
	switch e {
	case SupplyRequestSortAttributeTypeID:
		return true
	case SupplyRequestSortAttributeTypeREQUESTEDDATE:
		return true
	case SupplyRequestSortAttributeTypeSTATUS:
		return true
	case SupplyRequestSortAttributeTypeUPDATEDAT:
		return true
	default:
		return false
	}
}

func (e SupplyRequestStatusType) Valid() bool {
	switch e {
	case SupplyRequestStatusTypeACCEPTEDBYWAREHOUSESYSTEM:
		return true
	case SupplyRequestStatusTypeARRIVEDTOSERVICE:
		return true
	case SupplyRequestStatusTypeARRIVEDTOXDOCSERVICE:
		return true
	case SupplyRequestStatusTypeCANCELLATIONREJECTED:
		return true
	case SupplyRequestStatusTypeCANCELLATIONREQUESTED:
		return true
	case SupplyRequestStatusTypeCANCELLED:
		return true
	case SupplyRequestStatusTypeCREATED:
		return true
	case SupplyRequestStatusTypeFINISHED:
		return true
	case SupplyRequestStatusTypeINVALID:
		return true
	case SupplyRequestStatusTypeNEEDPREPARATION:
		return true
	case SupplyRequestStatusTypePUBLISHED:
		return true
	case SupplyRequestStatusTypeREADYFORUTILIZATION:
		return true
	case SupplyRequestStatusTypeREADYTOWITHDRAW:
		return true
	case SupplyRequestStatusTypeREGISTEREDINELECTRONICQUEUE:
		return true
	case SupplyRequestStatusTypeSHIPPEDTOSERVICE:
		return true
	case SupplyRequestStatusTypeTRANSITMOVING:
		return true
	case SupplyRequestStatusTypeVALIDATED:
		return true
	case SupplyRequestStatusTypeWAREHOUSEHANDLING:
		return true
	case SupplyRequestStatusTypeWAREHOUSESIGNEDACT:
		return true
	default:
		return false
	}
}

func (e SupplyRequestSubType) Valid() bool {
	switch e {
	case SupplyRequestSubTypeADDITIONALSUPPLY:
		return true
	case SupplyRequestSubTypeANOMALYWITHDRAW:
		return true
	case SupplyRequestSubTypeDEFAULT:
		return true
	case SupplyRequestSubTypeEXTERNALWITHDRAWINTOZON:
		return true
	case SupplyRequestSubTypeEXTERNALWITHDRAWINTWB:
		return true
	case SupplyRequestSubTypeFIXLOSTINVENTORYING:
		return true
	case SupplyRequestSubTypeFORCEPLAN:
		return true
	case SupplyRequestSubTypeFORCEPLANANOMALYPERSUPPLY:
		return true
	case SupplyRequestSubTypeINVENTORYINGSUPPLY:
		return true
	case SupplyRequestSubTypeINVENTORYINGSUPPLYWAREHOUSEBASEDPERSUPPLIER:
		return true
	case SupplyRequestSubTypeMANUTIL:
		return true
	case SupplyRequestSubTypeMISGRADINGSUPPLY:
		return true
	case SupplyRequestSubTypeMISGRADINGWITHDRAW:
		return true
	case SupplyRequestSubTypeMOVEMENTSUPPLY:
		return true
	case SupplyRequestSubTypeMOVEMENTWITHDRAW:
		return true
	case SupplyRequestSubTypeOPERLOSTINVENTORYING:
		return true
	case SupplyRequestSubTypePLANBYSUPPLIER:
		return true
	case SupplyRequestSubTypeVIRTUALDISTRIBUTIONCENTER:
		return true
	case SupplyRequestSubTypeVIRTUALDISTRIBUTIONCENTERCHILD:
		return true
	case SupplyRequestSubTypeWITHDRAWAUTOUTILIZATION:
		return true
	case SupplyRequestSubTypeXDOC:
		return true
	default:
		return false
	}
}

func (e SupplyRequestType) Valid() bool {
	switch e {
	case SupplyRequestTypeSUPPLY:
		return true
	case SupplyRequestTypeUTILIZATION:
		return true
	case SupplyRequestTypeWITHDRAW:
		return true
	default:
		return false
	}
}

const (
	AGENCYCOMMISSION            TariffType = "AGENCY_COMMISSION"
	CANCELLEDORDERFEEQI         TariffType = "CANCELLED_ORDER_FEE_QI"
	CARGOACCEPTANCE             TariffType = "CARGO_ACCEPTANCE"
	CROSSBORDERDELIVERY         TariffType = "CROSSBORDER_DELIVERY"
	CROSSREGIONALDELIVERY       TariffType = "CROSSREGIONAL_DELIVERY"
	CROSSREGIONALDELIVERYRETURN TariffType = "CROSSREGIONAL_DELIVERY_RETURN"
	DELIVERYTOCUSTOMER          TariffType = "DELIVERY_TO_CUSTOMER"
	DISPOSAL                    TariffType = "DISPOSAL"
	EXPRESSCANCELLEDBYPARTNER   TariffType = "EXPRESS_CANCELLED_BY_PARTNER"
	EXPRESSDELIVERY             TariffType = "EXPRESS_DELIVERY"
	FEE                         TariffType = "FEE"
	FFSTORAGEBILLING            TariffType = "FF_STORAGE_BILLING"
	FFXDOCSUPPLYBOX             TariffType = "FF_XDOC_SUPPLY_BOX"
	FFXDOCSUPPLYPALLET          TariffType = "FF_XDOC_SUPPLY_PALLET"
	GOODSACCEPTANCE             TariffType = "GOODS_ACCEPTANCE"
	INTAKESORTINGBULKYCARGO     TariffType = "INTAKE_SORTING_BULKY_CARGO"
	INTAKESORTINGDAILY          TariffType = "INTAKE_SORTING_DAILY"
	INTAKESORTINGSMALLGOODS     TariffType = "INTAKE_SORTING_SMALL_GOODS"
	LATEORDEREXECUTIONFEEQI     TariffType = "LATE_ORDER_EXECUTION_FEE_QI"
	MIDDLEMILE                  TariffType = "MIDDLE_MILE"
	ORDERPROCESSING             TariffType = "ORDER_PROCESSING"
	PAYMENTTRANSFER             TariffType = "PAYMENT_TRANSFER"
	RETURNPROCESSING            TariffType = "RETURN_PROCESSING"
	SORTING                     TariffType = "SORTING"
	SORTINGCENTERSTORAGE        TariffType = "SORTING_CENTER_STORAGE"
	STORAGE                     TariffType = "STORAGE"
	SURPLUS                     TariffType = "SURPLUS"
	VOLUMESTORAGE               TariffType = "VOLUME_STORAGE"
	WITHDRAW                    TariffType = "WITHDRAW"
	WITHDRAWEXTERNAL            TariffType = "WITHDRAW_EXTERNAL"
)

func (e TariffType) Valid() bool {
	switch e {
	case AGENCYCOMMISSION:
		return true
	case CANCELLEDORDERFEEQI:
		return true
	case CARGOACCEPTANCE:
		return true
	case CROSSBORDERDELIVERY:
		return true
	case CROSSREGIONALDELIVERY:
		return true
	case CROSSREGIONALDELIVERYRETURN:
		return true
	case DELIVERYTOCUSTOMER:
		return true
	case DISPOSAL:
		return true
	case EXPRESSCANCELLEDBYPARTNER:
		return true
	case EXPRESSDELIVERY:
		return true
	case FEE:
		return true
	case FFSTORAGEBILLING:
		return true
	case FFXDOCSUPPLYBOX:
		return true
	case FFXDOCSUPPLYPALLET:
		return true
	case GOODSACCEPTANCE:
		return true
	case INTAKESORTINGBULKYCARGO:
		return true
	case INTAKESORTINGDAILY:
		return true
	case INTAKESORTINGSMALLGOODS:
		return true
	case LATEORDEREXECUTIONFEEQI:
		return true
	case MIDDLEMILE:
		return true
	case ORDERPROCESSING:
		return true
	case PAYMENTTRANSFER:
		return true
	case RETURNPROCESSING:
		return true
	case SORTING:
		return true
	case SORTINGCENTERSTORAGE:
		return true
	case STORAGE:
		return true
	case SURPLUS:
		return true
	case VOLUMESTORAGE:
		return true
	case WITHDRAW:
		return true
	case WITHDRAWEXTERNAL:
		return true
	default:
		return false
	}
}

const (
	DAY   TimeUnitType = "DAY"
	HOUR  TimeUnitType = "HOUR"
	MONTH TimeUnitType = "MONTH"
	WEEK  TimeUnitType = "WEEK"
	YEAR  TimeUnitType = "YEAR"
)

func (e TimeUnitType) Valid() bool {
	switch e {
	case DAY:
		return true
	case HOUR:
		return true
	case MONTH:
		return true
	case WEEK:
		return true
	case YEAR:
		return true
	default:
		return false
	}
}

const (
	TurnoverTypeALMOSTLOW TurnoverType = "ALMOST_LOW"
	TurnoverTypeFREESTORE TurnoverType = "FREE_STORE"
	TurnoverTypeHIGH      TurnoverType = "HIGH"
	TurnoverTypeLOW       TurnoverType = "LOW"
	TurnoverTypeNOSALES   TurnoverType = "NO_SALES"
	TurnoverTypeVERYHIGH  TurnoverType = "VERY_HIGH"
)

func (e TurnoverType) Valid() bool {
	switch e {
	case TurnoverTypeALMOSTLOW:
		return true
	case TurnoverTypeFREESTORE:
		return true
	case TurnoverTypeHIGH:
		return true
	case TurnoverTypeLOW:
		return true
	case TurnoverTypeNOSALES:
		return true
	case TurnoverTypeVERYHIGH:
		return true
	default:
		return false
	}
}

const (
	UinStatusTypeFAILED          UinStatusType = "FAILED"
	UinStatusTypeINPROGRESS      UinStatusType = "IN_PROGRESS"
	UinStatusTypeNOTONVALIDATION UinStatusType = "NOT_ON_VALIDATION"
	UinStatusTypeOK              UinStatusType = "OK"
)

func (e UinStatusType) Valid() bool {
	switch e {
	case UinStatusTypeFAILED:
		return true
	case UinStatusTypeINPROGRESS:
		return true
	case UinStatusTypeNOTONVALIDATION:
		return true
	case UinStatusTypeOK:
		return true
	default:
		return false
	}
}

const (
	UINMERCHANTMISMATCH     UinSubstatusType = "UIN_MERCHANT_MISMATCH"
	UINMERCHANTUNREGISTERED UinSubstatusType = "UIN_MERCHANT_UNREGISTERED"
	UINNODATA               UinSubstatusType = "UIN_NO_DATA"
)

func (e UinSubstatusType) Valid() bool {
	switch e {
	case UINMERCHANTMISMATCH:
		return true
	case UINMERCHANTUNREGISTERED:
		return true
	case UINNODATA:
		return true
	default:
		return false
	}
}

func (e WarehouseComponentType) Valid() bool {
	switch e {
	case WarehouseComponentTypeADDRESS:
		return true
	case WarehouseComponentTypeSTATUS:
		return true
	default:
		return false
	}
}

func (e WarehouseStatusType) Valid() bool {
	switch e {
	case WarehouseStatusTypeDISABLEDMANUALLY:
		return true
	case WarehouseStatusTypeOTHER:
		return true
	default:
		return false
	}
}

func (e WarehouseStockType) Valid() bool {
	switch e {
	case WarehouseStockTypeAVAILABLE:
		return true
	case WarehouseStockTypeDEFECT:
		return true
	case WarehouseStockTypeEXPIRED:
		return true
	case WarehouseStockTypeFIT:
		return true
	case WarehouseStockTypeFREEZE:
		return true
	case WarehouseStockTypeQUARANTINE:
		return true
	case WarehouseStockTypeUTILIZATION:
		return true
	default:
		return false
	}
}

type AgeDTO struct {
	AgeUnit AgeUnitType `json:"ageUnit"`

	Value float32 `json:"value"`
}

type AgeUnitType string

type AnswerDTO struct {
	Author *QuestionsTextContentAuthorDTO `json:"author,omitempty"`

	CanModify bool `json:"canModify"`

	Comments *[]CommentDTO `json:"comments,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	Id int64 `json:"id"`

	QuestionId int64 `json:"questionId"`

	Status QuestionsTextContentModerationStatusType `json:"status"`

	Text string `json:"text"`

	Votes VotesDTO `json:"votes"`
}

type AnswerId = int64

type AnswerListDTO struct {
	Answers []AnswerDTO `json:"answers"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type ApiAvailabilityStatusType string

type ApiErrorDTO struct {
	Code string `json:"code"`

	Message *string `json:"message,omitempty"`
}

type ApiErrorResponse struct {
	Errors *[]ApiErrorDTO `json:"errors,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type ApiResponse struct {
	Status ApiResponseStatusType `json:"status"`
}

type ApiResponseStatusType string

type BarcodeFormatType string

type BasicCourierDeliveryAddressDTO struct {
	FullAddress string `json:"fullAddress"`
}

type BenefitType string

type CatalogLanguageType string

type ChannelType string

type CisDTO struct {
	CrptRequestDateTime *time.Time `json:"crptRequestDateTime,omitempty"`

	CrptRequestId *string `json:"crptRequestId,omitempty"`

	Status CisStatusType `json:"status"`

	Substatus *CisSubstatusType `json:"substatus,omitempty"`

	Value string `json:"value"`
}

type CisStatusType string

type CisSubstatusType string

type ClosureDocumentsContractType string

type ClosureDocumentsMonthOfYearDTO struct {
	Month int32 `json:"month"`

	Year int32 `json:"year"`
}

type CommentDTO struct {
	AnswerId int64 `json:"answerId"`

	Author *QuestionsTextContentAuthorDTO `json:"author,omitempty"`

	CanModify *bool `json:"canModify,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	Id int64 `json:"id"`

	ParentId *int64 `json:"parentId,omitempty"`

	Status QuestionsTextContentModerationStatusType `json:"status"`

	Text string `json:"text"`

	Votes *VotesDTO `json:"votes,omitempty"`
}

type CommodityCodeDTO struct {
	Code string `json:"code"`

	Type CommodityCodeType `json:"type"`
}

type CommodityCodeType string

type CountryDTO struct {
	CountryCode string `json:"countryCode"`

	Region RegionDTO `json:"region"`
}

type CourierDeliveryAddressDTO struct {
	Apartment *string `json:"apartment,omitempty"`

	Entrance *string `json:"entrance,omitempty"`

	Floor *int32 `json:"floor,omitempty"`

	FullAddress string `json:"fullAddress"`
}

type CourierDeliveryParametersDTO struct {
	FullAddress string `json:"fullAddress"`
}

type CurrencyType string

type CurrencyValueDTO struct {
	CurrencyId CurrencyType `json:"currencyId"`

	Value float32 `json:"value"`
}

type CustomerDTO struct {
	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	MiddleName *string `json:"middleName,omitempty"`

	Phone string `json:"phone"`
}

type DateDdMmYyyy = string

type DayOfWeekType string

type DeliveryDateIntervalDTO struct {
	FromDate openapi_types.Date `json:"fromDate"`

	ToDate openapi_types.Date `json:"toDate"`
}

type DeliveryIntervalsUpdateOptionDTO struct {
	DeliveryDateInterval DeliveryDateIntervalDTO `json:"deliveryDateInterval"`

	DeliveryTimeInterval TimeIntervalDTO `json:"deliveryTimeInterval"`
}

type DeliveryIntervalsUpdateOptionsDTO struct {
	AvailableDeliveryIntervals []DeliveryIntervalsUpdateOptionDTO `json:"availableDeliveryIntervals"`
}

type DeliveryPaymentType string

type DiscountBase = float32

type DocumentDTO struct {
	Date *openapi_types.Date `json:"date,omitempty"`

	Number *string `json:"number,omitempty"`

	Status *OrderDocumentStatusType `json:"status,omitempty"`
}

type EacVerificationResultDTO struct {
	AttemptsLeft *int32 `json:"attemptsLeft,omitempty"`

	VerificationResult *EacVerificationStatusType `json:"verificationResult,omitempty"`
}

type EacVerificationStatusType string

type FeedbackReactionStatusType string

type FlippingPagerDTO struct {
	CurrentPage *int32 `json:"currentPage,omitempty"`

	From *int32 `json:"from,omitempty"`

	PageSize *int32 `json:"pageSize,omitempty"`

	PagesCount *int32 `json:"pagesCount,omitempty"`

	To *int32 `json:"to,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

type GetCategoriesMaxSaleQuantumDTO struct {
	Errors *[]CategoryErrorDTO `json:"errors,omitempty"`

	Results []MaxSaleQuantumDTO `json:"results"`
}

type GetLogisticsPointsDTO struct {
	LogisticPoints []LogisticPointDTO `json:"logisticPoints"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetMappingDTO struct {
	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`

	MarketCategoryName *string `json:"marketCategoryName,omitempty"`

	MarketModelName *string `json:"marketModelName,omitempty"`

	MarketSku *int64 `json:"marketSku,omitempty"`

	MarketSkuName *string `json:"marketSkuName,omitempty"`
}

type GpsDTO struct {
	Latitude float32 `json:"latitude"`

	Longitude float32 `json:"longitude"`
}

type LabelsSortingType string

type LanguageType string

type LicenseCheckStatusType string

type LicenseType string

type LogisticPickupPointDTO struct {
	Address *PickupAddressDTO `json:"address,omitempty"`

	Id *int64 `json:"id,omitempty"`

	Instruction *string `json:"instruction,omitempty"`

	LogisticPartnerId *int64 `json:"logisticPartnerId,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *LogisticPointType `json:"type,omitempty"`
}

type MaxSaleQuantumDTO struct {
	Id int64 `json:"id"`

	MaxSaleQuantum *int `json:"maxSaleQuantum,omitempty"`

	Name *string `json:"name,omitempty"`
}

type MechanicsType string

type MediaFileUploadStateType string

type MonthOfYearDTO struct {
	Month int32 `json:"month"`

	Year int32 `json:"year"`
}

type OptionValuesLimitedDTO struct {
	LimitingOptionValueId int64 `json:"limitingOptionValueId"`

	OptionValueIds []int64 `json:"optionValueIds"`
}

type PackagingForwardScrollingPagerDTO struct {
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

type PackagingScrollingPagerDTO struct {
	NextPageToken *string `json:"nextPageToken,omitempty"`

	PrevPageToken *string `json:"prevPageToken,omitempty"`
}

type PageFormatType string

type PalletsCountDTO struct {
	Fact *int32 `json:"fact,omitempty"`

	Planned *int32 `json:"planned,omitempty"`
}

type ParameterType string

type ParameterValueConstraintsDTO struct {
	MaxLength *int32 `json:"maxLength,omitempty"`

	MaxValue *float64 `json:"maxValue,omitempty"`

	MinValue *float64 `json:"minValue,omitempty"`
}

type ParameterValueDTO struct {
	ParameterId int64 `json:"parameterId"`

	UnitId *int64 `json:"unitId,omitempty"`

	Value *string `json:"value,omitempty"`

	ValueId *int64 `json:"valueId,omitempty"`
}

type ParameterValueOptionDTO struct {
	Description *string `json:"description,omitempty"`

	Id int64 `json:"id"`

	Value string `json:"value"`
}

type ParcelBoxDTO struct {
	FulfilmentId *string `json:"fulfilmentId,omitempty"`

	Id *int64 `json:"id,omitempty"`
}

type ParcelBoxLabelDTO struct {
	BoxId int64 `json:"boxId"`

	DeliveryAddress *string `json:"deliveryAddress,omitempty"`

	DeliveryServiceId string `json:"deliveryServiceId"`

	DeliveryServiceName string `json:"deliveryServiceName"`

	FulfilmentId string `json:"fulfilmentId"`

	OrderId int64 `json:"orderId"`

	OrderNum string `json:"orderNum"`

	Place string `json:"place"`

	RecipientName string `json:"recipientName"`

	ShipmentDate *string `json:"shipmentDate,omitempty"`

	SupplierName string `json:"supplierName"`
	Url          string `json:"url"`

	Weight string `json:"weight"`
}

type ParcelBoxRequestDTO struct {
	FulfilmentId *string `json:"fulfilmentId,omitempty"`
}

type ParcelRequestDTO struct {
	Boxes []ParcelBoxRequestDTO `json:"boxes"`
}

type PartialCompensationBoundsDTO struct {
	MaxAmount BasePriceDTO `json:"maxAmount"`

	MaxPercent int64 `json:"maxPercent"`

	MinAmount BasePriceDTO `json:"minAmount"`
}

type PaymentFrequencyType string

type PickupAddressDTO struct {
	City *string `json:"city,omitempty"`

	Country *string `json:"country,omitempty"`

	House *string `json:"house,omitempty"`

	Postcode *string `json:"postcode,omitempty"`

	Street *string `json:"street,omitempty"`
}

type PickupDeliveryParametersDTO struct {
	LogisticPointsIds []LogisticPointId `json:"logisticPointsIds"`
}

type PickupOptionDTO struct {
	DeliveryDateInterval DeliveryDateIntervalDTO `json:"deliveryDateInterval"`

	Price DeliveryOptionPriceDTO `json:"price"`
}

type PickupOptionsDTO struct {
	LogisticPointId int64 `json:"logisticPointId"`

	Options []PickupOptionDTO `json:"options"`
}

type PlacementType string

type QuantumDTO struct {
	MinQuantity *int32 `json:"minQuantity,omitempty"`

	StepQuantity *int32 `json:"stepQuantity,omitempty"`
}

type QuestionDTO struct {
	Author QuestionsTextContentAuthorDTO `json:"author"`

	BusinessId int64 `json:"businessId"`

	CreatedAt time.Time `json:"createdAt"`

	QuestionIdentifiers QuestionIdentifiersDTO `json:"questionIdentifiers"`

	Text string `json:"text"`

	Votes VotesDTO `json:"votes"`
}

type QuestionId = int64

type QuestionIdentifiersDTO struct {
	CategoryId *int64 `json:"categoryId,omitempty"`

	Id int64 `json:"id"`

	OfferId string `json:"offerId"`
}

type QuestionListDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Questions []QuestionDTO `json:"questions"`

	TotalCount int64 `json:"totalCount"`
}

type QuestionsTextEntityType string

type RecipientType string

type RefundStatusType string

type ScheduleDayDTO struct {
	Day DayOfWeekType `json:"day"`

	EndTime string `json:"endTime"`

	StartTime string `json:"startTime"`
}

type SellingProgramType string

type ShopSku = string

type ShowcaseType string

type ShowcaseUrlDTO struct {
	ShowcaseType ShowcaseType `json:"showcaseType"`

	ShowcaseUrl string `json:"showcaseUrl"`
}

type ShowsSalesGroupingType string

type SignatureDTO struct {
	Signed bool `json:"signed"`
}

type StatisticsAttributionType string

type TimeIntervalDTO struct {
	FromTime string `json:"fromTime"`

	ToTime string `json:"toTime"`
}

type TimePeriodDTO struct {
	Comment *string `json:"comment,omitempty"`

	TimePeriod int `json:"timePeriod"`

	TimeUnit TimeUnitType `json:"timeUnit"`
}

type TimeUnitType string

type TrackDTO struct {
	TrackCode *string `json:"trackCode,omitempty"`
}

type TurnoverDTO struct {
	Turnover TurnoverType `json:"turnover"`

	TurnoverDays *float64 `json:"turnoverDays,omitempty"`
}

type TurnoverType string

type TypedQuestionsTextEntityIdDTO struct {
	Id int64 `json:"id"`

	Type QuestionsTextEntityType `json:"type"`
}

type UinDTO struct {
	Status UinStatusType `json:"status"`

	Substatus *UinSubstatusType `json:"substatus,omitempty"`

	Value string `json:"value"`
}

type UinStatusType string

type UinSubstatusType string

type UnitDTO struct {
	FullName string `json:"fullName"`

	Id int64 `json:"id"`

	Name string `json:"name"`
}

type UpdateMappingDTO struct {
	MarketSku *int64 `json:"marketSku,omitempty"`
}

type UpdateTimeDTO struct {
	UpdatedAt time.Time `json:"updatedAt"`
}

type Url = string

type ValueRestrictionDTO struct {
	LimitedValues []OptionValuesLimitedDTO `json:"limitedValues"`

	LimitingParameterId int64 `json:"limitingParameterId"`
}

type VatType = int32

type VotesDTO struct {
	Dislikes int64 `json:"dislikes"`

	Likes int64 `json:"likes"`
}

type apiKeyContextKey string

type oAuthContextKey string

type GetCategoriesMaxSaleQuantumJSONRequestBody = GetCategoriesMaxSaleQuantumRequest

type GetCategoriesTreeJSONRequestBody = GetCategoriesRequest
