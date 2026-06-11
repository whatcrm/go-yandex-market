package models

const (
	QuestionsTextContentAuthorTypeBRAND    QuestionsTextContentAuthorType = "BRAND"
	QuestionsTextContentAuthorTypeBUSINESS QuestionsTextContentAuthorType = "BUSINESS"
	QuestionsTextContentAuthorTypeUSER     QuestionsTextContentAuthorType = "USER"
	QuestionsTextContentAuthorTypeVENDOR   QuestionsTextContentAuthorType = "VENDOR"
)

const (
	QuestionsTextContentModerationStatusTypeBANNED      QuestionsTextContentModerationStatusType = "BANNED"
	QuestionsTextContentModerationStatusTypeDELETED     QuestionsTextContentModerationStatusType = "DELETED"
	QuestionsTextContentModerationStatusTypePUBLISHED   QuestionsTextContentModerationStatusType = "PUBLISHED"
	QuestionsTextContentModerationStatusTypeUNMODERATED QuestionsTextContentModerationStatusType = "UNMODERATED"
)

type GetCategoryContentParametersResponse struct {
	Result *CategoryContentParametersDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOfferCardsContentStatusRequest struct {
	CardStatuses *[]OfferCardStatusType `json:"cardStatuses,omitempty"`

	CategoryIds *[]CategoryId `json:"categoryIds,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	WithRecommendations *bool `json:"withRecommendations,omitempty"`
}

type GetOfferCardsContentStatusResponse struct {
	Result *OfferCardsContentStatusDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type QuestionsTextContentAuthorDTO struct {
	Name *string `json:"name,omitempty"`

	Type *QuestionsTextContentAuthorType `json:"type,omitempty"`
}

type QuestionsTextContentAuthorType string

type QuestionsTextContentModerationStatusType string

type UpdateOfferContentRequest struct {
	OffersContent []OfferContentDTO `json:"offersContent"`
}

type UpdateOfferContentResponse struct {
	Results *[]UpdateOfferContentResultDTO `json:"results,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}
