package models

import "time"

const (
	SupplyRequestDocumentTypeACTOFDISCREPANCY                  SupplyRequestDocumentType = "ACT_OF_DISCREPANCY"
	SupplyRequestDocumentTypeACTOFRECEPTIONTRANSFER            SupplyRequestDocumentType = "ACT_OF_RECEPTION_TRANSFER"
	SupplyRequestDocumentTypeACTOFWITHDRAW                     SupplyRequestDocumentType = "ACT_OF_WITHDRAW"
	SupplyRequestDocumentTypeACTOFWITHDRAWFROMSTORAGE          SupplyRequestDocumentType = "ACT_OF_WITHDRAW_FROM_STORAGE"
	SupplyRequestDocumentTypeADDITIONALSUPPLY                  SupplyRequestDocumentType = "ADDITIONAL_SUPPLY"
	SupplyRequestDocumentTypeADDITIONALSUPPLYACCEPTABLEGOODS   SupplyRequestDocumentType = "ADDITIONAL_SUPPLY_ACCEPTABLE_GOODS"
	SupplyRequestDocumentTypeADDITIONALSUPPLYUNACCEPTABLEGOODS SupplyRequestDocumentType = "ADDITIONAL_SUPPLY_UNACCEPTABLE_GOODS"
	SupplyRequestDocumentTypeANOMALYCONTAINERSWITHDRAWACT      SupplyRequestDocumentType = "ANOMALY_CONTAINERS_WITHDRAW_ACT"
	SupplyRequestDocumentTypeCARGOUNITS                        SupplyRequestDocumentType = "CARGO_UNITS"
	SupplyRequestDocumentTypeCISFACT                           SupplyRequestDocumentType = "CIS_FACT"
	SupplyRequestDocumentTypeIDENTIFIERS                       SupplyRequestDocumentType = "IDENTIFIERS"
	SupplyRequestDocumentTypeINBOUNDUTD                        SupplyRequestDocumentType = "INBOUND_UTD"
	SupplyRequestDocumentTypeITEMSWITHCISES                    SupplyRequestDocumentType = "ITEMS_WITH_CISES"
	SupplyRequestDocumentTypeOUTBOUNDUTD                       SupplyRequestDocumentType = "OUTBOUND_UTD"
	SupplyRequestDocumentTypeREPORTOFWITHDRAWWITHCISES         SupplyRequestDocumentType = "REPORT_OF_WITHDRAW_WITH_CISES"
	SupplyRequestDocumentTypeRNPTFACT                          SupplyRequestDocumentType = "RNPT_FACT"
	SupplyRequestDocumentTypeSECONDARYACCEPTANCECISES          SupplyRequestDocumentType = "SECONDARY_ACCEPTANCE_CISES"
	SupplyRequestDocumentTypeSECONDARYRECEPTIONACT             SupplyRequestDocumentType = "SECONDARY_RECEPTION_ACT"
	SupplyRequestDocumentTypeSUPPLY                            SupplyRequestDocumentType = "SUPPLY"
	SupplyRequestDocumentTypeTRANSFER                          SupplyRequestDocumentType = "TRANSFER"
	SupplyRequestDocumentTypeVALIDATIONERRORS                  SupplyRequestDocumentType = "VALIDATION_ERRORS"
	SupplyRequestDocumentTypeVIRTUALDISTRIBUTIONCENTERSUPPLY   SupplyRequestDocumentType = "VIRTUAL_DISTRIBUTION_CENTER_SUPPLY"
	SupplyRequestDocumentTypeWITHDRAW                          SupplyRequestDocumentType = "WITHDRAW"
)

const (
	SupplyRequestLocationTypeFULFILLMENT SupplyRequestLocationType = "FULFILLMENT"
	SupplyRequestLocationTypePICKUPPOINT SupplyRequestLocationType = "PICKUP_POINT"
	SupplyRequestLocationTypeXDOC        SupplyRequestLocationType = "XDOC"
)

const (
	SupplyRequestReferenceTypeADDITIONALSUPPLY    SupplyRequestReferenceType = "ADDITIONAL_SUPPLY"
	SupplyRequestReferenceTypeUTILIZATION         SupplyRequestReferenceType = "UTILIZATION"
	SupplyRequestReferenceTypeVIRTUALDISTRIBUTION SupplyRequestReferenceType = "VIRTUAL_DISTRIBUTION"
	SupplyRequestReferenceTypeWITHDRAW            SupplyRequestReferenceType = "WITHDRAW"
)

const (
	SupplyRequestSortAttributeTypeID            SupplyRequestSortAttributeType = "ID"
	SupplyRequestSortAttributeTypeREQUESTEDDATE SupplyRequestSortAttributeType = "REQUESTED_DATE"
	SupplyRequestSortAttributeTypeSTATUS        SupplyRequestSortAttributeType = "STATUS"
	SupplyRequestSortAttributeTypeUPDATEDAT     SupplyRequestSortAttributeType = "UPDATED_AT"
)

const (
	SupplyRequestStatusTypeACCEPTEDBYWAREHOUSESYSTEM   SupplyRequestStatusType = "ACCEPTED_BY_WAREHOUSE_SYSTEM"
	SupplyRequestStatusTypeARRIVEDTOSERVICE            SupplyRequestStatusType = "ARRIVED_TO_SERVICE"
	SupplyRequestStatusTypeARRIVEDTOXDOCSERVICE        SupplyRequestStatusType = "ARRIVED_TO_XDOC_SERVICE"
	SupplyRequestStatusTypeCANCELLATIONREJECTED        SupplyRequestStatusType = "CANCELLATION_REJECTED"
	SupplyRequestStatusTypeCANCELLATIONREQUESTED       SupplyRequestStatusType = "CANCELLATION_REQUESTED"
	SupplyRequestStatusTypeCANCELLED                   SupplyRequestStatusType = "CANCELLED"
	SupplyRequestStatusTypeCREATED                     SupplyRequestStatusType = "CREATED"
	SupplyRequestStatusTypeFINISHED                    SupplyRequestStatusType = "FINISHED"
	SupplyRequestStatusTypeINVALID                     SupplyRequestStatusType = "INVALID"
	SupplyRequestStatusTypeNEEDPREPARATION             SupplyRequestStatusType = "NEED_PREPARATION"
	SupplyRequestStatusTypePUBLISHED                   SupplyRequestStatusType = "PUBLISHED"
	SupplyRequestStatusTypeREADYFORUTILIZATION         SupplyRequestStatusType = "READY_FOR_UTILIZATION"
	SupplyRequestStatusTypeREADYTOWITHDRAW             SupplyRequestStatusType = "READY_TO_WITHDRAW"
	SupplyRequestStatusTypeREGISTEREDINELECTRONICQUEUE SupplyRequestStatusType = "REGISTERED_IN_ELECTRONIC_QUEUE"
	SupplyRequestStatusTypeSHIPPEDTOSERVICE            SupplyRequestStatusType = "SHIPPED_TO_SERVICE"
	SupplyRequestStatusTypeTRANSITMOVING               SupplyRequestStatusType = "TRANSIT_MOVING"
	SupplyRequestStatusTypeVALIDATED                   SupplyRequestStatusType = "VALIDATED"
	SupplyRequestStatusTypeWAREHOUSEHANDLING           SupplyRequestStatusType = "WAREHOUSE_HANDLING"
	SupplyRequestStatusTypeWAREHOUSESIGNEDACT          SupplyRequestStatusType = "WAREHOUSE_SIGNED_ACT"
)

const (
	SupplyRequestSubTypeADDITIONALSUPPLY                            SupplyRequestSubType = "ADDITIONAL_SUPPLY"
	SupplyRequestSubTypeANOMALYWITHDRAW                             SupplyRequestSubType = "ANOMALY_WITHDRAW"
	SupplyRequestSubTypeDEFAULT                                     SupplyRequestSubType = "DEFAULT"
	SupplyRequestSubTypeEXTERNALWITHDRAWINTOZON                     SupplyRequestSubType = "EXTERNAL_WITHDRAW_INT_OZON"
	SupplyRequestSubTypeEXTERNALWITHDRAWINTWB                       SupplyRequestSubType = "EXTERNAL_WITHDRAW_INT_WB"
	SupplyRequestSubTypeFIXLOSTINVENTORYING                         SupplyRequestSubType = "FIX_LOST_INVENTORYING"
	SupplyRequestSubTypeFORCEPLAN                                   SupplyRequestSubType = "FORCE_PLAN"
	SupplyRequestSubTypeFORCEPLANANOMALYPERSUPPLY                   SupplyRequestSubType = "FORCE_PLAN_ANOMALY_PER_SUPPLY"
	SupplyRequestSubTypeINVENTORYINGSUPPLY                          SupplyRequestSubType = "INVENTORYING_SUPPLY"
	SupplyRequestSubTypeINVENTORYINGSUPPLYWAREHOUSEBASEDPERSUPPLIER SupplyRequestSubType = "INVENTORYING_SUPPLY_WAREHOUSE_BASED_PER_SUPPLIER"
	SupplyRequestSubTypeMANUTIL                                     SupplyRequestSubType = "MAN_UTIL"
	SupplyRequestSubTypeMISGRADINGSUPPLY                            SupplyRequestSubType = "MISGRADING_SUPPLY"
	SupplyRequestSubTypeMISGRADINGWITHDRAW                          SupplyRequestSubType = "MISGRADING_WITHDRAW"
	SupplyRequestSubTypeMOVEMENTSUPPLY                              SupplyRequestSubType = "MOVEMENT_SUPPLY"
	SupplyRequestSubTypeMOVEMENTWITHDRAW                            SupplyRequestSubType = "MOVEMENT_WITHDRAW"
	SupplyRequestSubTypeOPERLOSTINVENTORYING                        SupplyRequestSubType = "OPER_LOST_INVENTORYING"
	SupplyRequestSubTypePLANBYSUPPLIER                              SupplyRequestSubType = "PLAN_BY_SUPPLIER"
	SupplyRequestSubTypeVIRTUALDISTRIBUTIONCENTER                   SupplyRequestSubType = "VIRTUAL_DISTRIBUTION_CENTER"
	SupplyRequestSubTypeVIRTUALDISTRIBUTIONCENTERCHILD              SupplyRequestSubType = "VIRTUAL_DISTRIBUTION_CENTER_CHILD"
	SupplyRequestSubTypeWITHDRAWAUTOUTILIZATION                     SupplyRequestSubType = "WITHDRAW_AUTO_UTILIZATION"
	SupplyRequestSubTypeXDOC                                        SupplyRequestSubType = "XDOC"
)

const (
	SupplyRequestTypeSUPPLY      SupplyRequestType = "SUPPLY"
	SupplyRequestTypeUTILIZATION SupplyRequestType = "UTILIZATION"
	SupplyRequestTypeWITHDRAW    SupplyRequestType = "WITHDRAW"
)

type GetSupplyRequestDocumentsDTO struct {
	Documents []SupplyRequestDocumentDTO `json:"documents"`
}

type GetSupplyRequestDocumentsRequest struct {
	RequestId int64 `json:"requestId"`
}

type GetSupplyRequestDocumentsResponse struct {
	Result *GetSupplyRequestDocumentsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetSupplyRequestItemsDTO struct {
	Items []SupplyRequestItemDTO `json:"items"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetSupplyRequestItemsRequest struct {
	RequestId int64 `json:"requestId"`
}

type GetSupplyRequestItemsResponse struct {
	Result *GetSupplyRequestItemsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetSupplyRequestsDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Requests []SupplyRequestDTO `json:"requests"`
}

type GetSupplyRequestsRequest struct {
	RequestDateFrom *time.Time `json:"requestDateFrom,omitempty"`

	RequestDateTo *time.Time `json:"requestDateTo,omitempty"`

	RequestIds *[]SupplyRequestId `json:"requestIds,omitempty"`

	RequestStatuses *[]SupplyRequestStatusType `json:"requestStatuses,omitempty"`

	RequestSubtypes *[]SupplyRequestSubType `json:"requestSubtypes,omitempty"`

	RequestTypes *[]SupplyRequestType `json:"requestTypes,omitempty"`

	Sorting *SupplyRequestSortingDTO `json:"sorting,omitempty"`
}

type GetSupplyRequestsResponse struct {
	Result *GetSupplyRequestsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SupplyRequestCountersDTO struct {
	AcceptableCount *int32 `json:"acceptableCount,omitempty"`

	ActualBoxCount *int32 `json:"actualBoxCount,omitempty"`

	ActualPalletsCount *int32 `json:"actualPalletsCount,omitempty"`

	DefectCount *int32 `json:"defectCount,omitempty"`

	FactCount *int32 `json:"factCount,omitempty"`

	PlanCount *int32 `json:"planCount,omitempty"`

	ShortageCount *int32 `json:"shortageCount,omitempty"`

	SurplusCount *int32 `json:"surplusCount,omitempty"`

	UnacceptableCount *int32 `json:"unacceptableCount,omitempty"`

	UndefinedCount *int32 `json:"undefinedCount,omitempty"`
}

type SupplyRequestDTO struct {
	ChildrenLinks *[]SupplyRequestReferenceDTO `json:"childrenLinks,omitempty"`

	Counters SupplyRequestCountersDTO `json:"counters"`

	Id SupplyRequestIdDTO `json:"id"`

	ParentLink *SupplyRequestReferenceDTO `json:"parentLink,omitempty"`

	Status SupplyRequestStatusType `json:"status"`

	Subtype SupplyRequestSubType `json:"subtype"`

	TargetLocation SupplyRequestLocationDTO `json:"targetLocation"`

	TransitLocation *SupplyRequestLocationDTO `json:"transitLocation,omitempty"`

	Type SupplyRequestType `json:"type"`

	UpdatedAt time.Time `json:"updatedAt"`
}

type SupplyRequestDocumentDTO struct {
	CreatedAt time.Time `json:"createdAt"`

	Type SupplyRequestDocumentType `json:"type"`
	Url  string                    `json:"url"`
}

type SupplyRequestDocumentType string

type SupplyRequestId = int64

type SupplyRequestIdDTO struct {
	Id int64 `json:"id"`

	MarketplaceRequestId *string `json:"marketplaceRequestId,omitempty"`

	WarehouseRequestId *string `json:"warehouseRequestId,omitempty"`
}

type SupplyRequestItemCountersDTO struct {
	DefectCount *int32 `json:"defectCount,omitempty"`

	FactCount *int32 `json:"factCount,omitempty"`

	PlanCount *int32 `json:"planCount,omitempty"`

	ShortageCount *int32 `json:"shortageCount,omitempty"`

	SurplusCount *int32 `json:"surplusCount,omitempty"`
}

type SupplyRequestItemDTO struct {
	Counters SupplyRequestItemCountersDTO `json:"counters"`

	Name string `json:"name"`

	OfferId string `json:"offerId"`

	Price *CurrencyValueDTO `json:"price,omitempty"`
}

type SupplyRequestLocationAddressDTO struct {
	FullAddress string `json:"fullAddress"`

	Gps GpsDTO `json:"gps"`
}

type SupplyRequestLocationDTO struct {
	Address SupplyRequestLocationAddressDTO `json:"address"`

	Name string `json:"name"`

	RequestedDate *time.Time `json:"requestedDate,omitempty"`

	ServiceId int64 `json:"serviceId"`

	Type SupplyRequestLocationType `json:"type"`
}

type SupplyRequestLocationType string

type SupplyRequestReferenceDTO struct {
	Id SupplyRequestIdDTO `json:"id"`

	Type SupplyRequestReferenceType `json:"type"`
}

type SupplyRequestReferenceType string

type SupplyRequestSortAttributeType string

type SupplyRequestSortingDTO struct {
	Attribute SupplyRequestSortAttributeType `json:"attribute"`

	Direction SortOrderType `json:"direction"`
}

type SupplyRequestStatusType string

type SupplyRequestSubType string

type SupplyRequestType string

type GetSupplyRequestsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetSupplyRequestItemsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetSupplyRequestsJSONRequestBody = GetSupplyRequestsRequest

type GetSupplyRequestDocumentsJSONRequestBody = GetSupplyRequestDocumentsRequest

type GetSupplyRequestItemsJSONRequestBody = GetSupplyRequestItemsRequest
