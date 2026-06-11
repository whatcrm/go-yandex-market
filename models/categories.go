package models

const (
	CategoryErrorTypeCATEGORYISNOTLEAF CategoryErrorType = "CATEGORY_IS_NOT_LEAF"
	CategoryErrorTypeUNKNOWNCATEGORY   CategoryErrorType = "UNKNOWN_CATEGORY"
)

type ApiClientDataErrorResponse = ApiErrorResponse

type ApiForbiddenErrorResponse = ApiErrorResponse

type ApiLimitErrorResponse = ApiErrorResponse

type ApiNotFoundErrorResponse = ApiErrorResponse

type ApiServerErrorResponse = ApiErrorResponse

type ApiUnauthorizedErrorResponse = ApiErrorResponse

type CategoryContentParametersDTO struct {
	CategoryId int32 `json:"categoryId"`

	Parameters *[]CategoryParameterDTO `json:"parameters,omitempty"`
}

type CategoryDTO struct {
	Children *[]CategoryDTO `json:"children,omitempty"`

	Id int64 `json:"id"`

	Name string `json:"name"`
}

type CategoryErrorDTO struct {
	CategoryId *int64 `json:"categoryId,omitempty"`

	Type *CategoryErrorType `json:"type,omitempty"`
}

type CategoryErrorType string

type CategoryId = int32

type CategoryParameterDTO struct {
	AllowCustomValues bool `json:"allowCustomValues"`

	Constraints *ParameterValueConstraintsDTO `json:"constraints,omitempty"`

	Description *string `json:"description,omitempty"`

	Distinctive bool `json:"distinctive"`

	Filtering bool `json:"filtering"`

	Id int64 `json:"id"`

	Multivalue bool `json:"multivalue"`

	Name *string `json:"name,omitempty"`

	RecommendationTypes *[]OfferCardRecommendationType `json:"recommendationTypes,omitempty"`

	Required bool `json:"required"`

	Type ParameterType `json:"type"`

	Unit *CategoryParameterUnitDTO `json:"unit,omitempty"`

	ValueRestrictions *[]ValueRestrictionDTO `json:"valueRestrictions,omitempty"`

	Values *[]ParameterValueOptionDTO `json:"values,omitempty"`
}

type CategoryParameterUnitDTO struct {
	DefaultUnitId int64 `json:"defaultUnitId"`

	Units []UnitDTO `json:"units"`
}

type GetCategoriesMaxSaleQuantumRequest struct {
	MarketCategoryIds []int64 `json:"marketCategoryIds"`
}

type GetCategoriesMaxSaleQuantumResponse struct {
	Errors *[]CategoryErrorDTO `json:"errors,omitempty"`

	Results []MaxSaleQuantumDTO `json:"results"`

	Status ApiResponseStatusType `json:"status"`
}

type GetCategoriesRequest struct {
	Language *LanguageType `json:"language,omitempty"`
}

type GetCategoriesResponse struct {
	Result *CategoryDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetCategoryContentParametersParams struct {
	BusinessId *int64 `form:"businessId,omitempty" json:"businessId,omitempty"`
}
