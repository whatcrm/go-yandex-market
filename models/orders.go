package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	AffectedOrderQualityRatingComponentTypeDBSCANCELLATIONRATE AffectedOrderQualityRatingComponentType = "DBS_CANCELLATION_RATE"
	AffectedOrderQualityRatingComponentTypeDBSLATEDELIVERYRATE AffectedOrderQualityRatingComponentType = "DBS_LATE_DELIVERY_RATE"
	AffectedOrderQualityRatingComponentTypeFBSCANCELLATIONRATE AffectedOrderQualityRatingComponentType = "FBS_CANCELLATION_RATE"
	AffectedOrderQualityRatingComponentTypeFBSLATESHIPRATE     AffectedOrderQualityRatingComponentType = "FBS_LATE_SHIP_RATE"
)

const (
	CreateOrderPackageTypeBRAND      CreateOrderPackageType = "BRAND"
	CreateOrderPackageTypeWHITELABEL CreateOrderPackageType = "WHITELABEL"
)

const (
	OrderBuyerTypeBUSINESS OrderBuyerType = "BUSINESS"
	OrderBuyerTypePERSON   OrderBuyerType = "PERSON"
)

const (
	OrderDeliveryDispatchTypeBUYER               OrderDeliveryDispatchType = "BUYER"
	OrderDeliveryDispatchTypeMARKETBRANDEDOUTLET OrderDeliveryDispatchType = "MARKET_BRANDED_OUTLET"
	OrderDeliveryDispatchTypeSHOPOUTLET          OrderDeliveryDispatchType = "SHOP_OUTLET"
	OrderDeliveryDispatchTypeUNKNOWN             OrderDeliveryDispatchType = "UNKNOWN"
)

const (
	OrderDeliveryPartnerTypeSHOP         OrderDeliveryPartnerType = "SHOP"
	OrderDeliveryPartnerTypeUNKNOWN      OrderDeliveryPartnerType = "UNKNOWN"
	OrderDeliveryPartnerTypeYANDEXMARKET OrderDeliveryPartnerType = "YANDEX_MARKET"
)

const (
	OrderDeliveryTypeDELIVERY OrderDeliveryType = "DELIVERY"
	OrderDeliveryTypeDIGITAL  OrderDeliveryType = "DIGITAL"
	OrderDeliveryTypePICKUP   OrderDeliveryType = "PICKUP"
	OrderDeliveryTypePOST     OrderDeliveryType = "POST"
	OrderDeliveryTypeUNKNOWN  OrderDeliveryType = "UNKNOWN"
)

const (
	OrderItemStatusTypeREJECTED OrderItemStatusType = "REJECTED"
	OrderItemStatusTypeRETURNED OrderItemStatusType = "RETURNED"
)

const (
	OrderItemSubsidyTypeSUBSIDY        OrderItemSubsidyType = "SUBSIDY"
	OrderItemSubsidyTypeYANDEXCASHBACK OrderItemSubsidyType = "YANDEX_CASHBACK"
)

const (
	OrderLiftTypeCARGOELEVATOR OrderLiftType = "CARGO_ELEVATOR"
	OrderLiftTypeELEVATOR      OrderLiftType = "ELEVATOR"
	OrderLiftTypeFREE          OrderLiftType = "FREE"
	OrderLiftTypeMANUAL        OrderLiftType = "MANUAL"
	OrderLiftTypeNOTNEEDED     OrderLiftType = "NOT_NEEDED"
	OrderLiftTypeUNKNOWN       OrderLiftType = "UNKNOWN"
)

const (
	OrderPaymentMethodTypeAPPLEPAY              OrderPaymentMethodType = "APPLE_PAY"
	OrderPaymentMethodTypeB2BACCOUNTPOSTPAYMENT OrderPaymentMethodType = "B2B_ACCOUNT_POSTPAYMENT"
	OrderPaymentMethodTypeB2BACCOUNTPREPAYMENT  OrderPaymentMethodType = "B2B_ACCOUNT_PREPAYMENT"
	OrderPaymentMethodTypeBNPLBANKONDELIVERY    OrderPaymentMethodType = "BNPL_BANK_ON_DELIVERY"
	OrderPaymentMethodTypeBNPLONDELIVERY        OrderPaymentMethodType = "BNPL_ON_DELIVERY"
	OrderPaymentMethodTypeBNPLTBYB              OrderPaymentMethodType = "BNPL_TBYB"
	OrderPaymentMethodTypeBOUNDCARDONDELIVERY   OrderPaymentMethodType = "BOUND_CARD_ON_DELIVERY"
	OrderPaymentMethodTypeCARDONDELIVERY        OrderPaymentMethodType = "CARD_ON_DELIVERY"
	OrderPaymentMethodTypeCASHONDELIVERY        OrderPaymentMethodType = "CASH_ON_DELIVERY"
	OrderPaymentMethodTypeCREDIT                OrderPaymentMethodType = "CREDIT"
	OrderPaymentMethodTypeEXTERNALCERTIFICATE   OrderPaymentMethodType = "EXTERNAL_CERTIFICATE"
	OrderPaymentMethodTypeGOOGLEPAY             OrderPaymentMethodType = "GOOGLE_PAY"
	OrderPaymentMethodTypeMICROCREDIT           OrderPaymentMethodType = "MICROCREDIT"
	OrderPaymentMethodTypeSBP                   OrderPaymentMethodType = "SBP"
	OrderPaymentMethodTypeTINKOFFCREDIT         OrderPaymentMethodType = "TINKOFF_CREDIT"
	OrderPaymentMethodTypeTINKOFFINSTALLMENTS   OrderPaymentMethodType = "TINKOFF_INSTALLMENTS"
	OrderPaymentMethodTypeUNKNOWN               OrderPaymentMethodType = "UNKNOWN"
	OrderPaymentMethodTypeYANDEX                OrderPaymentMethodType = "YANDEX"
)

const (
	OrderPaymentTypePOSTPAID OrderPaymentType = "POSTPAID"
	OrderPaymentTypePREPAID  OrderPaymentType = "PREPAID"
	OrderPaymentTypeUNKNOWN  OrderPaymentType = "UNKNOWN"
)

const (
	OrderPromoTypeBLUEFLASH             OrderPromoType = "BLUE_FLASH"
	OrderPromoTypeBLUESET               OrderPromoType = "BLUE_SET"
	OrderPromoTypeCASHBACK              OrderPromoType = "CASHBACK"
	OrderPromoTypeCHEAPESTASGIFT        OrderPromoType = "CHEAPEST_AS_GIFT"
	OrderPromoTypeDCOEXTRADISCOUNT      OrderPromoType = "DCO_EXTRA_DISCOUNT"
	OrderPromoTypeDIRECTDISCOUNT        OrderPromoType = "DIRECT_DISCOUNT"
	OrderPromoTypeDISCOUNTBYPAYMENTTYPE OrderPromoType = "DISCOUNT_BY_PAYMENT_TYPE"
	OrderPromoTypeGENERICBUNDLE         OrderPromoType = "GENERIC_BUNDLE"
	OrderPromoTypeMARKETBLUE            OrderPromoType = "MARKET_BLUE"
	OrderPromoTypeMARKETCOIN            OrderPromoType = "MARKET_COIN"
	OrderPromoTypeMARKETCOUPON          OrderPromoType = "MARKET_COUPON"
	OrderPromoTypeMARKETPROMOCODE       OrderPromoType = "MARKET_PROMOCODE"
	OrderPromoTypePERCENTDISCOUNT       OrderPromoType = "PERCENT_DISCOUNT"
	OrderPromoTypePRICEDROPASYOUSHOP    OrderPromoType = "PRICE_DROP_AS_YOU_SHOP"
	OrderPromoTypeSECRETSALE            OrderPromoType = "SECRET_SALE"
	OrderPromoTypeSPREADDISCOUNTCOUNT   OrderPromoType = "SPREAD_DISCOUNT_COUNT"
	OrderPromoTypeSPREADDISCOUNTRECEIPT OrderPromoType = "SPREAD_DISCOUNT_RECEIPT"
	OrderPromoTypeUNKNOWN               OrderPromoType = "UNKNOWN"
)

const (
	OrderSourcePlatformTypeMARKET      OrderSourcePlatformType = "MARKET"
	OrderSourcePlatformTypeOTHER       OrderSourcePlatformType = "OTHER"
	OrderSourcePlatformTypeOZON        OrderSourcePlatformType = "OZON"
	OrderSourcePlatformTypeWILDBERRIES OrderSourcePlatformType = "WILDBERRIES"
)

const (
	OrderStatsStatusTypeCANCELLEDBEFOREPROCESSING OrderStatsStatusType = "CANCELLED_BEFORE_PROCESSING"
	OrderStatsStatusTypeCANCELLEDINDELIVERY       OrderStatsStatusType = "CANCELLED_IN_DELIVERY"
	OrderStatsStatusTypeCANCELLEDINPROCESSING     OrderStatsStatusType = "CANCELLED_IN_PROCESSING"
	OrderStatsStatusTypeDELIVERED                 OrderStatsStatusType = "DELIVERED"
	OrderStatsStatusTypeDELIVERY                  OrderStatsStatusType = "DELIVERY"
	OrderStatsStatusTypeLOST                      OrderStatsStatusType = "LOST"
	OrderStatsStatusTypePARTIALLYDELIVERED        OrderStatsStatusType = "PARTIALLY_DELIVERED"
	OrderStatsStatusTypePARTIALLYRETURNED         OrderStatsStatusType = "PARTIALLY_RETURNED"
	OrderStatsStatusTypePENDING                   OrderStatsStatusType = "PENDING"
	OrderStatsStatusTypePICKUP                    OrderStatsStatusType = "PICKUP"
	OrderStatsStatusTypePROCESSING                OrderStatsStatusType = "PROCESSING"
	OrderStatsStatusTypeRESERVED                  OrderStatsStatusType = "RESERVED"
	OrderStatsStatusTypeRETURNED                  OrderStatsStatusType = "RETURNED"
	OrderStatsStatusTypeUNKNOWN                   OrderStatsStatusType = "UNKNOWN"
	OrderStatsStatusTypeUNPAID                    OrderStatsStatusType = "UNPAID"
)

const (
	OrderStatusTypeCANCELLED         OrderStatusType = "CANCELLED"
	OrderStatusTypeDELIVERED         OrderStatusType = "DELIVERED"
	OrderStatusTypeDELIVERY          OrderStatusType = "DELIVERY"
	OrderStatusTypePARTIALLYRETURNED OrderStatusType = "PARTIALLY_RETURNED"
	OrderStatusTypePENDING           OrderStatusType = "PENDING"
	OrderStatusTypePICKUP            OrderStatusType = "PICKUP"
	OrderStatusTypePLACING           OrderStatusType = "PLACING"
	OrderStatusTypePROCESSING        OrderStatusType = "PROCESSING"
	OrderStatusTypeRESERVED          OrderStatusType = "RESERVED"
	OrderStatusTypeRETURNED          OrderStatusType = "RETURNED"
	OrderStatusTypeUNKNOWN           OrderStatusType = "UNKNOWN"
	OrderStatusTypeUNPAID            OrderStatusType = "UNPAID"
)

const (
	OrderSubsidyTypeDELIVERY       OrderSubsidyType = "DELIVERY"
	OrderSubsidyTypeSUBSIDY        OrderSubsidyType = "SUBSIDY"
	OrderSubsidyTypeYANDEXCASHBACK OrderSubsidyType = "YANDEX_CASHBACK"
)

const (
	OrderSubstatusTypeANTIFRAUD                                  OrderSubstatusType = "ANTIFRAUD"
	OrderSubstatusTypeASPARTOFMULTIORDER                         OrderSubstatusType = "AS_PART_OF_MULTI_ORDER"
	OrderSubstatusTypeASYNCPROCESSING                            OrderSubstatusType = "ASYNC_PROCESSING"
	OrderSubstatusTypeAWAITAUTODELIVERYDATES                     OrderSubstatusType = "AWAIT_AUTO_DELIVERY_DATES"
	OrderSubstatusTypeAWAITCASHIER                               OrderSubstatusType = "AWAIT_CASHIER"
	OrderSubstatusTypeAWAITCONFIRMATION                          OrderSubstatusType = "AWAIT_CONFIRMATION"
	OrderSubstatusTypeAWAITCUSTOMPRICECONFIRMATION               OrderSubstatusType = "AWAIT_CUSTOM_PRICE_CONFIRMATION"
	OrderSubstatusTypeAWAITDELIVERYDATES                         OrderSubstatusType = "AWAIT_DELIVERY_DATES"
	OrderSubstatusTypeAWAITDELIVERYDATESCONFIRMATION             OrderSubstatusType = "AWAIT_DELIVERY_DATES_CONFIRMATION"
	OrderSubstatusTypeAWAITLAVKARESERVATION                      OrderSubstatusType = "AWAIT_LAVKA_RESERVATION"
	OrderSubstatusTypeAWAITPAYMENT                               OrderSubstatusType = "AWAIT_PAYMENT"
	OrderSubstatusTypeAWAITPAYMENTAFTERDELIVERY                  OrderSubstatusType = "AWAIT_PAYMENT_AFTER_DELIVERY"
	OrderSubstatusTypeAWAITSERVICEABLECONFIRMATION               OrderSubstatusType = "AWAIT_SERVICEABLE_CONFIRMATION"
	OrderSubstatusTypeAWAITUSERPERSONALDATA                      OrderSubstatusType = "AWAIT_USER_PERSONAL_DATA"
	OrderSubstatusTypeAWAITUSERSTEAMFASTURL                      OrderSubstatusType = "AWAIT_USER_STEAM_FAST_URL"
	OrderSubstatusTypeBANKREJECTCREDITOFFER                      OrderSubstatusType = "BANK_REJECT_CREDIT_OFFER"
	OrderSubstatusTypeBROKENITEM                                 OrderSubstatusType = "BROKEN_ITEM"
	OrderSubstatusTypeCANCELLEDCOURIERNOTFOUND                   OrderSubstatusType = "CANCELLED_COURIER_NOT_FOUND"
	OrderSubstatusTypeCOURIERARRIVEDTOSENDER                     OrderSubstatusType = "COURIER_ARRIVED_TO_SENDER"
	OrderSubstatusTypeCOURIERFOUND                               OrderSubstatusType = "COURIER_FOUND"
	OrderSubstatusTypeCOURIERINTRANSITTOSENDER                   OrderSubstatusType = "COURIER_IN_TRANSIT_TO_SENDER"
	OrderSubstatusTypeCOURIERNOTCOMEFORORDER                     OrderSubstatusType = "COURIER_NOT_COME_FOR_ORDER"
	OrderSubstatusTypeCOURIERNOTDELIVERORDER                     OrderSubstatusType = "COURIER_NOT_DELIVER_ORDER"
	OrderSubstatusTypeCOURIERNOTFOUND                            OrderSubstatusType = "COURIER_NOT_FOUND"
	OrderSubstatusTypeCOURIERRECEIVED                            OrderSubstatusType = "COURIER_RECEIVED"
	OrderSubstatusTypeCOURIERRETURNEDORDER                       OrderSubstatusType = "COURIER_RETURNED_ORDER"
	OrderSubstatusTypeCOURIERRETURNSORDER                        OrderSubstatusType = "COURIER_RETURNS_ORDER"
	OrderSubstatusTypeCOURIERSEARCH                              OrderSubstatusType = "COURIER_SEARCH"
	OrderSubstatusTypeCOURIERSEARCHNOTSTARTED                    OrderSubstatusType = "COURIER_SEARCH_NOT_STARTED"
	OrderSubstatusTypeCREDITOFFERFAILED                          OrderSubstatusType = "CREDIT_OFFER_FAILED"
	OrderSubstatusTypeCUSTOM                                     OrderSubstatusType = "CUSTOM"
	OrderSubstatusTypeCUSTOMERREJECTCREDITOFFER                  OrderSubstatusType = "CUSTOMER_REJECT_CREDIT_OFFER"
	OrderSubstatusTypeCUSTOMSFAILEDMARKET                        OrderSubstatusType = "CUSTOMS_FAILED_MARKET"
	OrderSubstatusTypeCUSTOMSFAILEDUSERADDITIONALDATANOTPROVIDED OrderSubstatusType = "CUSTOMS_FAILED_USER_ADDITIONAL_DATA_NOT_PROVIDED"
	OrderSubstatusTypeCUSTOMSFAILEDUSERCOMMERCIALITEMS           OrderSubstatusType = "CUSTOMS_FAILED_USER_COMMERCIAL_ITEMS"
	OrderSubstatusTypeCUSTOMSFAILEDUSERDUTYNOTPAID               OrderSubstatusType = "CUSTOMS_FAILED_USER_DUTY_NOT_PAID"
	OrderSubstatusTypeCUSTOMSFAILEDUSERINVALIDPERSONALDATA       OrderSubstatusType = "CUSTOMS_FAILED_USER_INVALID_PERSONAL_DATA"
	OrderSubstatusTypeCUSTOMSPROBLEMS                            OrderSubstatusType = "CUSTOMS_PROBLEMS"
	OrderSubstatusTypeDAMAGEDBOX                                 OrderSubstatusType = "DAMAGED_BOX"
	OrderSubstatusTypeDEFERREDPAYMENT                            OrderSubstatusType = "DEFERRED_PAYMENT"
	OrderSubstatusTypeDELIVEREDUSERNOTRECEIVED                   OrderSubstatusType = "DELIVERED_USER_NOT_RECEIVED"
	OrderSubstatusTypeDELIVEREDUSERRECEIVED                      OrderSubstatusType = "DELIVERED_USER_RECEIVED"
	OrderSubstatusTypeDELIVERYCUSTOMSARRIVED                     OrderSubstatusType = "DELIVERY_CUSTOMS_ARRIVED"
	OrderSubstatusTypeDELIVERYCUSTOMSCLEARED                     OrderSubstatusType = "DELIVERY_CUSTOMS_CLEARED"
	OrderSubstatusTypeDELIVERYNOTMANAGEDREGION                   OrderSubstatusType = "DELIVERY_NOT_MANAGED_REGION"
	OrderSubstatusTypeDELIVERYPROBLEMS                           OrderSubstatusType = "DELIVERY_PROBLEMS"
	OrderSubstatusTypeDELIVERYSERVICEDELIVERED                   OrderSubstatusType = "DELIVERY_SERVICE_DELIVERED"
	OrderSubstatusTypeDELIVERYSERVICEFAILED                      OrderSubstatusType = "DELIVERY_SERVICE_FAILED"
	OrderSubstatusTypeDELIVERYSERVICELOST                        OrderSubstatusType = "DELIVERY_SERVICE_LOST"
	OrderSubstatusTypeDELIVERYSERVICENOTRECEIVED                 OrderSubstatusType = "DELIVERY_SERVICE_NOT_RECEIVED"
	OrderSubstatusTypeDELIVERYSERVICERECEIVED                    OrderSubstatusType = "DELIVERY_SERVICE_RECEIVED"
	OrderSubstatusTypeDELIVERYSERVICEUNDELIVERED                 OrderSubstatusType = "DELIVERY_SERVICE_UNDELIVERED"
	OrderSubstatusTypeDELIVERYTOSTORESTARTED                     OrderSubstatusType = "DELIVERY_TO_STORE_STARTED"
	OrderSubstatusTypeDELIVERYUSERNOTRECEIVED                    OrderSubstatusType = "DELIVERY_USER_NOT_RECEIVED"
	OrderSubstatusTypeDROPOFFCLOSED                              OrderSubstatusType = "DROPOFF_CLOSED"
	OrderSubstatusTypeDROPOFFLOST                                OrderSubstatusType = "DROPOFF_LOST"
	OrderSubstatusTypeFIRSTMILEDELIVERYSERVICERECEIVED           OrderSubstatusType = "FIRST_MILE_DELIVERY_SERVICE_RECEIVED"
	OrderSubstatusTypeFULLNOTRANSOM                              OrderSubstatusType = "FULL_NOT_RANSOM"
	OrderSubstatusTypeINAPPROPRIATEWEIGHTSIZE                    OrderSubstatusType = "INAPPROPRIATE_WEIGHT_SIZE"
	OrderSubstatusTypeINCOMPLETECONTACTINFORMATION               OrderSubstatusType = "INCOMPLETE_CONTACT_INFORMATION"
	OrderSubstatusTypeINCOMPLETEMULTIORDER                       OrderSubstatusType = "INCOMPLETE_MULTI_ORDER"
	OrderSubstatusTypeINCORRECTPERSONALDATA                      OrderSubstatusType = "INCORRECT_PERSONAL_DATA"
	OrderSubstatusTypeLASTMILECOURIERSEARCH                      OrderSubstatusType = "LAST_MILE_COURIER_SEARCH"
	OrderSubstatusTypeLASTMILESTARTED                            OrderSubstatusType = "LAST_MILE_STARTED"
	OrderSubstatusTypeLATECONTACT                                OrderSubstatusType = "LATE_CONTACT"
	OrderSubstatusTypeLEGALINFOCHANGED                           OrderSubstatusType = "LEGAL_INFO_CHANGED"
	OrderSubstatusTypeLOST                                       OrderSubstatusType = "LOST"
	OrderSubstatusTypeMISSINGITEM                                OrderSubstatusType = "MISSING_ITEM"
	OrderSubstatusTypeNOPERSONALDATAEXPIRED                      OrderSubstatusType = "NO_PERSONAL_DATA_EXPIRED"
	OrderSubstatusTypePACKAGING                                  OrderSubstatusType = "PACKAGING"
	OrderSubstatusTypePENDINGCANCELLED                           OrderSubstatusType = "PENDING_CANCELLED"
	OrderSubstatusTypePENDINGEXPIRED                             OrderSubstatusType = "PENDING_EXPIRED"
	OrderSubstatusTypePICKUPEXPIRED                              OrderSubstatusType = "PICKUP_EXPIRED"
	OrderSubstatusTypePICKUPPOINTCLOSED                          OrderSubstatusType = "PICKUP_POINT_CLOSED"
	OrderSubstatusTypePICKUPSERVICERECEIVED                      OrderSubstatusType = "PICKUP_SERVICE_RECEIVED"
	OrderSubstatusTypePICKUPUSERRECEIVED                         OrderSubstatusType = "PICKUP_USER_RECEIVED"
	OrderSubstatusTypePOSTPAIDBUDGETRESERVATIONFAILED            OrderSubstatusType = "POSTPAID_BUDGET_RESERVATION_FAILED"
	OrderSubstatusTypePOSTPAIDFAILED                             OrderSubstatusType = "POSTPAID_FAILED"
	OrderSubstatusTypePREORDER                                   OrderSubstatusType = "PREORDER"
	OrderSubstatusTypePRESCRIPTIONMISMATCH                       OrderSubstatusType = "PRESCRIPTION_MISMATCH"
	OrderSubstatusTypePROCESSINGEXPIRED                          OrderSubstatusType = "PROCESSING_EXPIRED"
	OrderSubstatusTypeREADYFORLASTMILE                           OrderSubstatusType = "READY_FOR_LAST_MILE"
	OrderSubstatusTypeREADYFORPICKUP                             OrderSubstatusType = "READY_FOR_PICKUP"
	OrderSubstatusTypeREADYTOSHIP                                OrderSubstatusType = "READY_TO_SHIP"
	OrderSubstatusTypeREPLACINGORDER                             OrderSubstatusType = "REPLACING_ORDER"
	OrderSubstatusTypeRESERVATIONEXPIRED                         OrderSubstatusType = "RESERVATION_EXPIRED"
	OrderSubstatusTypeRESERVATIONFAILED                          OrderSubstatusType = "RESERVATION_FAILED"
	OrderSubstatusTypeSERVICEFAULT                               OrderSubstatusType = "SERVICE_FAULT"
	OrderSubstatusTypeSHIPPED                                    OrderSubstatusType = "SHIPPED"
	OrderSubstatusTypeSHIPPEDTOWRONGDELIVERYSERVICE              OrderSubstatusType = "SHIPPED_TO_WRONG_DELIVERY_SERVICE"
	OrderSubstatusTypeSHOPFAILED                                 OrderSubstatusType = "SHOP_FAILED"
	OrderSubstatusTypeSHOPPENDINGCANCELLED                       OrderSubstatusType = "SHOP_PENDING_CANCELLED"
	OrderSubstatusTypeSORTINGCENTERLOST                          OrderSubstatusType = "SORTING_CENTER_LOST"
	OrderSubstatusTypeSTARTED                                    OrderSubstatusType = "STARTED"
	OrderSubstatusTypeTECHNICALERROR                             OrderSubstatusType = "TECHNICAL_ERROR"
	OrderSubstatusTypeTOOLONGDELIVERY                            OrderSubstatusType = "TOO_LONG_DELIVERY"
	OrderSubstatusTypeTOOMANYDELIVERYDATECHANGES                 OrderSubstatusType = "TOO_MANY_DELIVERY_DATE_CHANGES"
	OrderSubstatusTypeUNKNOWN                                    OrderSubstatusType = "UNKNOWN"
	OrderSubstatusTypeUSERBOUGHTCHEAPER                          OrderSubstatusType = "USER_BOUGHT_CHEAPER"
	OrderSubstatusTypeUSERCHANGEDMIND                            OrderSubstatusType = "USER_CHANGED_MIND"
	OrderSubstatusTypeUSERFORGOTTOUSEBONUS                       OrderSubstatusType = "USER_FORGOT_TO_USE_BONUS"
	OrderSubstatusTypeUSERFRAUD                                  OrderSubstatusType = "USER_FRAUD"
	OrderSubstatusTypeUSERHASNOTIMETOPICKUPORDER                 OrderSubstatusType = "USER_HAS_NO_TIME_TO_PICKUP_ORDER"
	OrderSubstatusTypeUSERIDENTIFICATIONMISMATCH                 OrderSubstatusType = "USER_IDENTIFICATION_MISMATCH"
	OrderSubstatusTypeUSERNOTPAID                                OrderSubstatusType = "USER_NOT_PAID"
	OrderSubstatusTypeUSERPLACEDOTHERORDER                       OrderSubstatusType = "USER_PLACED_OTHER_ORDER"
	OrderSubstatusTypeUSERRECEIVED                               OrderSubstatusType = "USER_RECEIVED"
	OrderSubstatusTypeUSERRECEIVEDTECHNICALERROR                 OrderSubstatusType = "USER_RECEIVED_TECHNICAL_ERROR"
	OrderSubstatusTypeUSERREFUSEDDELIVERY                        OrderSubstatusType = "USER_REFUSED_DELIVERY"
	OrderSubstatusTypeUSERREFUSEDPRODUCT                         OrderSubstatusType = "USER_REFUSED_PRODUCT"
	OrderSubstatusTypeUSERREFUSEDQUALITY                         OrderSubstatusType = "USER_REFUSED_QUALITY"
	OrderSubstatusTypeUSERUNREACHABLE                            OrderSubstatusType = "USER_UNREACHABLE"
	OrderSubstatusTypeUSERWANTEDANOTHERPAYMENTMETHOD             OrderSubstatusType = "USER_WANTED_ANOTHER_PAYMENT_METHOD"
	OrderSubstatusTypeUSERWANTSTOCHANGEADDRESS                   OrderSubstatusType = "USER_WANTS_TO_CHANGE_ADDRESS"
	OrderSubstatusTypeUSERWANTSTOCHANGEDELIVERYDATE              OrderSubstatusType = "USER_WANTS_TO_CHANGE_DELIVERY_DATE"
	OrderSubstatusTypeWAITINGBANKDECISION                        OrderSubstatusType = "WAITING_BANK_DECISION"
	OrderSubstatusTypeWAITINGFORSTOCKS                           OrderSubstatusType = "WAITING_FOR_STOCKS"
	OrderSubstatusTypeWAITINGPOSTPAIDBUDGETRESERVATION           OrderSubstatusType = "WAITING_POSTPAID_BUDGET_RESERVATION"
	OrderSubstatusTypeWAITINGTINKOFFDECISION                     OrderSubstatusType = "WAITING_TINKOFF_DECISION"
	OrderSubstatusTypeWAITINGUSERDELIVERYINPUT                   OrderSubstatusType = "WAITING_USER_DELIVERY_INPUT"
	OrderSubstatusTypeWAITINGUSERINPUT                           OrderSubstatusType = "WAITING_USER_INPUT"
	OrderSubstatusTypeWAREHOUSEFAILEDTOSHIP                      OrderSubstatusType = "WAREHOUSE_FAILED_TO_SHIP"
	OrderSubstatusTypeWRONGITEM                                  OrderSubstatusType = "WRONG_ITEM"
	OrderSubstatusTypeWRONGITEMDELIVERED                         OrderSubstatusType = "WRONG_ITEM_DELIVERED"
)

const (
	OrderTaxSystemTypeAUSN          OrderTaxSystemType = "AUSN"
	OrderTaxSystemTypeAUSNMINUSCOST OrderTaxSystemType = "AUSN_MINUS_COST"
	OrderTaxSystemTypeECHN          OrderTaxSystemType = "ECHN"
	OrderTaxSystemTypeENVD          OrderTaxSystemType = "ENVD"
	OrderTaxSystemTypeNPD           OrderTaxSystemType = "NPD"
	OrderTaxSystemTypeOSN           OrderTaxSystemType = "OSN"
	OrderTaxSystemTypePSN           OrderTaxSystemType = "PSN"
	OrderTaxSystemTypeUNKNOWNVALUE  OrderTaxSystemType = "UNKNOWN_VALUE"
	OrderTaxSystemTypeUSN           OrderTaxSystemType = "USN"
	OrderTaxSystemTypeUSNMINUSCOST  OrderTaxSystemType = "USN_MINUS_COST"
)

const (
	OrderUpdateStatusTypeERROR OrderUpdateStatusType = "ERROR"
	OrderUpdateStatusTypeOK    OrderUpdateStatusType = "OK"
)

const (
	OrderVatTypeNOVAT        OrderVatType = "NO_VAT"
	OrderVatTypeUNKNOWNVALUE OrderVatType = "UNKNOWN_VALUE"
	OrderVatTypeVAT0         OrderVatType = "VAT_0"
	OrderVatTypeVAT05        OrderVatType = "VAT_05"
	OrderVatTypeVAT07        OrderVatType = "VAT_07"
	OrderVatTypeVAT10        OrderVatType = "VAT_10"
	OrderVatTypeVAT10110     OrderVatType = "VAT_10_110"
	OrderVatTypeVAT12        OrderVatType = "VAT_12"
	OrderVatTypeVAT18        OrderVatType = "VAT_18"
	OrderVatTypeVAT18118     OrderVatType = "VAT_18_118"
	OrderVatTypeVAT20        OrderVatType = "VAT_20"
	OrderVatTypeVAT20120     OrderVatType = "VAT_20_120"
	OrderVatTypeVAT22        OrderVatType = "VAT_22"
)

const (
	OrdersStatsCommissionTypeAGENCY                    OrdersStatsCommissionType = "AGENCY"
	OrdersStatsCommissionTypeAUCTIONPROMOTION          OrdersStatsCommissionType = "AUCTION_PROMOTION"
	OrdersStatsCommissionTypeCROSSREGIONALDELIVERY     OrdersStatsCommissionType = "CROSSREGIONAL_DELIVERY"
	OrdersStatsCommissionTypeDELIVERYTOCUSTOMER        OrdersStatsCommissionType = "DELIVERY_TO_CUSTOMER"
	OrdersStatsCommissionTypeEXPRESSDELIVERYTOCUSTOMER OrdersStatsCommissionType = "EXPRESS_DELIVERY_TO_CUSTOMER"
	OrdersStatsCommissionTypeFEE                       OrdersStatsCommissionType = "FEE"
	OrdersStatsCommissionTypeFULFILLMENT               OrdersStatsCommissionType = "FULFILLMENT"
	OrdersStatsCommissionTypeILLIQUIDGOODSSALE         OrdersStatsCommissionType = "ILLIQUID_GOODS_SALE"
	OrdersStatsCommissionTypeINSTALLMENT               OrdersStatsCommissionType = "INSTALLMENT"
	OrdersStatsCommissionTypeINTAKESORTING             OrdersStatsCommissionType = "INTAKE_SORTING"
	OrdersStatsCommissionTypeLOYALTYPARTICIPATIONFEE   OrdersStatsCommissionType = "LOYALTY_PARTICIPATION_FEE"
	OrdersStatsCommissionTypePAYMENTTRANSFER           OrdersStatsCommissionType = "PAYMENT_TRANSFER"
	OrdersStatsCommissionTypeRETURNEDORDERSSTORAGE     OrdersStatsCommissionType = "RETURNED_ORDERS_STORAGE"
	OrdersStatsCommissionTypeRETURNPROCESSING          OrdersStatsCommissionType = "RETURN_PROCESSING"
	OrdersStatsCommissionTypeSORTING                   OrdersStatsCommissionType = "SORTING"
)

const (
	OrdersStatsItemStatusTypeREJECTED OrdersStatsItemStatusType = "REJECTED"
	OrdersStatsItemStatusTypeRETURNED OrdersStatsItemStatusType = "RETURNED"
)

const (
	OrdersStatsOrderPaymentTypePOSTPAID OrdersStatsOrderPaymentType = "POSTPAID"
	OrdersStatsOrderPaymentTypePREPAID  OrdersStatsOrderPaymentType = "PREPAID"
	OrdersStatsOrderPaymentTypeUNKNOWN  OrdersStatsOrderPaymentType = "UNKNOWN"
)

const (
	OrdersStatsPaymentSourceTypeBUYER         OrdersStatsPaymentSourceType = "BUYER"
	OrdersStatsPaymentSourceTypeCASHBACK      OrdersStatsPaymentSourceType = "CASHBACK"
	OrdersStatsPaymentSourceTypeMARKETCESSION OrdersStatsPaymentSourceType = "MARKET_CESSION"
	OrdersStatsPaymentSourceTypeMARKETPLACE   OrdersStatsPaymentSourceType = "MARKETPLACE"
	OrdersStatsPaymentSourceTypeSPLIT         OrdersStatsPaymentSourceType = "SPLIT"
)

const (
	OrdersStatsPriceTypeBUYER       OrdersStatsPriceType = "BUYER"
	OrdersStatsPriceTypeCASHBACK    OrdersStatsPriceType = "CASHBACK"
	OrdersStatsPriceTypeMARKETPLACE OrdersStatsPriceType = "MARKETPLACE"
)

const (
	OrdersStatsStockTypeDEFECT  OrdersStatsStockType = "DEFECT"
	OrdersStatsStockTypeEXPIRED OrdersStatsStockType = "EXPIRED"
	OrdersStatsStockTypeFIT     OrdersStatsStockType = "FIT"
)

type AcceptOrderCancellationRequest struct {
	Accepted bool `json:"accepted"`

	Reason *OrderCancellationReasonType `json:"reason,omitempty"`
}

type AffectedOrderQualityRatingComponentType string

type BasicOrderItemDTO struct {
	Count int32 `json:"count"`

	OfferId string `json:"offerId"`
}

type BriefOrderItemDTO struct {
	Count *int32 `json:"count,omitempty"`

	Id *int64 `json:"id,omitempty"`

	Instances *[]OrderItemInstanceDTO `json:"instances,omitempty"`

	OfferId *string `json:"offerId,omitempty"`

	OfferName *string `json:"offerName,omitempty"`

	Price *float32 `json:"price,omitempty"`

	Vat *OrderVatType `json:"vat,omitempty"`
}

type BriefOrderItemInstanceDTO struct {
	Cis *string `json:"cis,omitempty"`

	CountryCode *string `json:"countryCode,omitempty"`

	Gtd *string `json:"gtd,omitempty"`

	Rnpt *string `json:"rnpt,omitempty"`

	Uin *string `json:"uin,omitempty"`
}

type BusinessOrderBoxLayoutDTO struct {
	Barcode string `json:"barcode"`

	BoxId int64 `json:"boxId"`

	Items []BusinessOrderBoxLayoutItemDTO `json:"items"`
}

type BusinessOrderBoxLayoutItemDTO struct {
	FullCount *int32 `json:"fullCount,omitempty"`

	Id int64 `json:"id"`

	Instances *[]BriefOrderItemInstanceDTO `json:"instances,omitempty"`

	PartialCount *BusinessOrderBoxLayoutPartialCountDTO `json:"partialCount,omitempty"`
}

type BusinessOrderBoxLayoutPartialCountDTO struct {
	Current int32 `json:"current"`

	Total int32 `json:"total"`
}

type BusinessOrderCourierDeliveryDTO struct {
	Address *BusinessOrderDeliveryAddressDTO `json:"address,omitempty"`

	Region *RegionDTO `json:"region,omitempty"`
}

type BusinessOrderDTO struct {
	BuyerType *OrderBuyerType `json:"buyerType,omitempty"`

	CampaignId int64 `json:"campaignId"`

	CancelRequested *bool `json:"cancelRequested,omitempty"`

	CreationDate time.Time `json:"creationDate"`

	Delivery BusinessOrderDeliveryDTO `json:"delivery"`

	ExternalOrderId *string `json:"externalOrderId,omitempty"`

	Fake bool `json:"fake"`

	Items []BusinessOrderItemDTO `json:"items"`

	Notes *string `json:"notes,omitempty"`

	OrderId int64 `json:"orderId"`

	PaymentMethod OrderPaymentMethodType `json:"paymentMethod"`

	PaymentType OrderPaymentType `json:"paymentType"`

	Prices *OrderPriceDTO `json:"prices,omitempty"`

	ProgramType *SellingProgramType `json:"programType,omitempty"`

	Services *BusinessOrderServicesDTO `json:"services,omitempty"`

	SourcePlatform *OrderSourcePlatformType `json:"sourcePlatform,omitempty"`

	Status OrderStatusType `json:"status"`

	Substatus OrderSubstatusType `json:"substatus"`

	UpdateDate *time.Time `json:"updateDate,omitempty"`
}

type BusinessOrderDeliveryAddressDTO struct {
	Apartment *string `json:"apartment,omitempty"`

	Block *string `json:"block,omitempty"`

	City *string `json:"city,omitempty"`

	Country *string `json:"country,omitempty"`

	District *string `json:"district,omitempty"`

	Entrance *string `json:"entrance,omitempty"`

	Entryphone *string `json:"entryphone,omitempty"`

	Floor *string `json:"floor,omitempty"`

	Gps *GpsDTO `json:"gps,omitempty"`

	House *string `json:"house,omitempty"`

	Postcode *string `json:"postcode,omitempty"`

	Street *string `json:"street,omitempty"`

	Subway *string `json:"subway,omitempty"`
}

type BusinessOrderDeliveryDTO struct {
	BoxesLayout *[]BusinessOrderBoxLayoutDTO `json:"boxesLayout,omitempty"`

	Courier *BusinessOrderCourierDeliveryDTO `json:"courier,omitempty"`

	Dates BusinessOrderDeliveryDatesDTO `json:"dates"`

	DeliveryPartnerType OrderDeliveryPartnerType `json:"deliveryPartnerType"`

	DeliveryServiceId int64 `json:"deliveryServiceId"`

	DispatchType *OrderDeliveryDispatchType `json:"dispatchType,omitempty"`

	Estimated *bool `json:"estimated,omitempty"`

	Pickup *BusinessOrderPickupDeliveryDTO `json:"pickup,omitempty"`

	ReceiveBarcode *string `json:"receiveBarcode,omitempty"`

	ReceiveCode *string `json:"receiveCode,omitempty"`

	ServiceName string `json:"serviceName"`

	Shipment *BusinessOrderShipmentDTO `json:"shipment,omitempty"`

	Tracks *[]OrderTrackDTO `json:"tracks,omitempty"`

	Transfer *BusinessOrderTransferDTO `json:"transfer,omitempty"`

	Type OrderDeliveryType `json:"type"`

	WarehouseId *string `json:"warehouseId,omitempty"`
}

type BusinessOrderDeliveryDatesDTO struct {
	FromDate openapi_types.Date `json:"fromDate"`

	FromTime *string `json:"fromTime,omitempty"`

	RealDeliveryDate *openapi_types.Date `json:"realDeliveryDate,omitempty"`

	ToDate *openapi_types.Date `json:"toDate,omitempty"`

	ToTime *string `json:"toTime,omitempty"`
}

type BusinessOrderEacDTO struct {
	EacCode *string `json:"eacCode,omitempty"`

	EacType OrderDeliveryEacType `json:"eacType"`
}

type BusinessOrderItemDTO struct {
	Count int `json:"count"`

	Id int64 `json:"id"`

	Instances *[]OrderItemInstanceDTO `json:"instances,omitempty"`

	OfferId string `json:"offerId"`

	OfferName string `json:"offerName"`

	Prices *ItemPriceDTO `json:"prices,omitempty"`

	RequiredInstanceTypes *[]OrderItemInstanceType `json:"requiredInstanceTypes,omitempty"`

	Tags *[]OrderItemTagType `json:"tags,omitempty"`
}

type BusinessOrderPickupDeliveryDTO struct {
	Address *BusinessOrderDeliveryAddressDTO `json:"address,omitempty"`

	LogisticPointId *int64 `json:"logisticPointId,omitempty"`

	OutletCode *string `json:"outletCode,omitempty"`

	OutletStorageLimitDate *openapi_types.Date `json:"outletStorageLimitDate,omitempty"`

	Region *RegionDTO `json:"region,omitempty"`
}

type BusinessOrderServicesDTO struct {
	LiftType *OrderLiftType `json:"liftType,omitempty"`
}

type BusinessOrderShipmentDTO struct {
	Id *int64 `json:"id,omitempty"`

	ShipmentDate openapi_types.Date `json:"shipmentDate"`

	ShipmentTime *string `json:"shipmentTime,omitempty"`
}

type BusinessOrderTransferDTO struct {
	Courier *OrderCourierDTO `json:"courier,omitempty"`

	Eac *BusinessOrderEacDTO `json:"eac,omitempty"`
}

type CreateOrderDTO struct {
	Customer CustomerDTO `json:"customer"`

	Destination CreateOrderDeliveryOptionDTO `json:"destination"`

	Draft *bool `json:"draft,omitempty"`

	ExternalOrderId string `json:"externalOrderId"`

	Fake *bool `json:"fake,omitempty"`

	ItemsDelivery []CreateOrderWarehouseItemsDTO `json:"itemsDelivery"`

	Packaging CreateOrderPackagingDTO `json:"packaging"`

	PaymentType DeliveryPaymentType `json:"paymentType"`
}

type CreateOrderDeliveryOptionDTO struct {
	CourierDelivery *OrderCourierDeliveryDTO `json:"courierDelivery,omitempty"`

	PickupDelivery *OrderPickupDeliveryDTO `json:"pickupDelivery,omitempty"`
}

type CreateOrderItemDTO struct {
	Count int32 `json:"count"`

	OfferId string `json:"offerId"`

	Price *BasePriceDTO `json:"price,omitempty"`
}

type CreateOrderPackageType string

type CreateOrderPackagingDTO struct {
	PackageType *CreateOrderPackageType `json:"packageType,omitempty"`
}

type CreateOrderRequest struct {
	Order CreateOrderDTO `json:"order"`
}

type CreateOrderResponse struct {
	Result *CreatedOrdersDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type CreateOrderWarehouseItemsDTO struct {
	DeliveryDateInterval DeliveryDateIntervalDTO `json:"deliveryDateInterval"`

	DeliveryTimeInterval *TimeIntervalDTO `json:"deliveryTimeInterval,omitempty"`

	Items []CreateOrderItemDTO `json:"items"`

	WarehouseId int64 `json:"warehouseId"`
}

type CreatedOrderDTO struct {
	Id int64 `json:"id"`
}

type CreatedOrdersDTO struct {
	Orders []CreatedOrderDTO `json:"orders"`
}

type EnrichedOrderBoxLayoutDTO struct {
	BoxId *int64 `json:"boxId,omitempty"`

	Items []OrderBoxLayoutItemDTO `json:"items"`
}

type ExternalOrderId = string

type GetBusinessOrdersRequest struct {
	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	Dates *OrderDatesFilterDTO `json:"dates,omitempty"`

	ExternalOrderIds *[]ExternalOrderId `json:"externalOrderIds,omitempty"`

	Fake *bool `json:"fake,omitempty"`

	OrderIds *[]int64 `json:"orderIds,omitempty"`

	ProgramTypes *[]SellingProgramType `json:"programTypes,omitempty"`

	SourcePlatforms *[]OrderSourcePlatformType `json:"sourcePlatforms,omitempty"`

	Statuses *[]OrderStatusType `json:"statuses,omitempty"`

	Substatuses *[]OrderSubstatusType `json:"substatuses,omitempty"`

	WaitingForCancellationApprove *bool `json:"waitingForCancellationApprove,omitempty"`
}

type GetBusinessOrdersResponse struct {
	Orders []BusinessOrderDTO `json:"orders"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetOrderIdentifiersStatusDTO struct {
	Items []OrderItemValidationStatusDTO `json:"items"`
}

type GetOrderIdentifiersStatusResponse struct {
	Result *GetOrderIdentifiersStatusDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOrderResponse struct {
	Order *OrderDTO `json:"order,omitempty"`
}

type GetOrderUpdateOptionsRequest struct {
	Id int64 `json:"id"`
}

type GetOrderUpdateOptionsResponse struct {
	Result *OrderUpdateOptionsDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOrdersResponse struct {
	Orders []OrderDTO `json:"orders"`

	Pager *FlippingPagerDTO `json:"pager,omitempty"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OrderBoxLayoutDTO struct {
	Items []OrderBoxLayoutItemDTO `json:"items"`
}

type OrderBoxLayoutItemDTO struct {
	FullCount *int32 `json:"fullCount,omitempty"`

	Id int64 `json:"id"`

	Instances *[]BriefOrderItemInstanceDTO `json:"instances,omitempty"`

	PartialCount *OrderBoxLayoutPartialCountDTO `json:"partialCount,omitempty"`
}

type OrderBoxLayoutPartialCountDTO struct {
	Current int32 `json:"current"`

	Total int32 `json:"total"`
}

type OrderBoxesLayoutDTO struct {
	Boxes []EnrichedOrderBoxLayoutDTO `json:"boxes"`
}

type OrderBusinessBuyerDTO struct {
	Inn *string `json:"inn,omitempty"`

	Kpp *string `json:"kpp,omitempty"`

	OrganizationJurAddress *string `json:"organizationJurAddress,omitempty"`

	OrganizationName *string `json:"organizationName,omitempty"`
}

type OrderBusinessDocumentsDTO struct {
	Ksf *DocumentDTO `json:"ksf,omitempty"`

	Sf *DocumentDTO `json:"sf,omitempty"`

	TorgTwelve *DocumentDTO `json:"torgTwelve,omitempty"`

	Ukd *DocumentDTO `json:"ukd,omitempty"`

	Upd *DocumentDTO `json:"upd,omitempty"`
}

type OrderBuyerBasicInfoDTO struct {
	FirstName *string `json:"firstName,omitempty"`

	Id *string `json:"id,omitempty"`

	LastName *string `json:"lastName,omitempty"`

	MiddleName *string `json:"middleName,omitempty"`

	Type OrderBuyerType `json:"type"`
}

type OrderBuyerDTO struct {
	FirstName *string `json:"firstName,omitempty"`

	Id *string `json:"id,omitempty"`

	LastName *string `json:"lastName,omitempty"`

	MiddleName *string `json:"middleName,omitempty"`

	Type OrderBuyerType `json:"type"`
}

type OrderBuyerInfoDTO struct {
	FirstName *string `json:"firstName,omitempty"`

	Id *string `json:"id,omitempty"`

	LastName *string `json:"lastName,omitempty"`

	MiddleName *string `json:"middleName,omitempty"`

	Phone *string `json:"phone,omitempty"`

	Trusted *bool `json:"trusted,omitempty"`

	Type OrderBuyerType `json:"type"`
}

type OrderBuyerType string

type OrderCancellationReasonType string

type OrderCourierDTO struct {
	FullName *string `json:"fullName,omitempty"`

	Phone *string `json:"phone,omitempty"`

	PhoneExtension *string `json:"phoneExtension,omitempty"`

	VehicleDescription *string `json:"vehicleDescription,omitempty"`

	VehicleNumber *string `json:"vehicleNumber,omitempty"`
}

type OrderCourierDeliveryDTO struct {
	Address CourierDeliveryAddressDTO `json:"address"`

	Notes *string `json:"notes,omitempty"`
}

type OrderDTO struct {
	Buyer OrderBuyerDTO `json:"buyer"`

	BuyerItemsTotal *float32 `json:"buyerItemsTotal,omitempty"`

	BuyerItemsTotalBeforeDiscount float32 `json:"buyerItemsTotalBeforeDiscount"`

	BuyerTotal *float32 `json:"buyerTotal,omitempty"`

	BuyerTotalBeforeDiscount *float32 `json:"buyerTotalBeforeDiscount,omitempty"`

	CancelRequested *bool  `json:"cancelRequested,omitempty"`
	CreationDate    string `json:"creationDate"`

	Currency CurrencyType `json:"currency"`

	Delivery OrderDeliveryDTO `json:"delivery"`

	DeliveryTotal float32 `json:"deliveryTotal"`
	ExpiryDate    *string `json:"expiryDate,omitempty"`

	ExternalOrderId *string `json:"externalOrderId,omitempty"`

	Fake bool `json:"fake"`

	Id int64 `json:"id"`

	Items []OrderItemDTO `json:"items"`

	ItemsTotal float32 `json:"itemsTotal"`

	Notes *string `json:"notes,omitempty"`

	PaymentMethod OrderPaymentMethodType `json:"paymentMethod"`

	PaymentType OrderPaymentType `json:"paymentType"`

	Status OrderStatusType `json:"status"`

	Subsidies *[]OrderSubsidyDTO `json:"subsidies,omitempty"`

	Substatus OrderSubstatusType `json:"substatus"`

	TaxSystem OrderTaxSystemType `json:"taxSystem"`
	UpdatedAt *string            `json:"updatedAt,omitempty"`
}

type OrderDatesFilterDTO struct {
	CreationDateFrom *openapi_types.Date `json:"creationDateFrom,omitempty"`

	CreationDateTo *openapi_types.Date `json:"creationDateTo,omitempty"`

	ShipmentDateFrom *openapi_types.Date `json:"shipmentDateFrom,omitempty"`

	ShipmentDateTo *openapi_types.Date `json:"shipmentDateTo,omitempty"`

	UpdateDateFrom *time.Time `json:"updateDateFrom,omitempty"`

	UpdateDateTo *time.Time `json:"updateDateTo,omitempty"`
}

type OrderDeliveryAddressDTO struct {
	Apartment *string `json:"apartment,omitempty"`

	Block *string `json:"block,omitempty"`

	Building *string `json:"building,omitempty"`

	City *string `json:"city,omitempty"`

	Country *string `json:"country,omitempty"`

	District *string `json:"district,omitempty"`

	Entrance *string `json:"entrance,omitempty"`

	Entryphone *string `json:"entryphone,omitempty"`

	Estate *string `json:"estate,omitempty"`

	Floor *string `json:"floor,omitempty"`

	Gps *GpsDTO `json:"gps,omitempty"`

	House *string `json:"house,omitempty"`

	Phone *string `json:"phone,omitempty"`

	Postcode *string `json:"postcode,omitempty"`

	Recipient *string `json:"recipient,omitempty"`

	Street *string `json:"street,omitempty"`

	Subway *string `json:"subway,omitempty"`
}

type OrderDeliveryDTO struct {
	Address *OrderDeliveryAddressDTO `json:"address,omitempty"`

	Courier *OrderCourierDTO `json:"courier,omitempty"`

	Dates OrderDeliveryDatesDTO `json:"dates"`

	DeliveryPartnerType OrderDeliveryPartnerType `json:"deliveryPartnerType"`

	DeliveryServiceId int64 `json:"deliveryServiceId"`

	DispatchType *OrderDeliveryDispatchType `json:"dispatchType,omitempty"`

	EacCode *string `json:"eacCode,omitempty"`

	EacType *OrderDeliveryEacType `json:"eacType,omitempty"`

	Estimated *bool `json:"estimated,omitempty"`

	Id *string `json:"id,omitempty"`

	LiftPrice *float32 `json:"liftPrice,omitempty"`

	LiftType *OrderLiftType `json:"liftType,omitempty"`

	OutletCode *string `json:"outletCode,omitempty"`

	OutletStorageLimitDate *string `json:"outletStorageLimitDate,omitempty"`

	Price *float32 `json:"price,omitempty"`

	ReceiveCode *string `json:"receiveCode,omitempty"`

	Region *RegionDTO `json:"region,omitempty"`

	ServiceName string `json:"serviceName"`

	Shipments *[]OrderShipmentDTO `json:"shipments,omitempty"`

	Tracks *[]OrderTrackDTO `json:"tracks,omitempty"`

	Type OrderDeliveryType `json:"type"`

	Vat *OrderVatType `json:"vat,omitempty"`
}

type OrderDeliveryDateDTO struct {
	ToDate openapi_types.Date `json:"toDate"`
}

type OrderDeliveryDateReasonType string

type OrderDeliveryDatesDTO struct {
	FromDate string `json:"fromDate"`

	FromTime *string `json:"fromTime,omitempty"`

	RealDeliveryDate *string `json:"realDeliveryDate,omitempty"`

	ToDate *string `json:"toDate,omitempty"`

	ToTime *string `json:"toTime,omitempty"`
}

type OrderDeliveryDispatchType string

type OrderDeliveryEacType string

type OrderDeliveryPartnerType string

type OrderDeliveryType string

type OrderDigitalItemDTO struct {
	ActivateTill openapi_types.Date `json:"activate_till"`

	Codes *[]string `json:"codes,omitempty"`

	Id int64 `json:"id"`

	Slip string `json:"slip"`
}

type OrderDocumentStatusType string

type OrderItemDTO struct {
	BuyerPrice float32 `json:"buyerPrice"`

	BuyerPriceBeforeDiscount float32 `json:"buyerPriceBeforeDiscount"`

	Count int `json:"count"`

	Details *[]OrderItemDetailDTO `json:"details,omitempty"`

	Id int64 `json:"id"`

	Instances *[]OrderItemInstanceDTO `json:"instances,omitempty"`

	OfferId string `json:"offerId"`

	OfferName string `json:"offerName"`

	PartnerWarehouseId *string `json:"partnerWarehouseId,omitempty"`

	Price float32 `json:"price"`

	PriceBeforeDiscount *float32 `json:"priceBeforeDiscount,omitempty"`

	Promos *[]OrderItemPromoDTO `json:"promos,omitempty"`

	RequiredInstanceTypes *[]OrderItemInstanceType `json:"requiredInstanceTypes,omitempty"`

	ShopSku *string `json:"shopSku,omitempty"`

	Subsidies *[]OrderItemSubsidyDTO `json:"subsidies,omitempty"`

	Subsidy *float32 `json:"subsidy,omitempty"`

	Tags *[]OrderItemTagType `json:"tags,omitempty"`

	Vat *OrderVatType `json:"vat,omitempty"`
}

type OrderItemDetailDTO struct {
	ItemCount int64 `json:"itemCount"`

	ItemStatus OrderItemStatusType `json:"itemStatus"`

	UpdateDate string `json:"updateDate"`
}

type OrderItemInstanceDTO struct {
	Cis *string `json:"cis,omitempty"`

	CisFull *string `json:"cisFull,omitempty"`

	CountryCode *string `json:"countryCode,omitempty"`

	Gtd *string `json:"gtd,omitempty"`

	Rnpt *string `json:"rnpt,omitempty"`

	Uin *string `json:"uin,omitempty"`
}

type OrderItemInstanceModificationDTO struct {
	Id int64 `json:"id"`

	Instances []BriefOrderItemInstanceDTO `json:"instances"`
}

type OrderItemInstanceType string

type OrderItemModificationDTO struct {
	Count int32 `json:"count"`

	Id int64 `json:"id"`

	Instances *[]BriefOrderItemInstanceDTO `json:"instances,omitempty"`
}

type OrderItemPromoDTO struct {
	Discount *float32 `json:"discount,omitempty"`

	MarketPromoId *string `json:"marketPromoId,omitempty"`

	ShopPromoId *string `json:"shopPromoId,omitempty"`

	Subsidy float32 `json:"subsidy"`

	Type OrderPromoType `json:"type"`
}

type OrderItemStatusType string

type OrderItemSubsidyDTO struct {
	Amount float32 `json:"amount"`

	Type OrderItemSubsidyType `json:"type"`
}

type OrderItemSubsidyType string

type OrderItemTagType string

type OrderItemValidationStatusDTO struct {
	Cis *[]CisDTO `json:"cis,omitempty"`

	Id int64 `json:"id"`

	Uin *[]UinDTO `json:"uin,omitempty"`
}

type OrderItemsModificationRequestReasonType string

type OrderItemsModificationResultDTO struct {
	Items []BriefOrderItemDTO `json:"items"`
}

type OrderLabelDTO struct {
	OrderId int64 `json:"orderId"`

	ParcelBoxLabels []ParcelBoxLabelDTO `json:"parcelBoxLabels"`

	PlacesNumber int32  `json:"placesNumber"`
	Url          string `json:"url"`
}

type OrderLiftType string

type OrderParcelBoxDTO struct {
	FulfilmentId string `json:"fulfilmentId"`

	Id int64 `json:"id"`
}

type OrderPaymentMethodType string

type OrderPaymentType string

type OrderPickupDeliveryDTO struct {
	LogisticPointId int64 `json:"logisticPointId"`
}

type OrderPickupReturnDTO struct {
	LogisticPointId int64 `json:"logisticPointId"`
}

type OrderPriceDTO struct {
	Cashback *CurrencyValueDTO `json:"cashback,omitempty"`

	Delivery *DeliveryPriceDTO `json:"delivery,omitempty"`

	Payment *CurrencyValueDTO `json:"payment,omitempty"`

	Subsidy *CurrencyValueDTO `json:"subsidy,omitempty"`
}

type OrderPromoType string

type OrderShipmentDTO struct {
	Boxes *[]OrderParcelBoxDTO `json:"boxes,omitempty"`

	Id *int64 `json:"id,omitempty"`

	ShipmentDate *string `json:"shipmentDate,omitempty"`

	ShipmentTime *string `json:"shipmentTime,omitempty"`

	Tracks *[]OrderTrackDTO `json:"tracks,omitempty"`
}

type OrderSourcePlatformType string

type OrderStateDTO struct {
	Id int64 `json:"id"`

	Status OrderStatusType `json:"status"`

	Substatus *OrderSubstatusType `json:"substatus,omitempty"`
}

type OrderStatsStatusType string

type OrderStatusChangeDTO struct {
	Delivery *OrderStatusChangeDeliveryDTO `json:"delivery,omitempty"`

	Status OrderStatusType `json:"status"`

	Substatus *OrderSubstatusType `json:"substatus,omitempty"`
}

type OrderStatusChangeDeliveryDTO struct {
	Dates *OrderStatusChangeDeliveryDatesDTO `json:"dates,omitempty"`
}

type OrderStatusChangeDeliveryDatesDTO struct {
	RealDeliveryDate *openapi_types.Date `json:"realDeliveryDate,omitempty"`
}

type OrderStatusType string

type OrderSubsidyDTO struct {
	Amount float32 `json:"amount"`

	Type OrderSubsidyType `json:"type"`
}

type OrderSubsidyType string

type OrderSubstatusType string

type OrderTaxSystemType string

type OrderTrackDTO struct {
	DeliveryServiceId int64 `json:"deliveryServiceId"`

	TrackCode *string `json:"trackCode,omitempty"`
}

type OrderUpdateOptionsDTO struct {
	DeliveryIntervals DeliveryIntervalsUpdateOptionsDTO `json:"deliveryIntervals"`
}

type OrderUpdateStatusType string

type OrderVatType string

type OrdersShipmentInfoDTO struct {
	OrderIdsWithLabels []int64 `json:"orderIdsWithLabels"`

	OrderIdsWithoutLabels []int64 `json:"orderIdsWithoutLabels"`
}

type OrdersStatsCommissionDTO struct {
	Actual *float32 `json:"actual,omitempty"`

	Type *OrdersStatsCommissionType `json:"type,omitempty"`
}

type OrdersStatsCommissionType string

type OrdersStatsDTO struct {
	Orders []OrdersStatsOrderDTO `json:"orders"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OrdersStatsDeliveryRegionDTO struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type OrdersStatsDetailsDTO struct {
	ItemCount *int64 `json:"itemCount,omitempty"`

	ItemStatus *OrdersStatsItemStatusType `json:"itemStatus,omitempty"`

	StockType *OrdersStatsStockType `json:"stockType,omitempty"`

	UpdateDate *openapi_types.Date `json:"updateDate,omitempty"`
}

type OrdersStatsItemDTO struct {
	BidFee *int32 `json:"bidFee,omitempty"`

	CisList *[]string `json:"cisList,omitempty"`

	CofinanceThreshold *float32 `json:"cofinanceThreshold,omitempty"`

	CofinanceValue *float32 `json:"cofinanceValue,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Details *[]OrdersStatsDetailsDTO `json:"details,omitempty"`

	InitialCount *int32 `json:"initialCount,omitempty"`

	MarketSku *int64 `json:"marketSku,omitempty"`

	OfferName *string `json:"offerName,omitempty"`

	Prices *[]OrdersStatsPriceDTO `json:"prices,omitempty"`

	ShopSku *string `json:"shopSku,omitempty"`

	Warehouse *OrdersStatsWarehouseDTO `json:"warehouse,omitempty"`
}

type OrdersStatsItemStatusType string

type OrdersStatsOrderDTO struct {
	Commissions []OrdersStatsCommissionDTO `json:"commissions"`

	CreationDate *openapi_types.Date `json:"creationDate,omitempty"`

	Currency CurrencyType `json:"currency"`

	DeliveryRegion *OrdersStatsDeliveryRegionDTO `json:"deliveryRegion,omitempty"`

	Fake *bool `json:"fake,omitempty"`

	Id *int64 `json:"id,omitempty"`

	InitialItems *[]OrdersStatsItemDTO `json:"initialItems,omitempty"`

	Items []OrdersStatsItemDTO `json:"items"`

	PartnerOrderId *string `json:"partnerOrderId,omitempty"`

	PaymentType *OrdersStatsOrderPaymentType `json:"paymentType,omitempty"`

	Payments []OrdersStatsPaymentDTO `json:"payments"`

	Status *OrderStatsStatusType `json:"status,omitempty"`

	StatusUpdateDate *time.Time `json:"statusUpdateDate,omitempty"`

	Subsidies *[]OrdersStatsSubsidyDTO `json:"subsidies,omitempty"`
}

type OrdersStatsOrderPaymentType string

type OrdersStatsPaymentDTO struct {
	Date *openapi_types.Date `json:"date,omitempty"`

	Id *string `json:"id,omitempty"`

	PaymentOrder *OrdersStatsPaymentOrderDTO `json:"paymentOrder,omitempty"`

	Source *OrdersStatsPaymentSourceType `json:"source,omitempty"`

	Total *float32 `json:"total,omitempty"`

	Type *OrdersStatsPaymentType `json:"type,omitempty"`
}

type OrdersStatsPaymentOrderDTO struct {
	Date *openapi_types.Date `json:"date,omitempty"`

	Id *string `json:"id,omitempty"`
}

type OrdersStatsPaymentSourceType string

type OrdersStatsPaymentType string

type OrdersStatsPriceDTO struct {
	CostPerItem *float32 `json:"costPerItem,omitempty"`

	Total *float32 `json:"total,omitempty"`

	Type *OrdersStatsPriceType `json:"type,omitempty"`
}

type OrdersStatsPriceType string

type OrdersStatsStockType string

type OrdersStatsSubsidyDTO struct {
	Amount float32 `json:"amount"`

	OperationType OrdersStatsSubsidyOperationType `json:"operationType"`

	Type OrdersStatsSubsidyType `json:"type"`
}

type OrdersStatsSubsidyOperationType string

type OrdersStatsSubsidyType string

type OrdersStatsWarehouseDTO struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type ProvideOrderDigitalCodesRequest struct {
	Items []OrderDigitalItemDTO `json:"items"`
}

type ProvideOrderItemIdentifiersRequest struct {
	Items []OrderItemInstanceModificationDTO `json:"items"`
}

type ProvideOrderItemIdentifiersResponse struct {
	Result *OrderItemsModificationResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type QualityRatingAffectedOrderDTO struct {
	ComponentType AffectedOrderQualityRatingComponentType `json:"componentType"`

	Description string `json:"description"`

	OrderId int64 `json:"orderId"`
}

type QuestionSortOrderType string

type SetOrderBoxLayoutRequest struct {
	AllowRemove *bool `json:"allowRemove,omitempty"`

	Boxes []OrderBoxLayoutDTO `json:"boxes"`
}

type SetOrderBoxLayoutResponse struct {
	Result *OrderBoxesLayoutDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SetOrderShipmentBoxesRequest = ParcelRequestDTO

type SetOrderShipmentBoxesResponse struct {
	Result *ShipmentBoxesDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SortOrderType string

type UpdateExternalOrderIdRequest struct {
	ExternalOrderId string `json:"externalOrderId"`
}

type UpdateOrderDTO struct {
	Customer *CustomerDTO `json:"customer,omitempty"`

	DeliveryInterval *DeliveryIntervalsUpdateOptionDTO `json:"deliveryInterval,omitempty"`

	Id int64 `json:"id"`
}

type UpdateOrderItemRequest struct {
	Items []OrderItemModificationDTO `json:"items"`

	Reason *OrderItemsModificationRequestReasonType `json:"reason,omitempty"`
}

type UpdateOrderRequest struct {
	Order UpdateOrderDTO `json:"order"`
}

type UpdateOrderResponse struct {
	Result *UpdateOrderResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type UpdateOrderResultDTO struct {
	Operations []OperationDTO `json:"operations"`
}

type UpdateOrderStatusDTO struct {
	ErrorDetails *string `json:"errorDetails,omitempty"`

	Id *int64 `json:"id,omitempty"`

	Operation *OperationDTO `json:"operation,omitempty"`

	Status *OrderStatusType `json:"status,omitempty"`

	Substatus *OrderSubstatusType `json:"substatus,omitempty"`

	UpdateStatus *OrderUpdateStatusType `json:"updateStatus,omitempty"`
}

type UpdateOrderStatusRequest struct {
	Order OrderStatusChangeDTO `json:"order"`
}

type UpdateOrderStatusResponse struct {
	Operation *OperationDTO `json:"operation,omitempty"`

	Order *OrderDTO `json:"order,omitempty"`
}

type UpdateOrderStatusesDTO struct {
	Orders []UpdateOrderStatusDTO `json:"orders"`
}

type UpdateOrderStatusesRequest struct {
	Orders []OrderStateDTO `json:"orders"`
}

type UpdateOrderStatusesResponse struct {
	Result *UpdateOrderStatusesDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetBusinessOrdersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetOrdersParams struct {
	OrderIds *[]int64 `form:"orderIds,omitempty" json:"orderIds,omitempty"`

	Status *[]OrderStatusType `form:"status,omitempty" json:"status,omitempty"`

	Substatus *[]OrderSubstatusType `form:"substatus,omitempty" json:"substatus,omitempty"`

	FromDate *openapi_types.Date `form:"fromDate,omitempty" json:"fromDate,omitempty"`

	ToDate *openapi_types.Date `form:"toDate,omitempty" json:"toDate,omitempty"`

	SupplierShipmentDateFrom *openapi_types.Date `form:"supplierShipmentDateFrom,omitempty" json:"supplierShipmentDateFrom,omitempty"`

	SupplierShipmentDateTo *openapi_types.Date `form:"supplierShipmentDateTo,omitempty" json:"supplierShipmentDateTo,omitempty"`

	UpdatedAtFrom *time.Time `form:"updatedAtFrom,omitempty" json:"updatedAtFrom,omitempty"`

	UpdatedAtTo *time.Time `form:"updatedAtTo,omitempty" json:"updatedAtTo,omitempty"`

	DispatchType *OrderDeliveryDispatchType `form:"dispatchType,omitempty" json:"dispatchType,omitempty"`

	Fake *bool `form:"fake,omitempty" json:"fake,omitempty"`

	HasCis *bool `form:"hasCis,omitempty" json:"hasCis,omitempty"`

	OnlyWaitingForCancellationApprove *bool `form:"onlyWaitingForCancellationApprove,omitempty" json:"onlyWaitingForCancellationApprove,omitempty"`

	OnlyEstimatedDelivery *bool `form:"onlyEstimatedDelivery,omitempty" json:"onlyEstimatedDelivery,omitempty"`

	BuyerType *OrderBuyerType `form:"buyerType,omitempty" json:"buyerType,omitempty"`

	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	PageSize *int32 `form:"pageSize,omitempty" json:"pageSize,omitempty"`

	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GenerateOrderLabelsParams struct {
	Format *PageFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateOrderLabelParams struct {
	Format *PageFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GetOrdersStatsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GenerateMassOrderLabelsReportParams struct {
	Format *PageFormatType `form:"format,omitempty" json:"format,omitempty"`
}

type GenerateUnitedOrdersReportParams struct {
	Format *ReportFormatType `form:"format,omitempty" json:"format,omitempty"`

	Language *ReportLanguageType `form:"language,omitempty" json:"language,omitempty"`
}

type GetBusinessOrdersJSONRequestBody = GetBusinessOrdersRequest

type CreateOrderJSONRequestBody = CreateOrderRequest

type UpdateOrderJSONRequestBody = UpdateOrderRequest

type GetOrderUpdateOptionsJSONRequestBody = GetOrderUpdateOptionsRequest

type TransferOrdersFromShipmentJSONRequestBody = TransferOrdersFromShipmentRequest

type UpdateOrderStatusesJSONRequestBody = UpdateOrderStatusesRequest

type SetOrderBoxLayoutJSONRequestBody = SetOrderBoxLayoutRequest

type AcceptOrderCancellationJSONRequestBody = AcceptOrderCancellationRequest

type ProvideOrderDigitalCodesJSONRequestBody = ProvideOrderDigitalCodesRequest

type SetOrderDeliveryDateJSONRequestBody = SetOrderDeliveryDateRequest

type SetOrderShipmentBoxesJSONRequestBody = SetOrderShipmentBoxesRequest

type UpdateOrderStorageLimitJSONRequestBody = UpdateOrderStorageLimitRequest

type SetOrderDeliveryTrackCodeJSONRequestBody = SetOrderDeliveryTrackCodeRequest

type UpdateExternalOrderIdJSONRequestBody = UpdateExternalOrderIdRequest

type ProvideOrderItemIdentifiersJSONRequestBody = ProvideOrderItemIdentifiersRequest

type UpdateOrderItemsJSONRequestBody = UpdateOrderItemRequest

type UpdateOrderStatusJSONRequestBody = UpdateOrderStatusRequest

type VerifyOrderEacJSONRequestBody = VerifyOrderEacRequest

type GetOrdersStatsJSONRequestBody = GetOrdersStatsRequest

type GenerateMassOrderLabelsReportJSONRequestBody = GenerateMassOrderLabelsRequest

type GenerateUnitedOrdersReportJSONRequestBody = GenerateUnitedOrdersRequest
