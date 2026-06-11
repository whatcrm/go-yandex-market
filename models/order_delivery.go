package models

import openapi_types "github.com/oapi-codegen/runtime/types"

type GetOrderBuyerInfoResponse struct {
	Result *OrderBuyerInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SetOrderDeliveryDateRequest struct {
	Dates OrderDeliveryDateDTO `json:"dates"`

	Reason OrderDeliveryDateReasonType `json:"reason"`
}

type SetOrderDeliveryTrackCodeRequest struct {
	DeliveryServiceId int64 `json:"deliveryServiceId"`

	TrackCode string `json:"trackCode"`
}

type UpdateOrderStorageLimitRequest struct {
	NewDate openapi_types.Date `json:"newDate"`
}

type VerifyOrderEacRequest struct {
	Code string `json:"code"`
}

type VerifyOrderEacResponse struct {
	Result *EacVerificationResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
