package models

const (
	CalculatedTariffTypeAGENCYCOMMISSION      CalculatedTariffType = "AGENCY_COMMISSION"
	CalculatedTariffTypeCROSSREGIONALDELIVERY CalculatedTariffType = "CROSSREGIONAL_DELIVERY"
	CalculatedTariffTypeDELIVERYTOCUSTOMER    CalculatedTariffType = "DELIVERY_TO_CUSTOMER"
	CalculatedTariffTypeEXPRESSDELIVERY       CalculatedTariffType = "EXPRESS_DELIVERY"
	CalculatedTariffTypeFEE                   CalculatedTariffType = "FEE"
	CalculatedTariffTypeMIDDLEMILE            CalculatedTariffType = "MIDDLE_MILE"
	CalculatedTariffTypePAYMENTTRANSFER       CalculatedTariffType = "PAYMENT_TRANSFER"
	CalculatedTariffTypeSORTING               CalculatedTariffType = "SORTING"
)

type CalculateTariffsParametersDTO struct {
	CampaignId *int64 `json:"campaignId,omitempty"`

	Currency *CurrencyType `json:"currency,omitempty"`

	Frequency *PaymentFrequencyType `json:"frequency,omitempty"`

	PaymentDelayWeeks *int32 `json:"paymentDelayWeeks,omitempty"`

	SellingProgram *SellingProgramType `json:"sellingProgram,omitempty"`
}

type CalculateTariffsRequest struct {
	Offers []CalculateTariffsOfferDTO `json:"offers"`

	Parameters CalculateTariffsParametersDTO `json:"parameters"`
}

type CalculateTariffsResponse struct {
	Result *CalculateTariffsResponseDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type CalculateTariffsResponseDTO struct {
	Offers []CalculateTariffsOfferInfoDTO `json:"offers"`
}

type CalculatedTariffDTO struct {
	Amount *float32 `json:"amount,omitempty"`

	Currency *CurrencyType `json:"currency,omitempty"`

	Parameters []TariffParameterDTO `json:"parameters"`

	Type CalculatedTariffType `json:"type"`
}

type CalculatedTariffType string

type TariffDTO struct {
	Amount float32 `json:"amount"`

	Currency CurrencyType `json:"currency"`

	Parameters []TariffParameterDTO `json:"parameters"`

	Percent *float32 `json:"percent,omitempty"`

	Type TariffType `json:"type"`
}

type TariffParameterDTO struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

type TariffType string

type CalculateTariffsJSONRequestBody = CalculateTariffsRequest
