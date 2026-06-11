package models

import openapi_types "github.com/oapi-codegen/runtime/types"

const (
	QualityRatingComponentTypeDBSCANCELLATIONRATE QualityRatingComponentType = "DBS_CANCELLATION_RATE"
	QualityRatingComponentTypeDBSLATEDELIVERYRATE QualityRatingComponentType = "DBS_LATE_DELIVERY_RATE"
	QualityRatingComponentTypeFBSCANCELLATIONRATE QualityRatingComponentType = "FBS_CANCELLATION_RATE"
	QualityRatingComponentTypeFBSLATESHIPRATE     QualityRatingComponentType = "FBS_LATE_SHIP_RATE"
	QualityRatingComponentTypeFBYCANCELLATIONRATE QualityRatingComponentType = "FBY_CANCELLATION_RATE"
	QualityRatingComponentTypeFBYDELIVERYDIFFRATE QualityRatingComponentType = "FBY_DELIVERY_DIFF_RATE"
	QualityRatingComponentTypeFBYLATEDELIVERYRATE QualityRatingComponentType = "FBY_LATE_DELIVERY_RATE"
	QualityRatingComponentTypeFBYLATEEDITINGRATE  QualityRatingComponentType = "FBY_LATE_EDITING_RATE"
)

type GetQualityRatingDetailsResponse struct {
	Result *QualityRatingDetailsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetQualityRatingRequest struct {
	CampaignIds []CampaignId `json:"campaignIds"`

	DateFrom *openapi_types.Date `json:"dateFrom,omitempty"`

	DateTo *openapi_types.Date `json:"dateTo,omitempty"`
}

type GetQualityRatingResponse struct {
	Result *CampaignsQualityRatingDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type QualityRatingComponentDTO struct {
	ComponentType QualityRatingComponentType `json:"componentType"`

	Value float64 `json:"value"`
}

type QualityRatingComponentType string

type QualityRatingDTO struct {
	CalculationDate openapi_types.Date `json:"calculationDate"`

	Components []QualityRatingComponentDTO `json:"components"`

	Rating int64 `json:"rating"`
}

type QualityRatingDetailsDTO struct {
	AffectedOrders []QualityRatingAffectedOrderDTO `json:"affectedOrders"`
}

type GetQualityRatingsJSONRequestBody = GetQualityRatingRequest
