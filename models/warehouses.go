package models

const (
	WarehouseComponentTypeADDRESS WarehouseComponentType = "ADDRESS"
	WarehouseComponentTypeSTATUS  WarehouseComponentType = "STATUS"
)

const (
	WarehouseStatusTypeDISABLEDMANUALLY WarehouseStatusType = "DISABLED_MANUALLY"
	WarehouseStatusTypeOTHER            WarehouseStatusType = "OTHER"
)

const (
	WarehouseStockTypeAVAILABLE   WarehouseStockType = "AVAILABLE"
	WarehouseStockTypeDEFECT      WarehouseStockType = "DEFECT"
	WarehouseStockTypeEXPIRED     WarehouseStockType = "EXPIRED"
	WarehouseStockTypeFIT         WarehouseStockType = "FIT"
	WarehouseStockTypeFREEZE      WarehouseStockType = "FREEZE"
	WarehouseStockTypeQUARANTINE  WarehouseStockType = "QUARANTINE"
	WarehouseStockTypeUTILIZATION WarehouseStockType = "UTILIZATION"
)

type FulfillmentWarehouseDTO struct {
	Address *WarehouseAddressDTO `json:"address,omitempty"`

	Id int64 `json:"id"`

	Name string `json:"name"`
}

type FulfillmentWarehousesDTO struct {
	Warehouses []FulfillmentWarehouseDTO `json:"warehouses"`
}

type GetFulfillmentWarehousesResponse struct {
	Result *FulfillmentWarehousesDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPagedWarehousesRequest struct {
	CampaignIds *[]CampaignId `json:"campaignIds,omitempty"`

	Components *[]WarehouseComponentType `json:"components,omitempty"`
}

type GetPagedWarehousesResponse struct {
	Result *PagedWarehousesDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetWarehouseStocksDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Warehouses []WarehouseOffersDTO `json:"warehouses"`
}

type GetWarehousesResponse struct {
	Result *WarehousesDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GoodsStatsWarehouseDTO struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Stocks []WarehouseStockDTO `json:"stocks"`
}

type PagedWarehousesDTO struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Warehouses []WarehouseDetailsDTO `json:"warehouses"`
}

type UpdateWarehouseStatusRequest struct {
	Enabled bool `json:"enabled"`
}

type UpdateWarehouseStatusResponse struct {
	Result *WarehouseStatusDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type WarehouseAddressDTO struct {
	Block *string `json:"block,omitempty"`

	Building *string `json:"building,omitempty"`

	City string `json:"city"`

	Gps GpsDTO `json:"gps"`

	Number *string `json:"number,omitempty"`

	Street *string `json:"street,omitempty"`
}

type WarehouseComponentType string

type WarehouseDTO struct {
	Address *WarehouseAddressDTO `json:"address,omitempty"`

	CampaignId int64 `json:"campaignId"`

	Express bool `json:"express"`

	Id int64 `json:"id"`

	Name string `json:"name"`
}

type WarehouseDeliveryOptionsDTO struct {
	CourierDelivery *CourierDeliveryOptionsDTO `json:"courierDelivery,omitempty"`

	PickupDelivery *PickupDeliveryOptionsDTO `json:"pickupDelivery,omitempty"`
}

type WarehouseDetailsDTO struct {
	Address *WarehouseAddressDTO `json:"address,omitempty"`

	CampaignId int64 `json:"campaignId"`

	Express bool `json:"express"`

	GroupInfo *WarehouseGroupInfoDTO `json:"groupInfo,omitempty"`

	Id int64 `json:"id"`

	Name string `json:"name"`

	Status *WarehouseStatusDTO `json:"status,omitempty"`
}

type WarehouseGroupDTO struct {
	MainWarehouse WarehouseDTO `json:"mainWarehouse"`

	Name string `json:"name"`

	Warehouses []WarehouseDTO `json:"warehouses"`
}

type WarehouseGroupInfoDTO struct {
	Id int64 `json:"id"`

	Name string `json:"name"`
}

type WarehouseStatusDTO struct {
	Status WarehouseStatusType `json:"status"`
}

type WarehouseStatusType string

type WarehouseStockDTO struct {
	Count int64 `json:"count"`

	Type WarehouseStockType `json:"type"`
}

type WarehouseStockType string

type WarehousesDTO struct {
	WarehouseGroups []WarehouseGroupDTO `json:"warehouseGroups"`

	Warehouses []WarehouseDTO `json:"warehouses"`
}

type WarehousesDeliveryOptionsDTO struct {
	DeliveryOptions WarehouseDeliveryOptionsDTO `json:"deliveryOptions"`

	Items []BasicOrderItemDTO `json:"items"`

	WarehouseId int64 `json:"warehouseId"`
}

type GetPagedWarehousesParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetFulfillmentWarehousesParams struct {
	CampaignId *CampaignId `form:"campaignId,omitempty" json:"campaignId,omitempty"`
}

type GetPagedWarehousesJSONRequestBody = GetPagedWarehousesRequest

type UpdateWarehouseStatusJSONRequestBody = UpdateWarehouseStatusRequest
