package models

type BusinessDTO struct {
	Id *int64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type BusinessSettingsDTO struct {
	Currency *CurrencyType `json:"currency,omitempty"`

	OnlyDefaultPrice *bool `json:"onlyDefaultPrice,omitempty"`
}

type BusinessSubscriptionLevelType string

type BusinessTraitType string

type GetBusinessSettingsInfoDTO struct {
	Info *BusinessDTO `json:"info,omitempty"`

	Settings *BusinessSettingsDTO `json:"settings,omitempty"`

	SubscriptionLevel *BusinessSubscriptionLevelType `json:"subscriptionLevel,omitempty"`

	Traits *[]BusinessTraitType `json:"traits,omitempty"`
}

type GetBusinessSettingsResponse struct {
	Result *GetBusinessSettingsInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
