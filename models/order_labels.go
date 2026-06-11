package models

type GetOrderLabelsDataResponse struct {
	Result *OrderLabelDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
