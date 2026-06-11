package models

import "time"

type FullOutletLicenseDTO struct {
	CheckComment *string `json:"checkComment,omitempty"`

	CheckStatus *LicenseCheckStatusType `json:"checkStatus,omitempty"`

	DateOfExpiry time.Time `json:"dateOfExpiry"`

	DateOfIssue time.Time `json:"dateOfIssue"`

	Id *int64 `json:"id,omitempty"`

	LicenseType LicenseType `json:"licenseType"`

	Number string `json:"number"`

	OutletId int64 `json:"outletId"`
}

type GetOutletLicensesResponse struct {
	Result *OutletLicensesResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type OutletLicenseDTO struct {
	DateOfExpiry time.Time `json:"dateOfExpiry"`

	DateOfIssue time.Time `json:"dateOfIssue"`

	Id *int64 `json:"id,omitempty"`

	LicenseType LicenseType `json:"licenseType"`

	Number string `json:"number"`

	OutletId int64 `json:"outletId"`
}

type OutletLicensesResponseDTO struct {
	Licenses []FullOutletLicenseDTO `json:"licenses"`
}

type UpdateOutletLicenseRequest struct {
	Licenses []OutletLicenseDTO `json:"licenses"`
}

type DeleteOutletLicensesParams struct {
	Ids []int64 `form:"ids" json:"ids"`
}

type GetOutletLicensesParams struct {
	OutletIds *[]int64 `form:"outletIds,omitempty" json:"outletIds,omitempty"`

	Ids *[]int64 `form:"ids,omitempty" json:"ids,omitempty"`
}

type UpdateOutletLicensesJSONRequestBody = UpdateOutletLicenseRequest
