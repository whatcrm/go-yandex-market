package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	ReturnShipmentStatusTypeCANCELLED                ReturnShipmentStatusType = "CANCELLED"
	ReturnShipmentStatusTypeCREATED                  ReturnShipmentStatusType = "CREATED"
	ReturnShipmentStatusTypeEXPIRED                  ReturnShipmentStatusType = "EXPIRED"
	ReturnShipmentStatusTypeFULFILMENTRECEIVED       ReturnShipmentStatusType = "FULFILMENT_RECEIVED"
	ReturnShipmentStatusTypeINTRANSIT                ReturnShipmentStatusType = "IN_TRANSIT"
	ReturnShipmentStatusTypeLOST                     ReturnShipmentStatusType = "LOST"
	ReturnShipmentStatusTypeNOTINDEMAND              ReturnShipmentStatusType = "NOT_IN_DEMAND"
	ReturnShipmentStatusTypePICKED                   ReturnShipmentStatusType = "PICKED"
	ReturnShipmentStatusTypePREPAREDFORUTILIZATION   ReturnShipmentStatusType = "PREPARED_FOR_UTILIZATION"
	ReturnShipmentStatusTypeREADYFOREXPROPRIATION    ReturnShipmentStatusType = "READY_FOR_EXPROPRIATION"
	ReturnShipmentStatusTypeREADYFORPICKUP           ReturnShipmentStatusType = "READY_FOR_PICKUP"
	ReturnShipmentStatusTypeRECEIVED                 ReturnShipmentStatusType = "RECEIVED"
	ReturnShipmentStatusTypeRECEIVEDFOREXPROPRIATION ReturnShipmentStatusType = "RECEIVED_FOR_EXPROPRIATION"
	ReturnShipmentStatusTypeUNKNOWN                  ReturnShipmentStatusType = "UNKNOWN"
	ReturnShipmentStatusTypeUTILIZED                 ReturnShipmentStatusType = "UTILIZED"
)

const (
	ShipmentPalletLabelPageFormatTypeA4 ShipmentPalletLabelPageFormatType = "A4"
	ShipmentPalletLabelPageFormatTypeA8 ShipmentPalletLabelPageFormatType = "A8"
)

const (
	ShipmentStatusTypeACCEPTED                     ShipmentStatusType = "ACCEPTED"
	ShipmentStatusTypeACCEPTEDWITHDISCREPANCIES    ShipmentStatusType = "ACCEPTED_WITH_DISCREPANCIES"
	ShipmentStatusTypeERROR                        ShipmentStatusType = "ERROR"
	ShipmentStatusTypeFINISHED                     ShipmentStatusType = "FINISHED"
	ShipmentStatusTypeOUTBOUNDCONFIRMED            ShipmentStatusType = "OUTBOUND_CONFIRMED"
	ShipmentStatusTypeOUTBOUNDCREATED              ShipmentStatusType = "OUTBOUND_CREATED"
	ShipmentStatusTypeOUTBOUNDREADYFORCONFIRMATION ShipmentStatusType = "OUTBOUND_READY_FOR_CONFIRMATION"
	ShipmentStatusTypeOUTBOUNDSIGNED               ShipmentStatusType = "OUTBOUND_SIGNED"
)

const (
	ShipmentTypeIMPORT   ShipmentType = "IMPORT"
	ShipmentTypeWITHDRAW ShipmentType = "WITHDRAW"
)

type BaseShipmentDTO struct {
	DeliveryService *DeliveryServiceDTO `json:"deliveryService,omitempty"`

	DraftCount int32 `json:"draftCount"`

	ExternalId *string `json:"externalId,omitempty"`

	FactCount int32 `json:"factCount"`

	Id int64 `json:"id"`

	OrderIds []int64 `json:"orderIds"`

	PalletsCount *PalletsCountDTO `json:"palletsCount,omitempty"`

	PlanIntervalFrom time.Time `json:"planIntervalFrom"`

	PlanIntervalTo time.Time `json:"planIntervalTo"`

	PlannedCount int32 `json:"plannedCount"`

	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`

	Signature SignatureDTO `json:"signature"`

	Warehouse *PartnerShipmentWarehouseDTO `json:"warehouse,omitempty"`

	WarehouseTo *PartnerShipmentWarehouseDTO `json:"warehouseTo,omitempty"`
}

type ConfirmShipmentRequest struct {
	ExternalShipmentId *string `json:"externalShipmentId,omitempty"`

	Signatory *string `json:"signatory,omitempty"`
}

type ExtensionShipmentDTO struct {
	AvailableActions []ShipmentActionType `json:"availableActions"`

	CurrentStatus *ShipmentStatusChangeDTO `json:"currentStatus,omitempty"`
}

type GetShipmentOrdersInfoResponse struct {
	Result *OrdersShipmentInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetShipmentResponse struct {
	Result *ShipmentDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type PartnerShipmentWarehouseDTO struct {
	Address *string `json:"address,omitempty"`

	Id int64 `json:"id"`

	Name *string `json:"name,omitempty"`
}

type ReturnShipmentStatusType string

type SearchShipmentsRequest struct {
	CancelledOrders *bool `json:"cancelledOrders,omitempty"`

	DateFrom openapi_types.Date `json:"dateFrom"`

	DateTo openapi_types.Date `json:"dateTo"`

	OrderIds *[]int64 `json:"orderIds,omitempty"`

	Statuses *[]ShipmentStatusType `json:"statuses,omitempty"`
}

type SearchShipmentsResponse struct {
	Result *SearchShipmentsResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SearchShipmentsResponseDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Shipments []ShipmentInfoDTO `json:"shipments"`
}

type SetShipmentPalletsCountRequest struct {
	PlacesCount int32 `json:"placesCount"`
}

type ShipmentActionType string

type ShipmentBoxesDTO struct {
	Boxes []ParcelBoxDTO `json:"boxes"`
}

type ShipmentDTO struct {
	AvailableActions []ShipmentActionType `json:"availableActions"`

	CurrentStatus *ShipmentStatusChangeDTO `json:"currentStatus,omitempty"`

	DeliveryService *DeliveryServiceDTO `json:"deliveryService,omitempty"`

	DraftCount int32 `json:"draftCount"`

	ExternalId *string `json:"externalId,omitempty"`

	FactCount int32 `json:"factCount"`

	Id int64 `json:"id"`

	OrderIds []int64 `json:"orderIds"`

	PalletsCount *PalletsCountDTO `json:"palletsCount,omitempty"`

	PlanIntervalFrom time.Time `json:"planIntervalFrom"`

	PlanIntervalTo time.Time `json:"planIntervalTo"`

	PlannedCount int32 `json:"plannedCount"`

	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`

	Signature SignatureDTO `json:"signature"`

	Warehouse *PartnerShipmentWarehouseDTO `json:"warehouse,omitempty"`

	WarehouseTo *PartnerShipmentWarehouseDTO `json:"warehouseTo,omitempty"`
}

type ShipmentInfoDTO struct {
	DeliveryService *DeliveryServiceDTO `json:"deliveryService,omitempty"`

	DraftCount int32 `json:"draftCount"`

	ExternalId *string `json:"externalId,omitempty"`

	FactCount int32 `json:"factCount"`

	Id int64 `json:"id"`

	OrderIds []int64 `json:"orderIds"`

	PalletsCount *PalletsCountDTO `json:"palletsCount,omitempty"`

	PlanIntervalFrom time.Time `json:"planIntervalFrom"`

	PlanIntervalTo time.Time `json:"planIntervalTo"`

	PlannedCount int32 `json:"plannedCount"`

	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`

	Signature SignatureDTO `json:"signature"`

	Status *ShipmentStatusType `json:"status,omitempty"`

	StatusDescription *string `json:"statusDescription,omitempty"`

	StatusUpdateTime *time.Time `json:"statusUpdateTime,omitempty"`

	Warehouse *PartnerShipmentWarehouseDTO `json:"warehouse,omitempty"`

	WarehouseTo *PartnerShipmentWarehouseDTO `json:"warehouseTo,omitempty"`
}

type ShipmentPalletLabelPageFormatType string

type ShipmentStatusChangeDTO struct {
	Description *string `json:"description,omitempty"`

	Status *ShipmentStatusType `json:"status,omitempty"`

	UpdateTime *time.Time `json:"updateTime,omitempty"`
}

type ShipmentStatusType string

type ShipmentType string

type TransferOrdersFromShipmentRequest struct {
	OrderIds []int64 `json:"orderIds"`
}

type SearchShipmentsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetShipmentParams struct {
	CancelledOrders *bool `form:"cancelledOrders,omitempty" json:"cancelledOrders,omitempty"`
}

type DownloadShipmentPalletLabelsParams struct {
	Format *ShipmentPalletLabelPageFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type DownloadShipmentReceptionTransferActParams struct {
	WarehouseId *int32 `form:"warehouse_id,omitempty" json:"warehouse_id,omitempty"`

	Signatory *string `form:"signatory,omitempty" json:"signatory,omitempty"`
}

type SearchShipmentsJSONRequestBody = SearchShipmentsRequest

type ConfirmShipmentJSONRequestBody = ConfirmShipmentRequest

type SetShipmentPalletsCountJSONRequestBody = SetShipmentPalletsCountRequest

type GenerateShipmentListDocumentReportJSONRequestBody = GenerateShipmentListDocumentReportRequest
