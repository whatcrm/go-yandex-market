package orderlabels

const (
	GenerateOrderLabelEndpoint  = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes/{boxId}/label"
	GenerateOrderLabelsEndpoint = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/labels"
	GetOrderLabelsDataEndpoint  = "/v2/campaigns/{campaignId}/orders/{orderId}/delivery/labels/data"
)
