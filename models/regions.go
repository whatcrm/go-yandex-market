package models

const (
	RegionTypeCITY            RegionType = "CITY"
	RegionTypeCITYDISTRICT    RegionType = "CITY_DISTRICT"
	RegionTypeCONTINENT       RegionType = "CONTINENT"
	RegionTypeCOUNTRY         RegionType = "COUNTRY"
	RegionTypeCOUNTRYDISTRICT RegionType = "COUNTRY_DISTRICT"
	RegionTypeOTHER           RegionType = "OTHER"
	RegionTypeREGION          RegionType = "REGION"
	RegionTypeREPUBLIC        RegionType = "REPUBLIC"
	RegionTypeREPUBLICAREA    RegionType = "REPUBLIC_AREA"
	RegionTypeSUBWAYSTATION   RegionType = "SUBWAY_STATION"
	RegionTypeVILLAGE         RegionType = "VILLAGE"
)

type GetRegionByIdResponse struct {
	Region RegionDTO `json:"region"`

	Regions *[]RegionDTO `json:"regions,omitempty"`
}

type GetRegionWithChildrenResponse struct {
	Pager *FlippingPagerDTO `json:"pager,omitempty"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Regions *RegionWithChildrenDTO `json:"regions,omitempty"`
}

type GetRegionsCodesResponse struct {
	Countries []CountryDTO `json:"countries"`
}

type GetRegionsResponse struct {
	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`

	Regions []RegionDTO `json:"regions"`
}

type RegionDTO struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Parent *RegionDTO `json:"parent,omitempty"`

	Type RegionType `json:"type"`
}

type RegionType string

type RegionWithChildrenDTO struct {
	Children *[]RegionDTO `json:"children,omitempty"`

	Id int64 `json:"id"`

	Name string `json:"name"`

	Parent *RegionDTO `json:"parent,omitempty"`

	Type RegionType `json:"type"`
}

type SearchRegionsByNameParams struct {
	Name string `form:"name" json:"name"`

	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type SearchRegionChildrenParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	Page *int32 `form:"page,omitempty" json:"page,omitempty"`

	PageSize *int32 `form:"pageSize,omitempty" json:"pageSize,omitempty"`
}
