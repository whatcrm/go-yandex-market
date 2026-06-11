package orders

const (
	AcceptOrderCancellationEndpoint     = "/v2/campaigns/{campaignId}/orders/{orderId}/cancellation/accept"
	CreateOrderEndpoint                 = "/v1/campaigns/{campaignId}/orders/create"
	GetBusinessOrdersEndpoint           = "/v1/businesses/{businessId}/orders"
	GetOrderEndpoint                    = "/v2/campaigns/{campaignId}/orders/{orderId}"
	GetOrderIdentifiersStatusEndpoint   = "/v2/campaigns/{campaignId}/orders/{orderId}/identifiers/status"
	GetOrderUpdateOptionsEndpoint       = "/v1/campaigns/{campaignId}/orders/update-options"
	GetOrdersEndpoint                   = "/v2/campaigns/{campaignId}/orders"
	ProvideOrderDigitalCodesEndpoint    = "/v2/campaigns/{campaignId}/orders/{orderId}/deliverDigitalGoods"
	ProvideOrderItemIdentifiersEndpoint = "/v2/campaigns/{campaignId}/orders/{orderId}/identifiers"
	SetOrderBoxLayoutEndpoint           = "/v2/campaigns/{campaignId}/orders/{orderId}/boxes"
	SetOrderShipmentBoxesEndpoint       = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes"
	UpdateExternalOrderIdEndpoint       = "/v2/campaigns/{campaignId}/orders/{orderId}/external-id"
	UpdateOrderEndpoint                 = "/v1/campaigns/{campaignId}/orders/update"
	UpdateOrderItemsEndpoint            = "/v2/campaigns/{campaignId}/orders/{orderId}/items"
	UpdateOrderStatusEndpoint           = "/v2/campaigns/{campaignId}/orders/{orderId}/status"
	UpdateOrderStatusesEndpoint         = "/v2/campaigns/{campaignId}/orders/status-update"
)
