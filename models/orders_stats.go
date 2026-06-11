package models

import openapi_types "github.com/oapi-codegen/runtime/types"

type GetOrdersStatsRequest struct {
	DateFrom *openapi_types.Date `json:"dateFrom,omitempty"`

	DateTo *openapi_types.Date `json:"dateTo,omitempty"`

	HasCis *bool `json:"hasCis,omitempty"`

	Orders *[]int64 `json:"orders,omitempty"`

	Statuses *[]OrderStatsStatusType `json:"statuses,omitempty"`

	UpdateFrom *openapi_types.Date `json:"updateFrom,omitempty"`

	UpdateTo *openapi_types.Date `json:"updateTo,omitempty"`
}

type GetOrdersStatsResponse struct {
	Result *OrdersStatsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
