package models

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	ChatContextIdentifiableTypeORDER  ChatContextIdentifiableType = "ORDER"
	ChatContextIdentifiableTypeRETURN ChatContextIdentifiableType = "RETURN"
)

const (
	ChatContextTypeDIRECT ChatContextType = "DIRECT"
	ChatContextTypeORDER  ChatContextType = "ORDER"
	ChatContextTypeRETURN ChatContextType = "RETURN"
)

const (
	ChatMessageSenderTypeCUSTOMER ChatMessageSenderType = "CUSTOMER"
	ChatMessageSenderTypeMARKET   ChatMessageSenderType = "MARKET"
	ChatMessageSenderTypePARTNER  ChatMessageSenderType = "PARTNER"
	ChatMessageSenderTypeSUPPORT  ChatMessageSenderType = "SUPPORT"
)

const (
	ChatStatusTypeFINISHED           ChatStatusType = "FINISHED"
	ChatStatusTypeNEW                ChatStatusType = "NEW"
	ChatStatusTypeWAITINGFORARBITER  ChatStatusType = "WAITING_FOR_ARBITER"
	ChatStatusTypeWAITINGFORCUSTOMER ChatStatusType = "WAITING_FOR_CUSTOMER"
	ChatStatusTypeWAITINGFORMARKET   ChatStatusType = "WAITING_FOR_MARKET"
	ChatStatusTypeWAITINGFORPARTNER  ChatStatusType = "WAITING_FOR_PARTNER"
)

type ChatContextDTO struct {
	Id int64 `json:"id"`

	Type ChatContextIdentifiableType `json:"type"`
}

type ChatContextIdentifiableType string

type ChatContextType string

type ChatCustomerDTO struct {
	Name *string `json:"name,omitempty"`

	PublicId *string `json:"publicId,omitempty"`
}

type ChatFullContextDTO struct {
	CampaignId *int64 `json:"campaignId,omitempty"`

	Customer *ChatCustomerDTO `json:"customer,omitempty"`

	OrderId *int64 `json:"orderId,omitempty"`

	ReturnId *int64 `json:"returnId,omitempty"`

	Type ChatContextType `json:"type"`
}

type ChatMessageDTO struct {
	CreatedAt time.Time `json:"createdAt"`

	Message *string `json:"message,omitempty"`

	MessageId int64 `json:"messageId"`

	Payload *[]ChatMessagePayloadDTO `json:"payload,omitempty"`

	Sender ChatMessageSenderType `json:"sender"`
}

type ChatMessagePayloadDTO struct {
	Name string `json:"name"`

	Size int32  `json:"size"`
	Url  string `json:"url"`
}

type ChatMessageSenderType string

type ChatMessagesResultDTO struct {
	Context ChatFullContextDTO `json:"context"`

	Messages []ChatMessageDTO `json:"messages"`

	OrderId *int64 `json:"orderId,omitempty"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type ChatStatusType string

type ChatType string

type CreateChatRequest struct {
	Context *ChatContextDTO `json:"context,omitempty"`

	OrderId *int64 `json:"orderId,omitempty"`
}

type CreateChatResponse struct {
	Result *CreateChatResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type CreateChatResultDTO struct {
	ChatId int64 `json:"chatId"`
}

type GetChatHistoryRequest struct {
	MessageIdFrom *int64 `json:"messageIdFrom,omitempty"`
}

type GetChatHistoryResponse struct {
	Result *ChatMessagesResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetChatInfoDTO struct {
	ChatId int64 `json:"chatId"`

	Context ChatFullContextDTO `json:"context"`

	CreatedAt time.Time `json:"createdAt"`

	OrderId *int64 `json:"orderId,omitempty"`

	Status ChatStatusType `json:"status"`

	Type ChatType `json:"type"`

	UpdatedAt time.Time `json:"updatedAt"`
}

type GetChatMessageResponse struct {
	Result *ChatMessageDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetChatResponse struct {
	Result *GetChatInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetChatsInfoDTO struct {
	Chats []GetChatInfoDTO `json:"chats"`

	Paging *PackagingForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type GetChatsRequest struct {
	ContextTypes *[]ChatContextType `json:"contextTypes,omitempty"`

	Contexts *[]ChatContextDTO `json:"contexts,omitempty"`

	OrderIds *[]int64 `json:"orderIds,omitempty"`

	Statuses *[]ChatStatusType `json:"statuses,omitempty"`

	Types *[]ChatType `json:"types,omitempty"`
}

type GetChatsResponse struct {
	Result *GetChatsInfoDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type SendFileToChatRequest struct {
	File openapi_types.File `json:"file"`
}

type SendMessageToChatRequest struct {
	Message string `json:"message"`
}

type GetChatParams struct {
	ChatId int64 `url:"chatId" json:"chatId"`
}

type GetChatsParams struct {
	PageToken *string `url:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `url:"limit,omitempty" json:"limit,omitempty"`
}

type SendFileToChatParams struct {
	ChatId int64 `url:"chatId" json:"chatId"`
}

type GetChatHistoryParams struct {
	ChatId int64 `url:"chatId" json:"chatId"`

	PageToken *string `url:"pageToken,omitempty" json:"pageToken,omitempty"`

	Limit *int32 `url:"limit,omitempty" json:"limit,omitempty"`
}

type GetChatMessageParams struct {
	ChatId int64 `url:"chatId" json:"chatId"`

	MessageId int64 `url:"messageId" json:"messageId"`
}

type SendMessageToChatParams struct {
	ChatId int64 `url:"chatId" json:"chatId"`
}

type GetChatsJSONRequestBody = GetChatsRequest

type SendFileToChatMultipartRequestBody = SendFileToChatRequest

type GetChatHistoryJSONRequestBody = GetChatHistoryRequest

type SendMessageToChatJSONRequestBody = SendMessageToChatRequest

type CreateChatJSONRequestBody = CreateChatRequest
