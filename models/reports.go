package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	KeyIndicatorsReportDetalizationLevelTypeMONTH KeyIndicatorsReportDetalizationLevelType = "MONTH"
	KeyIndicatorsReportDetalizationLevelTypeWEEK  KeyIndicatorsReportDetalizationLevelType = "WEEK"
)

const (
	ReportStatusTypeDONE       ReportStatusType = "DONE"
	ReportStatusTypeFAILED     ReportStatusType = "FAILED"
	ReportStatusTypePENDING    ReportStatusType = "PENDING"
	ReportStatusTypePROCESSING ReportStatusType = "PROCESSING"
)

type GenerateBannersStatisticsRequest struct {
	BusinessId int64 `json:"businessId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateBarcodesReportRequest struct {
	BarcodeFormat BarcodeFormatType `json:"barcodeFormat"`

	BarcodeOfferInfos *[]BarcodeOfferInfoDTO `json:"barcodeOfferInfos,omitempty"`

	CampaignId int64 `json:"campaignId"`

	SupplyRequestId *int64 `json:"supplyRequestId,omitempty"`
}

type GenerateBoostConsolidatedRequest struct {
	BusinessId int64 `json:"businessId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateClosureDocumentsDetalizationRequest struct {
	CampaignId int64 `json:"campaignId"`

	ContractType ClosureDocumentsContractType `json:"contractType"`

	MonthOfYear ClosureDocumentsMonthOfYearDTO `json:"monthOfYear"`
}

type GenerateClosureDocumentsRequest struct {
	CampaignId int64 `json:"campaignId"`

	ContractTypes *[]ClosureDocumentsContractType `json:"contractTypes,omitempty"`

	MonthOfYear ClosureDocumentsMonthOfYearDTO `json:"monthOfYear"`
}

type GenerateCompetitorsPositionReportRequest struct {
	BusinessId int64 `json:"businessId"`

	CategoryId int64 `json:"categoryId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateGoodsFeedbackRequest struct {
	BusinessId int64 `json:"businessId"`
}

type GenerateGoodsMovementReportRequest struct {
	CampaignId int64 `json:"campaignId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	ShopSku *string `json:"shopSku,omitempty"`
}

type GenerateGoodsPricesReportRequest struct {
	BusinessId *int64 `json:"businessId,omitempty"`

	CampaignId *int64 `json:"campaignId,omitempty"`

	CategoryIds *[]int32 `json:"categoryIds,omitempty"`
}

type GenerateGoodsRealizationReportRequest struct {
	CampaignId int64 `json:"campaignId"`

	Month int32 `json:"month"`

	Year int32 `json:"year"`
}

type GenerateGoodsTurnoverRequest struct {
	CampaignId int64 `json:"campaignId"`

	Date *openapi_types.Date `json:"date,omitempty"`
}

type GenerateJewelryFiscalReportRequest struct {
	CampaignId int64 `json:"campaignId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateKeyIndicatorsRequest struct {
	BusinessId *int64 `json:"businessId,omitempty"`

	CampaignId *int64 `json:"campaignId,omitempty"`

	DetalizationLevel KeyIndicatorsReportDetalizationLevelType `json:"detalizationLevel"`
}

type GenerateMassOrderLabelsRequest struct {
	BusinessId int64 `json:"businessId"`

	OrderIds []int64 `json:"orderIds"`

	SortingType *LabelsSortingType `json:"sortingType,omitempty"`
}

type GenerateReportDTO struct {
	EstimatedGenerationTime int64 `json:"estimatedGenerationTime"`

	ReportId string `json:"reportId"`
}

type GenerateReportResponse struct {
	Result *GenerateReportDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GenerateSalesGeographyRequest struct {
	BusinessId int64 `json:"businessId"`

	CategoryIds *[]int32 `json:"categoryIds,omitempty"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`
}

type GenerateShelfsStatisticsRequest struct {
	AttributionType StatisticsAttributionType `json:"attributionType"`

	BusinessId int64 `json:"businessId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateShipmentListDocumentReportRequest struct {
	CampaignId int64 `json:"campaignId"`

	OrderIds *[]int64 `json:"orderIds,omitempty"`

	ShipmentId *int64 `json:"shipmentId,omitempty"`
}

type GenerateShowsBoostRequest struct {
	AttributionType StatisticsAttributionType `json:"attributionType"`

	BusinessId int64 `json:"businessId"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`
}

type GenerateShowsSalesReportRequest struct {
	BusinessId *int64 `json:"businessId,omitempty"`

	CampaignId *int64 `json:"campaignId,omitempty"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	Grouping ShowsSalesGroupingType `json:"grouping"`
}

type GenerateStocksOnWarehousesReportRequest struct {
	BusinessId *int64 `json:"businessId,omitempty"`

	CampaignId *int64 `json:"campaignId,omitempty"`

	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	CategoryIds *[]int32 `json:"categoryIds,omitempty"`

	HasStocks *bool `json:"hasStocks,omitempty"`

	ReportDate *openapi_types.Date `json:"reportDate,omitempty"`

	WarehouseIds *[]int64 `json:"warehouseIds,omitempty"`
}

type GenerateUnitedMarketplaceServicesReportRequest struct {
	BusinessId int64 `json:"businessId"`

	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	DateFrom *openapi_types.Date `json:"dateFrom,omitempty"`

	DateTimeFrom *time.Time `json:"dateTimeFrom,omitempty"`

	DateTimeTo *time.Time `json:"dateTimeTo,omitempty"`

	DateTo *openapi_types.Date `json:"dateTo,omitempty"`

	Inns *[]string `json:"inns,omitempty"`

	MonthFrom *int32 `json:"monthFrom,omitempty"`

	MonthTo *int32 `json:"monthTo,omitempty"`

	PlacementPrograms *[]PlacementType `json:"placementPrograms,omitempty"`

	YearFrom *int32 `json:"yearFrom,omitempty"`

	YearTo *int32 `json:"yearTo,omitempty"`
}

type GenerateUnitedNettingReportRequest struct {
	BankOrderDateTime *time.Time `json:"bankOrderDateTime,omitempty"`

	BankOrderId *int64 `json:"bankOrderId,omitempty"`

	BusinessId int64 `json:"businessId"`

	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	DateFrom *openapi_types.Date `json:"dateFrom,omitempty"`

	DateTimeFrom *time.Time `json:"dateTimeFrom,omitempty"`

	DateTimeTo *time.Time `json:"dateTimeTo,omitempty"`

	DateTo *openapi_types.Date `json:"dateTo,omitempty"`

	Inns *[]string `json:"inns,omitempty"`

	MonthOfYear *MonthOfYearDTO `json:"monthOfYear,omitempty"`

	PlacementPrograms *[]PlacementType `json:"placementPrograms,omitempty"`
}

type GenerateUnitedOrdersRequest struct {
	BusinessId int64 `json:"businessId"`

	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	PromoId *string `json:"promoId,omitempty"`
}

type GenerateUnitedReturnsRequest struct {
	BusinessId int64 `json:"businessId"`

	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	ReturnStatusTypes *[]ReturnShipmentStatusType `json:"returnStatusTypes,omitempty"`

	ReturnType *ReturnType `json:"returnType,omitempty"`
}

type GetReportInfoResponse struct {
	Result *ReportInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type KeyIndicatorsReportDetalizationLevelType string

type ReportFormatType string

type ReportInfoDTO struct {
	EstimatedGenerationTime *int64 `json:"estimatedGenerationTime,omitempty"`

	File *string `json:"file,omitempty"`

	GenerationFinishedAt *time.Time `json:"generationFinishedAt,omitempty"`

	GenerationRequestedAt time.Time `json:"generationRequestedAt"`

	Status ReportStatusType `json:"status"`

	SubStatus *ReportSubStatusType `json:"subStatus,omitempty"`
}

type ReportLanguageType string

type ReportStatusType string

type ReportSubStatusType string

type GenerateBannersStatisticsReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateBoostConsolidatedReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateClosureDocumentsDetalizationReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateCompetitorsPositionReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateGoodsFeedbackReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateGoodsMovementReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateGoodsPricesReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateGoodsRealizationReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateGoodsTurnoverReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateJewelryFiscalReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateKeyIndicatorsReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateSalesGeographyReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateShelfsStatisticsReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateShowsBoostReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateShowsSalesReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateStocksOnWarehousesReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateUnitedMarketplaceServicesReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`

	Language *ReportLanguageType `form:"language,omitempty" json:"language,omitempty"`
}

type GenerateUnitedNettingReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`

	Language *ReportLanguageType `form:"language,omitempty" json:"language,omitempty"`
}

type GenerateBarcodesReportJSONRequestBody = GenerateBarcodesReportRequest

type GenerateBannersStatisticsReportJSONRequestBody = GenerateBannersStatisticsRequest

type GenerateBoostConsolidatedReportJSONRequestBody = GenerateBoostConsolidatedRequest

type GenerateClosureDocumentsDetalizationReportJSONRequestBody = GenerateClosureDocumentsDetalizationRequest

type GenerateClosureDocumentsReportJSONRequestBody = GenerateClosureDocumentsRequest

type GenerateCompetitorsPositionReportJSONRequestBody = GenerateCompetitorsPositionReportRequest

type GenerateGoodsFeedbackReportJSONRequestBody = GenerateGoodsFeedbackRequest

type GenerateGoodsMovementReportJSONRequestBody = GenerateGoodsMovementReportRequest

type GenerateGoodsPricesReportJSONRequestBody = GenerateGoodsPricesReportRequest

type GenerateGoodsRealizationReportJSONRequestBody = GenerateGoodsRealizationReportRequest

type GenerateGoodsTurnoverReportJSONRequestBody = GenerateGoodsTurnoverRequest

type GenerateJewelryFiscalReportJSONRequestBody = GenerateJewelryFiscalReportRequest

type GenerateKeyIndicatorsReportJSONRequestBody = GenerateKeyIndicatorsRequest

type GenerateSalesGeographyReportJSONRequestBody = GenerateSalesGeographyRequest

type GenerateShelfsStatisticsReportJSONRequestBody = GenerateShelfsStatisticsRequest

type GenerateShowsBoostReportJSONRequestBody = GenerateShowsBoostRequest

type GenerateShowsSalesReportJSONRequestBody = GenerateShowsSalesReportRequest

type GenerateStocksOnWarehousesReportJSONRequestBody = GenerateStocksOnWarehousesReportRequest

type GenerateUnitedMarketplaceServicesReportJSONRequestBody = GenerateUnitedMarketplaceServicesReportRequest

type GenerateUnitedNettingReportJSONRequestBody = GenerateUnitedNettingReportRequest
