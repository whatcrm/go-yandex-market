package models

import "time"

const (
	GoodsFeedbackCommentAuthorTypeBUSINESS GoodsFeedbackCommentAuthorType = "BUSINESS"
	GoodsFeedbackCommentAuthorTypeUSER     GoodsFeedbackCommentAuthorType = "USER"
)

const (
	GoodsFeedbackCommentStatusTypeBANNED      GoodsFeedbackCommentStatusType = "BANNED"
	GoodsFeedbackCommentStatusTypeDELETED     GoodsFeedbackCommentStatusType = "DELETED"
	GoodsFeedbackCommentStatusTypePUBLISHED   GoodsFeedbackCommentStatusType = "PUBLISHED"
	GoodsFeedbackCommentStatusTypeUNMODERATED GoodsFeedbackCommentStatusType = "UNMODERATED"
)

type DeleteGoodsFeedbackCommentRequest struct {
	Id int64 `json:"id"`
}

type GetGoodsFeedbackCommentsRequest struct {
	CommentIds *[]GoodsFeedbackCommentId `json:"commentIds,omitempty"`

	FeedbackId *int64 `json:"feedbackId,omitempty"`
}

type GetGoodsFeedbackCommentsResponse struct {
	Result *GoodsFeedbackCommentListDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetGoodsFeedbackRequest struct {
	DateTimeFrom *time.Time `json:"dateTimeFrom,omitempty"`

	DateTimeTo *time.Time `json:"dateTimeTo,omitempty"`

	FeedbackIds *[]GoodsFeedbackId `json:"feedbackIds,omitempty"`

	OfferIds *[]ShopSku `json:"offerIds,omitempty"`

	Paid *bool `json:"paid,omitempty"`

	RatingValues *[]int32 `json:"ratingValues,omitempty"`

	ReactionStatus *FeedbackReactionStatusType `json:"reactionStatus,omitempty"`
}

type GetGoodsFeedbackResponse struct {
	Result *GoodsFeedbackListDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GoodsFeedbackCommentAuthorDTO struct {
	Name *string `json:"name,omitempty"`

	Type *GoodsFeedbackCommentAuthorType `json:"type,omitempty"`
}

type GoodsFeedbackCommentAuthorType string

type GoodsFeedbackCommentDTO struct {
	Author *GoodsFeedbackCommentAuthorDTO `json:"author,omitempty"`

	CanModify *bool `json:"canModify,omitempty"`

	FeedbackId int64 `json:"feedbackId"`

	Id int64 `json:"id"`

	ParentId *int64 `json:"parentId,omitempty"`

	Status GoodsFeedbackCommentStatusType `json:"status"`

	Text string `json:"text"`
}

type GoodsFeedbackCommentId = int64

type GoodsFeedbackCommentListDTO struct {
	Comments []GoodsFeedbackCommentDTO `json:"comments"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GoodsFeedbackCommentStatusType string

type GoodsFeedbackDTO struct {
	Author *string `json:"author,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	Description *GoodsFeedbackDescriptionDTO `json:"description,omitempty"`

	FeedbackId int64 `json:"feedbackId"`

	Identifiers GoodsFeedbackIdentifiersDTO `json:"identifiers"`

	Media *GoodsFeedbackMediaDTO `json:"media,omitempty"`

	NeedReaction bool `json:"needReaction"`

	Statistics GoodsFeedbackStatisticsDTO `json:"statistics"`
}

type GoodsFeedbackDescriptionDTO struct {
	Advantages *string `json:"advantages,omitempty"`

	Comment *string `json:"comment,omitempty"`

	Disadvantages *string `json:"disadvantages,omitempty"`
}

type GoodsFeedbackId = int64

type GoodsFeedbackIdentifiersDTO struct {
	OfferId *string `json:"offerId,omitempty"`

	OrderId *int64 `json:"orderId,omitempty"`
}

type GoodsFeedbackListDTO struct {
	Feedbacks []GoodsFeedbackDTO `json:"feedbacks"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GoodsFeedbackMediaDTO struct {
	Photos *[]string `json:"photos,omitempty"`

	Videos *[]string `json:"videos,omitempty"`
}

type GoodsFeedbackStatisticsDTO struct {
	CommentsCount int64 `json:"commentsCount"`

	PaidAmount *int64 `json:"paidAmount,omitempty"`

	Rating int32 `json:"rating"`

	Recommended *bool `json:"recommended,omitempty"`
}

type SkipGoodsFeedbackReactionRequest struct {
	FeedbackIds []GoodsFeedbackId `json:"feedbackIds"`
}

type UpdateGoodsFeedbackCommentDTO struct {
	Id *int64 `json:"id,omitempty"`

	ParentId *int64 `json:"parentId,omitempty"`

	Text string `json:"text"`
}

type UpdateGoodsFeedbackCommentRequest struct {
	Comment UpdateGoodsFeedbackCommentDTO `json:"comment"`

	FeedbackId int64 `json:"feedbackId"`
}

type UpdateGoodsFeedbackCommentResponse struct {
	Result *GoodsFeedbackCommentDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetGoodsFeedbacksParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetGoodsFeedbackCommentsParams struct {
	PageToken *string `form:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`
}

type GetGoodsFeedbacksJSONRequestBody = GetGoodsFeedbackRequest

type GetGoodsFeedbackCommentsJSONRequestBody = GetGoodsFeedbackCommentsRequest

type DeleteGoodsFeedbackCommentJSONRequestBody = DeleteGoodsFeedbackCommentRequest

type UpdateGoodsFeedbackCommentJSONRequestBody = UpdateGoodsFeedbackCommentRequest

type SkipGoodsFeedbacksReactionJSONRequestBody = SkipGoodsFeedbackReactionRequest
