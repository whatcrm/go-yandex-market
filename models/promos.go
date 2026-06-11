package models

import "time"

type DeletePromoOffersRequest struct {
	DeleteAllOffers *bool `json:"deleteAllOffers,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	PromoId string `json:"promoId"`
}

type DeletePromoOffersResponse struct {
	Result *DeletePromoOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPromoAssortmentInfoDTO struct {
	ActiveOffers int32 `json:"activeOffers"`

	PotentialOffers *int32 `json:"potentialOffers,omitempty"`

	Processing *bool `json:"processing,omitempty"`
}

type GetPromoBestsellerInfoDTO struct {
	Bestseller bool `json:"bestseller"`

	EntryDeadline *time.Time `json:"entryDeadline,omitempty"`

	RenewalEnabled *bool `json:"renewalEnabled,omitempty"`
}

type GetPromoConstraintsDTO struct {
	WarehouseIds *[]int64 `json:"warehouseIds,omitempty"`
}

type GetPromoDTO struct {
	AssortmentInfo GetPromoAssortmentInfoDTO `json:"assortmentInfo"`

	BestsellerInfo GetPromoBestsellerInfoDTO `json:"bestsellerInfo"`

	Channels *[]ChannelType `json:"channels,omitempty"`

	Constraints *GetPromoConstraintsDTO `json:"constraints,omitempty"`

	Id string `json:"id"`

	MechanicsInfo GetPromoMechanicsInfoDTO `json:"mechanicsInfo"`

	Name string `json:"name"`

	Participating bool `json:"participating"`

	Period PromoPeriodDTO `json:"period"`
}

type GetPromoMechanicsInfoDTO struct {
	PromocodeInfo *GetPromoPromocodeInfoDTO `json:"promocodeInfo,omitempty"`

	Type MechanicsType `json:"type"`
}

type GetPromoOffersRequest struct {
	PromoId string `json:"promoId"`

	StatusType *PromoOfferParticipationStatusFilterType `json:"statusType,omitempty"`

	Statuses *[]PromoOfferParticipationStatusMultiFilterType `json:"statuses,omitempty"`
}

type GetPromoOffersResponse struct {
	Result *GetPromoOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPromoPromocodeInfoDTO struct {
	Discount int32 `json:"discount"`

	Promocode string `json:"promocode"`
}

type GetPromosRequest struct {
	Mechanics *MechanicsType `json:"mechanics,omitempty"`

	Participation *PromoParticipationType `json:"participation,omitempty"`
}

type GetPromosResponse struct {
	Result *GetPromosResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPromosResultDTO struct {
	Promos []GetPromoDTO `json:"promos"`
}

type PromoParticipationType string

type PromoPeriodDTO struct {
	DateTimeFrom time.Time `json:"dateTimeFrom"`

	DateTimeTo time.Time `json:"dateTimeTo"`
}

type UpdatePromoOffersRequest struct {
	Offers []UpdatePromoOfferDTO `json:"offers"`

	PromoId string `json:"promoId"`
}

type UpdatePromoOffersResponse struct {
	Result *UpdatePromoOffersResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetPromosJSONRequestBody = GetPromosRequest
