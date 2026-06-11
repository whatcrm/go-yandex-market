package models

import openapi_types "github.com/oapi-codegen/runtime/types"

type GetAnswersRequest struct {
	AnswerIds *[]AnswerId `json:"answerIds,omitempty"`

	QuestionId *int64 `json:"questionId,omitempty"`
}

type GetAnswersResponse struct {
	Result *AnswerListDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetQuestionsRequest struct {
	CategoryIds *[]int64 `json:"categoryIds,omitempty"`

	DateFrom *openapi_types.Date `json:"dateFrom,omitempty"`

	DateTo *openapi_types.Date `json:"dateTo,omitempty"`

	NeedAnswer *bool `json:"needAnswer,omitempty"`

	QuestionIds *[]QuestionId `json:"questionIds,omitempty"`

	Sort *QuestionSortOrderType `json:"sort,omitempty"`
}

type GetQuestionsResponse struct {
	Result *QuestionListDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type UpdateGoodsQuestionTextEntityDTO struct {
	Entity TypedQuestionsTextEntityIdDTO `json:"entity"`
}

type UpdateGoodsQuestionTextEntityRequest struct {
	EntityId *TypedQuestionsTextEntityIdDTO `json:"entityId,omitempty"`

	OperationType QuestionsTextEntityOperationType `json:"operationType"`

	ParentEntityId *TypedQuestionsTextEntityIdDTO `json:"parentEntityId,omitempty"`

	Text *string `json:"text,omitempty"`
}

type UpdateGoodsQuestionTextEntityResponse struct {
	Result *UpdateGoodsQuestionTextEntityDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetGoodsQuestionsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetGoodsQuestionAnswersParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetGoodsQuestionsJSONRequestBody = GetQuestionsRequest

type GetGoodsQuestionAnswersJSONRequestBody = GetAnswersRequest

type UpdateGoodsQuestionTextEntityJSONRequestBody = UpdateGoodsQuestionTextEntityRequest
