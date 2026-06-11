package models

const (
	OutletStatusTypeATMODERATION OutletStatusType = "AT_MODERATION"
	OutletStatusTypeFAILED       OutletStatusType = "FAILED"
	OutletStatusTypeMODERATED    OutletStatusType = "MODERATED"
	OutletStatusTypeNONMODERATED OutletStatusType = "NONMODERATED"
	OutletStatusTypeUNKNOWN      OutletStatusType = "UNKNOWN"
)

const (
	OutletVisibilityTypeHIDDEN  OutletVisibilityType = "HIDDEN"
	OutletVisibilityTypeUNKNOWN OutletVisibilityType = "UNKNOWN"
	OutletVisibilityTypeVISIBLE OutletVisibilityType = "VISIBLE"
)

type ChangeOutletRequest = OutletDTO

type CreateOutletResponse struct {
	Result *OutletResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type FullOutletDTO struct {
	Address OutletAddressDTO `json:"address"`

	Coords *string `json:"coords,omitempty"`

	DeliveryRules *[]OutletDeliveryRuleDTO `json:"deliveryRules,omitempty"`

	Id *int64 `json:"id,omitempty"`

	IsMain *bool `json:"isMain,omitempty"`

	ModerationReason *string `json:"moderationReason,omitempty"`

	Name string `json:"name"`

	Phones []string `json:"phones"`

	Region *RegionDTO `json:"region,omitempty"`

	ShopOutletCode *string `json:"shopOutletCode,omitempty"`

	ShopOutletId *string `json:"shopOutletId,omitempty"`

	Status *OutletStatusType `json:"status,omitempty"`

	StoragePeriod *int64 `json:"storagePeriod,omitempty"`

	Type OutletType `json:"type"`

	Visibility *OutletVisibilityType `json:"visibility,omitempty"`

	WorkingSchedule OutletWorkingScheduleDTO `json:"workingSchedule"`

	WorkingTime *string `json:"workingTime,omitempty"`
}

type GetOutletResponse struct {
	Outlet *FullOutletDTO `json:"outlet,omitempty"`
}

type GetOutletsResponse struct {
	Outlets []FullOutletDTO `json:"outlets"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type OutletAddressDTO struct {
	Additional *string `json:"additional,omitempty"`

	Block *string `json:"block,omitempty"`

	Building *string `json:"building,omitempty"`

	City *string `json:"city,omitempty"`

	Estate *string `json:"estate,omitempty"`

	Km *int32 `json:"km,omitempty"`

	Number *string `json:"number,omitempty"`

	RegionId int64 `json:"regionId"`

	Street *string `json:"street,omitempty"`
}

type OutletDTO struct {
	Address OutletAddressDTO `json:"address"`

	Coords *string `json:"coords,omitempty"`

	DeliveryRules *[]OutletDeliveryRuleDTO `json:"deliveryRules,omitempty"`

	IsMain *bool `json:"isMain,omitempty"`

	Name string `json:"name"`

	Phones []string `json:"phones"`

	ShopOutletCode *string `json:"shopOutletCode,omitempty"`

	StoragePeriod *int64 `json:"storagePeriod,omitempty"`

	Type OutletType `json:"type"`

	Visibility *OutletVisibilityType `json:"visibility,omitempty"`

	WorkingSchedule OutletWorkingScheduleDTO `json:"workingSchedule"`
}

type OutletDeliveryRuleDTO struct {
	DeliveryServiceId *int64 `json:"deliveryServiceId,omitempty"`

	MaxDeliveryDays *int32 `json:"maxDeliveryDays,omitempty"`

	MinDeliveryDays *int32 `json:"minDeliveryDays,omitempty"`

	OrderBefore *int32 `json:"orderBefore,omitempty"`

	PriceFreePickup *float32 `json:"priceFreePickup,omitempty"`

	UnspecifiedDeliveryInterval *bool `json:"unspecifiedDeliveryInterval,omitempty"`
}

type OutletResponseDTO struct {
	Id *int64 `json:"id,omitempty"`
}

type OutletStatusType string

type OutletType string

type OutletVisibilityType string

type OutletWorkingScheduleDTO struct {
	ScheduleItems []OutletWorkingScheduleItemDTO `json:"scheduleItems"`

	WorkInHoliday *bool `json:"workInHoliday,omitempty"`
}

type OutletWorkingScheduleItemDTO struct {
	EndDay DayOfWeekType `json:"endDay"`

	EndTime string `json:"endTime"`

	StartDay DayOfWeekType `json:"startDay"`

	StartTime string `json:"startTime"`
}

type GetOutletsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	RegionId *int64 `form:"region_id,omitempty" json:"region_id,omitempty"`

	ShopOutletCode *string `form:"shop_outlet_code,omitempty" json:"shop_outlet_code,omitempty"`

	RegionIdLegacy *int64 `form:"regionId,omitempty" json:"regionId,omitempty"`
}

type CreateOutletJSONRequestBody = ChangeOutletRequest

type UpdateOutletJSONRequestBody = ChangeOutletRequest
