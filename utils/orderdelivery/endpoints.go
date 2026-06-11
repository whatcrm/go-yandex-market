package orderdelivery

const (
	GetOrderBuyerInfoEndpoint         = "/v2/campaigns/{campaignId}/orders/{orderId}/buyer"
	SetOrderDeliveryDateEndpoint      = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/date"
	SetOrderDeliveryTrackCodeEndpoint = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/track"
	UpdateOrderStorageLimitEndpoint   = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/storage-limit"
	VerifyOrderEacEndpoint            = "/v2/campaigns/{campaignId}/orders/{orderId}/verifyEac"
)
