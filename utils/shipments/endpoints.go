package shipments

const (
	ConfirmShipmentEndpoint                       = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/confirm"
	DownloadShipmentActEndpoint                   = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/act"
	DownloadShipmentDiscrepancyActEndpoint        = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/discrepancy-act"
	DownloadShipmentInboundActEndpoint            = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/inbound-act"
	DownloadShipmentPalletLabelsEndpoint          = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/pallet/labels"
	DownloadShipmentReceptionTransferActEndpoint  = "/v2/campaigns/{campaignId}/shipments/reception-transfer-act"
	DownloadShipmentTransportationWaybillEndpoint = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/transportation-waybill"
	GetShipmentEndpoint                           = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}"
	GetShipmentOrdersInfoEndpoint                 = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/orders/info"
	SearchShipmentsEndpoint                       = "/v2/campaigns/{campaignId}/first-mile/shipments"
	SetShipmentPalletsCountEndpoint               = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/pallets"
	TransferOrdersFromShipmentEndpoint            = "/v2/campaigns/{campaignId}/first-mile/shipments/{shipmentId}/orders/transfer"
)
