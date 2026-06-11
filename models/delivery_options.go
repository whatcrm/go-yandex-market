package models

type CourierDeliveryOptionDTO struct {
	DeliveryDateInterval DeliveryDateIntervalDTO `json:"deliveryDateInterval"`

	DeliveryTimeInterval TimeIntervalDTO `json:"deliveryTimeInterval"`

	Price DeliveryOptionPriceDTO `json:"price"`
}

type CourierDeliveryOptionsDTO struct {
	CourierDeliveryOptions []CourierDeliveryOptionDTO `json:"courierDeliveryOptions"`
}

type GetDeliveryOptionsDTO struct {
	WarehousesDeliveryOptions []WarehousesDeliveryOptionsDTO `json:"warehousesDeliveryOptions"`
}

type GetDeliveryOptionsItemDTO struct {
	Count int32 `json:"count"`

	OfferId string `json:"offerId"`

	WarehouseId *int64 `json:"warehouseId,omitempty"`
}

type GetDeliveryOptionsRequest struct {
	CourierDelivery *CourierDeliveryParametersDTO `json:"courierDelivery,omitempty"`

	Items []GetDeliveryOptionsItemDTO `json:"items"`

	PickupDelivery *PickupDeliveryParametersDTO `json:"pickupDelivery,omitempty"`
}

type GetDeliveryOptionsResponse struct {
	Result *GetDeliveryOptionsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetReturnDeliveryOptionsRequest struct {
	Items []BasicOrderItemDTO `json:"items"`

	PickupDelivery PickupDeliveryParametersDTO `json:"pickupDelivery"`
}

type GetReturnDeliveryOptionsResponse struct {
	Result *GetReturnDeliveryOptionsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type PickupDeliveryOptionsDTO struct {
	PickupOptions []PickupOptionsDTO `json:"pickupOptions"`
}

type GetDeliveryOptionsJSONRequestBody = GetDeliveryOptionsRequest
