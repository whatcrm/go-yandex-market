package models

type DeliveryServiceDTO struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type DeliveryServiceInfoDTO struct {
	Id int64 `json:"id"`

	Name string `json:"name"`
}

type DeliveryServicesDTO struct {
	DeliveryService []DeliveryServiceInfoDTO `json:"deliveryService"`
}

type GetDeliveryServicesResponse struct {
	Result *DeliveryServicesDTO `json:"result,omitempty"`
}
