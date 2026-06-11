package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	ExternalReturnDecisionReasonTypeBADQUALITY ExternalReturnDecisionReasonType = "BAD_QUALITY"
	ExternalReturnDecisionReasonTypeDOESNOTFIT ExternalReturnDecisionReasonType = "DOES_NOT_FIT"
	ExternalReturnDecisionReasonTypeWRONGITEM  ExternalReturnDecisionReasonType = "WRONG_ITEM"
)

const (
	ExternalReturnDecisionSubreasonTypeBADPACKAGE             ExternalReturnDecisionSubreasonType = "BAD_PACKAGE"
	ExternalReturnDecisionSubreasonTypeDAMAGED                ExternalReturnDecisionSubreasonType = "DAMAGED"
	ExternalReturnDecisionSubreasonTypeDELIVEREDTOOLONG       ExternalReturnDecisionSubreasonType = "DELIVERED_TOO_LONG"
	ExternalReturnDecisionSubreasonTypeDIDNOTMATCHDESCRIPTION ExternalReturnDecisionSubreasonType = "DID_NOT_MATCH_DESCRIPTION"
	ExternalReturnDecisionSubreasonTypeINCOMPLETENESS         ExternalReturnDecisionSubreasonType = "INCOMPLETENESS"
	ExternalReturnDecisionSubreasonTypeNOTWORKING             ExternalReturnDecisionSubreasonType = "NOT_WORKING"
	ExternalReturnDecisionSubreasonTypeUSERCHANGEDMIND        ExternalReturnDecisionSubreasonType = "USER_CHANGED_MIND"
	ExternalReturnDecisionSubreasonTypeUSERDIDNOTLIKE         ExternalReturnDecisionSubreasonType = "USER_DID_NOT_LIKE"
	ExternalReturnDecisionSubreasonTypeWRONGCOLOR             ExternalReturnDecisionSubreasonType = "WRONG_COLOR"
	ExternalReturnDecisionSubreasonTypeWRONGITEM              ExternalReturnDecisionSubreasonType = "WRONG_ITEM"
)

const (
	ReturnDecisionReasonTypeBADQUALITY     ReturnDecisionReasonType = "BAD_QUALITY"
	ReturnDecisionReasonTypeCONTENTFAIL    ReturnDecisionReasonType = "CONTENT_FAIL"
	ReturnDecisionReasonTypeDAMAGEDELIVERY ReturnDecisionReasonType = "DAMAGE_DELIVERY"
	ReturnDecisionReasonTypeDELIVERYFAIL   ReturnDecisionReasonType = "DELIVERY_FAIL"
	ReturnDecisionReasonTypeDOESNOTFIT     ReturnDecisionReasonType = "DOES_NOT_FIT"
	ReturnDecisionReasonTypeLOYALTYFAIL    ReturnDecisionReasonType = "LOYALTY_FAIL"
	ReturnDecisionReasonTypeUNKNOWN        ReturnDecisionReasonType = "UNKNOWN"
	ReturnDecisionReasonTypeWRONGITEM      ReturnDecisionReasonType = "WRONG_ITEM"
)

const (
	ReturnDecisionSubreasonTypeBADFLOWERS             ReturnDecisionSubreasonType = "BAD_FLOWERS"
	ReturnDecisionSubreasonTypeBADPACKAGE             ReturnDecisionSubreasonType = "BAD_PACKAGE"
	ReturnDecisionSubreasonTypeBROKEN                 ReturnDecisionSubreasonType = "BROKEN"
	ReturnDecisionSubreasonTypeDAMAGED                ReturnDecisionSubreasonType = "DAMAGED"
	ReturnDecisionSubreasonTypeDELIVEREDTOOLONG       ReturnDecisionSubreasonType = "DELIVERED_TOO_LONG"
	ReturnDecisionSubreasonTypeDIDNOTMATCHDESCRIPTION ReturnDecisionSubreasonType = "DID_NOT_MATCH_DESCRIPTION"
	ReturnDecisionSubreasonTypeINCOMPLETE             ReturnDecisionSubreasonType = "INCOMPLETE"
	ReturnDecisionSubreasonTypeINCOMPLETENESS         ReturnDecisionSubreasonType = "INCOMPLETENESS"
	ReturnDecisionSubreasonTypeITEMWASUSED            ReturnDecisionSubreasonType = "ITEM_WAS_USED"
	ReturnDecisionSubreasonTypeNOTWORKING             ReturnDecisionSubreasonType = "NOT_WORKING"
	ReturnDecisionSubreasonTypePARCELMISSING          ReturnDecisionSubreasonType = "PARCEL_MISSING"
	ReturnDecisionSubreasonTypeUNKNOWN                ReturnDecisionSubreasonType = "UNKNOWN"
	ReturnDecisionSubreasonTypeUSERCHANGEDMIND        ReturnDecisionSubreasonType = "USER_CHANGED_MIND"
	ReturnDecisionSubreasonTypeUSERDIDNOTLIKE         ReturnDecisionSubreasonType = "USER_DID_NOT_LIKE"
	ReturnDecisionSubreasonTypeWRAPPINGDAMAGED        ReturnDecisionSubreasonType = "WRAPPING_DAMAGED"
	ReturnDecisionSubreasonTypeWRONGAMOUNTDELIVERED   ReturnDecisionSubreasonType = "WRONG_AMOUNT_DELIVERED"
	ReturnDecisionSubreasonTypeWRONGCOLOR             ReturnDecisionSubreasonType = "WRONG_COLOR"
	ReturnDecisionSubreasonTypeWRONGITEM              ReturnDecisionSubreasonType = "WRONG_ITEM"
	ReturnDecisionSubreasonTypeWRONGORDER             ReturnDecisionSubreasonType = "WRONG_ORDER"
)

const (
	ReturnDecisionTypeDECLINEREFUND                ReturnDecisionType = "DECLINE_REFUND"
	ReturnDecisionTypeFASTREFUNDMONEY              ReturnDecisionType = "FAST_REFUND_MONEY"
	ReturnDecisionTypeOTHERDECISION                ReturnDecisionType = "OTHER_DECISION"
	ReturnDecisionTypePARTIALMONEYREFUND           ReturnDecisionType = "PARTIAL_MONEY_REFUND"
	ReturnDecisionTypeREFUNDMONEY                  ReturnDecisionType = "REFUND_MONEY"
	ReturnDecisionTypeREFUNDMONEYINCLUDINGSHIPMENT ReturnDecisionType = "REFUND_MONEY_INCLUDING_SHIPMENT"
	ReturnDecisionTypeREPAIR                       ReturnDecisionType = "REPAIR"
	ReturnDecisionTypeREPLACE                      ReturnDecisionType = "REPLACE"
	ReturnDecisionTypeSENDTOEXAMINATION            ReturnDecisionType = "SEND_TO_EXAMINATION"
	ReturnDecisionTypeUNKNOWN                      ReturnDecisionType = "UNKNOWN"
)

const (
	ReturnInstanceStatusTypeCANCELLED              ReturnInstanceStatusType = "CANCELLED"
	ReturnInstanceStatusTypeCREATED                ReturnInstanceStatusType = "CREATED"
	ReturnInstanceStatusTypeEXPROPRIATED           ReturnInstanceStatusType = "EXPROPRIATED"
	ReturnInstanceStatusTypeINTRANSIT              ReturnInstanceStatusType = "IN_TRANSIT"
	ReturnInstanceStatusTypeLOST                   ReturnInstanceStatusType = "LOST"
	ReturnInstanceStatusTypeNOTINDEMAND            ReturnInstanceStatusType = "NOT_IN_DEMAND"
	ReturnInstanceStatusTypePICKED                 ReturnInstanceStatusType = "PICKED"
	ReturnInstanceStatusTypePREPAREDFORUTILIZATION ReturnInstanceStatusType = "PREPARED_FOR_UTILIZATION"
	ReturnInstanceStatusTypeREADYFORPICKUP         ReturnInstanceStatusType = "READY_FOR_PICKUP"
	ReturnInstanceStatusTypeRECEIVED               ReturnInstanceStatusType = "RECEIVED"
	ReturnInstanceStatusTypeRECEIVEDONFULFILLMENT  ReturnInstanceStatusType = "RECEIVED_ON_FULFILLMENT"
	ReturnInstanceStatusTypeUTILIZED               ReturnInstanceStatusType = "UTILIZED"
)

const (
	ReturnInstanceStockTypeANOMALY               ReturnInstanceStockType = "ANOMALY"
	ReturnInstanceStockTypeDEFECT                ReturnInstanceStockType = "DEFECT"
	ReturnInstanceStockTypeDEMO                  ReturnInstanceStockType = "DEMO"
	ReturnInstanceStockTypeEXPIRED               ReturnInstanceStockType = "EXPIRED"
	ReturnInstanceStockTypeFIRMWARE              ReturnInstanceStockType = "FIRMWARE"
	ReturnInstanceStockTypeFIT                   ReturnInstanceStockType = "FIT"
	ReturnInstanceStockTypeINCORRECTCIS          ReturnInstanceStockType = "INCORRECT_CIS"
	ReturnInstanceStockTypeINCORRECTIMEI         ReturnInstanceStockType = "INCORRECT_IMEI"
	ReturnInstanceStockTypeINCORRECTSERIALNUMBER ReturnInstanceStockType = "INCORRECT_SERIAL_NUMBER"
	ReturnInstanceStockTypeMARKDOWN              ReturnInstanceStockType = "MARKDOWN"
	ReturnInstanceStockTypeMISGRADING            ReturnInstanceStockType = "MISGRADING"
	ReturnInstanceStockTypeNONCOMPLIENT          ReturnInstanceStockType = "NON_COMPLIENT"
	ReturnInstanceStockTypeNOTACCEPTABLE         ReturnInstanceStockType = "NOT_ACCEPTABLE"
	ReturnInstanceStockTypePARTMISSING           ReturnInstanceStockType = "PART_MISSING"
	ReturnInstanceStockTypeREPAIR                ReturnInstanceStockType = "REPAIR"
	ReturnInstanceStockTypeSERVICE               ReturnInstanceStockType = "SERVICE"
	ReturnInstanceStockTypeSURPLUS               ReturnInstanceStockType = "SURPLUS"
	ReturnInstanceStockTypeUNDEFINED             ReturnInstanceStockType = "UNDEFINED"
	ReturnInstanceStockTypeUNKNOWN               ReturnInstanceStockType = "UNKNOWN"
)

type CancelReturnDTO struct {
	Operation OperationDTO `json:"operation"`
}

type CancelReturnRequest struct {
	ReturnId int64 `json:"returnId"`
}

type CancelReturnResponse struct {
	Result *CancelReturnDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type CreateReturnDTO struct {
	Customer CustomerDTO `json:"customer"`

	ExternalReturnId string `json:"externalReturnId"`

	Items []CreateReturnItemDTO `json:"items"`

	OrderId int64 `json:"orderId"`

	ReturnOption CreateReturnOptionDTO `json:"returnOption"`
}

type CreateReturnItemDTO struct {
	Comment *string `json:"comment,omitempty"`

	Count int32 `json:"count"`

	OfferId string `json:"offerId"`

	Pictures *[]string `json:"pictures,omitempty"`

	ReasonType ExternalReturnDecisionReasonType `json:"reasonType"`

	SubreasonType ExternalReturnDecisionSubreasonType `json:"subreasonType"`
}

type CreateReturnOptionDTO struct {
	PickupReturn OrderPickupReturnDTO `json:"pickupReturn"`
}

type CreateReturnRequest struct {
	NewReturn CreateReturnDTO `json:"newReturn"`
}

type CreateReturnResponse struct {
	Result *CreatedReturnDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type CreatedReturnDTO struct {
	Id int64 `json:"id"`
}

type ExternalReturnDecisionReasonType string

type ExternalReturnDecisionSubreasonType string

type GetReturnAvailableDecisionsRequest struct {
	CampaignId int64 `json:"campaignId"`

	ReturnId int64 `json:"returnId"`
}

type GetReturnAvailableDecisionsResponse struct {
	Result *ReturnAvailableDecisionsResponse `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetReturnDeliveryOptionsDTO struct {
	PickupDelivery PickupReturnDeliveryOptionsDTO `json:"pickupDelivery"`
}

type GetReturnResponse struct {
	Result *ReturnDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetReturnsResponse struct {
	Result *PagedReturnsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type PagedReturnsDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Returns []ReturnDTO `json:"returns"`
}

type PickupReturnDeliveryOptionsDTO struct {
	LogisticPointIds []LogisticPointId `json:"logisticPointIds"`
}

type ReturnAvailableDecisionDTO struct {
	DecisionReasonTypes *[]ReturnRequestDecisionReasonType `json:"decisionReasonTypes,omitempty"`

	DecisionType ReturnDecisionType `json:"decisionType"`

	PartialCompensationBounds *PartialCompensationBoundsDTO `json:"partialCompensationBounds,omitempty"`
}

type ReturnAvailableDecisionsResponse struct {
	AvailableDecisions []ReturnAvailableDecisionDTO `json:"availableDecisions"`
}

type ReturnDTO struct {
	Amount *CurrencyValueDTO `json:"amount,omitempty"`

	CreationDate *time.Time `json:"creationDate,omitempty"`

	FastReturn *bool `json:"fastReturn,omitempty"`

	Id int64 `json:"id"`

	Items []ReturnItemDTO `json:"items"`

	LogisticPickupPoint *LogisticPickupPointDTO `json:"logisticPickupPoint,omitempty"`

	OrderId int64 `json:"orderId"`

	PickupTillDate *time.Time `json:"pickupTillDate,omitempty"`

	RefundAmount *int64 `json:"refundAmount,omitempty"`

	RefundStatus *RefundStatusType `json:"refundStatus,omitempty"`

	ReturnType ReturnType `json:"returnType"`

	ShipmentRecipientType *RecipientType `json:"shipmentRecipientType,omitempty"`

	ShipmentStatus *ReturnShipmentStatusType `json:"shipmentStatus,omitempty"`

	UpdateDate *time.Time `json:"updateDate,omitempty"`
}

type ReturnDecisionDTO struct {
	Amount *CurrencyValueDTO `json:"amount,omitempty"`

	Comment *string `json:"comment,omitempty"`

	Count *int32 `json:"count,omitempty"`

	DecisionType *ReturnDecisionType `json:"decisionType,omitempty"`

	Images *[]string `json:"images,omitempty"`

	PartnerCompensation *int64 `json:"partnerCompensation,omitempty"`

	PartnerCompensationAmount *CurrencyValueDTO `json:"partnerCompensationAmount,omitempty"`

	ReasonType *ReturnDecisionReasonType `json:"reasonType,omitempty"`

	RefundAmount *int64 `json:"refundAmount,omitempty"`

	ReturnItemId *int64 `json:"returnItemId,omitempty"`

	SubreasonType *ReturnDecisionSubreasonType `json:"subreasonType,omitempty"`
}

type ReturnDecisionReasonType string

type ReturnDecisionSubreasonType string

type ReturnDecisionType string

type ReturnInstanceDTO struct {
	Cis *string `json:"cis,omitempty"`

	Imei *string `json:"imei,omitempty"`

	Status *ReturnInstanceStatusType `json:"status,omitempty"`

	StockType *ReturnInstanceStockType `json:"stockType,omitempty"`
}

type ReturnInstanceStatusType string

type ReturnInstanceStockType string

type ReturnItemDTO struct {
	Count int64 `json:"count"`

	Decisions *[]ReturnDecisionDTO `json:"decisions,omitempty"`

	Instances *[]ReturnInstanceDTO `json:"instances,omitempty"`

	MarketSku *int64 `json:"marketSku,omitempty"`

	ShopSku string `json:"shopSku"`

	Tracks *[]TrackDTO `json:"tracks,omitempty"`
}

type ReturnItemDecisionDTO struct {
	Comment *string `json:"comment,omitempty"`

	Compensation *BasePriceDTO `json:"compensation,omitempty"`

	DecisionReasonType *ReturnRequestDecisionReasonType `json:"decisionReasonType,omitempty"`

	DecisionType ReturnRequestDecisionType `json:"decisionType"`

	ReturnItemId int64 `json:"returnItemId"`
}

type ReturnRequestDecisionReasonType string

type ReturnRequestDecisionType string

type ReturnType string

type SetReturnDecisionRequest struct {
	Comment *string `json:"comment,omitempty"`

	Compensation *BasePriceDTO `json:"compensation,omitempty"`

	DecisionType ReturnRequestDecisionType `json:"decisionType"`

	ReturnItemId int64 `json:"returnItemId"`
}

type SubmitReturnDecisionRequest struct {
	ReturnItemDecisions []ReturnItemDecisionDTO `json:"returnItemDecisions"`
}

type GetReturnsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	OrderIds *[]int64 `form:"orderIds,omitempty" json:"orderIds,omitempty"`

	Statuses *[]RefundStatusType `form:"statuses,omitempty" json:"statuses,omitempty"`

	ShipmentStatuses *[]ReturnShipmentStatusType `form:"shipmentStatuses,omitempty" json:"shipmentStatuses,omitempty"`

	Type *ReturnType `form:"type,omitempty" json:"type,omitempty"`

	FromDate *openapi_types.Date `form:"fromDate,omitempty" json:"fromDate,omitempty"`

	ToDate *openapi_types.Date `form:"toDate,omitempty" json:"toDate,omitempty"`

	FromDateLegacy *openapi_types.Date `form:"from_date,omitempty" json:"from_date,omitempty"`

	ToDateLegacy *openapi_types.Date `form:"to_date,omitempty" json:"to_date,omitempty"`
}

type GenerateUnitedReturnsReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GetReturnAvailableDecisionsJSONRequestBody = GetReturnAvailableDecisionsRequest

type GetReturnDeliveryOptionsJSONRequestBody = GetReturnDeliveryOptionsRequest

type CancelReturnJSONRequestBody = CancelReturnRequest

type CreateReturnJSONRequestBody = CreateReturnRequest

type SetReturnDecisionJSONRequestBody = SetReturnDecisionRequest

type SubmitReturnDecisionJSONRequestBody = SubmitReturnDecisionRequest

type GenerateUnitedReturnsReportJSONRequestBody = GenerateUnitedReturnsRequest
