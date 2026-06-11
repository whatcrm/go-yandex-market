package returns

const (
	CancelReturnEndpoint                = "/v1/campaigns/{campaignId}/returns/cancel"
	CreateReturnEndpoint                = "/v1/campaigns/{campaignId}/returns/create"
	GetReturnEndpoint                   = "/v2/campaigns/{campaignId}/orders/{orderId}/returns/{returnId}"
	GetReturnApplicationEndpoint        = "/v2/campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/application"
	GetReturnAvailableDecisionsEndpoint = "/v1/businesses/{businessId}/returns/decisions"
	GetReturnPhotoEndpoint              = "/v2/campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/{itemId}/image/{imageHash}"
	GetReturnsEndpoint                  = "/v2/campaigns/{campaignId}/returns"
	SetReturnDecisionEndpoint           = "/v2/campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision"
	SubmitReturnDecisionEndpoint        = "/v2/campaigns/{campaignId}/orders/{orderId}/returns/{returnId}/decision/submit"
)
