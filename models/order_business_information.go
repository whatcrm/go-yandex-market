package models

type GetBusinessBuyerInfoResponse struct {
	Result *OrderBusinessBuyerDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetBusinessDocumentsInfoResponse struct {
	Result *OrderBusinessDocumentsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
