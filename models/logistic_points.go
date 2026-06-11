package models

import openapi_types "github.com/oapi-codegen/runtime/types"

const (
	LogisticPointBrandTypeMARKET LogisticPointBrandType = "MARKET"
)

const (
	LogisticPointTypePICKUPMIXED      LogisticPointType = "PICKUP_MIXED"
	LogisticPointTypePICKUPPOINT      LogisticPointType = "PICKUP_POINT"
	LogisticPointTypePICKUPPOSTOFFICE LogisticPointType = "PICKUP_POST_OFFICE"
	LogisticPointTypePICKUPRETAIL     LogisticPointType = "PICKUP_RETAIL"
	LogisticPointTypePICKUPTERMINAL   LogisticPointType = "PICKUP_TERMINAL"
	LogisticPointTypeWAREHOUSE        LogisticPointType = "WAREHOUSE"
)

type GetLogisticPointsResponse struct {
	Result *GetLogisticsPointsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type LogisticPointAddressDTO struct {
	Additional *string `json:"additional,omitempty"`

	Block *string `json:"block,omitempty"`

	Building *string `json:"building,omitempty"`

	City *string `json:"city,omitempty"`

	FullAddress string `json:"fullAddress"`

	Gps GpsDTO `json:"gps"`

	House *string `json:"house,omitempty"`

	Km *int32 `json:"km,omitempty"`

	RegionId int64 `json:"regionId"`

	Street *string `json:"street,omitempty"`
}

type LogisticPointBrandType string

type LogisticPointDTO struct {
	Address LogisticPointAddressDTO `json:"address"`

	Brand LogisticPointBrandType `json:"brand"`

	DeliveryRestrictions LogisticPointDeliveryRestrictionDTO `json:"deliveryRestrictions"`

	Features *[]LogisticPointFeatureType `json:"features,omitempty"`

	LogisticPointId int64 `json:"logisticPointId"`

	StoragePeriod int64 `json:"storagePeriod"`

	WorkingSchedule LogisticPointScheduleDTO `json:"workingSchedule"`
}

type LogisticPointDeliveryRestrictionDTO struct {
	DimensionsRestrictions LogisticPointDimensionRestrictionsDTO `json:"dimensionsRestrictions"`
}

type LogisticPointDimensionRestrictionsDTO struct {
	DimensionsSum int32 `json:"dimensionsSum"`

	Height int32 `json:"height"`

	Length int32 `json:"length"`

	Weight int32 `json:"weight"`

	Width int32 `json:"width"`
}

type LogisticPointFeatureType string

type LogisticPointId = int64

type LogisticPointScheduleDTO struct {
	Holidays *[]openapi_types.Date `json:"holidays,omitempty"`

	Schedule []ScheduleDayDTO `json:"schedule"`
}

type LogisticPointType string

type GetLogisticPointsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}
