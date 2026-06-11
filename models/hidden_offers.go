package models

type AddHiddenOffersRequest struct {
	HiddenOffers []HiddenOfferDTO `json:"hiddenOffers"`
}

type DeleteHiddenOffersRequest struct {
	HiddenOffers []HiddenOfferDTO `json:"hiddenOffers"`
}

type GetHiddenOffersResponse struct {
	Result *GetHiddenOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
