package chats

const (
	CreateChatEndpoint        = "/v2/businesses/{businessId}/chats/new"
	GetChatEndpoint           = "/v2/businesses/{businessId}/chat"
	GetChatHistoryEndpoint    = "/v2/businesses/{businessId}/chats/history"
	GetChatMessageEndpoint    = "/v2/businesses/{businessId}/chats/message"
	GetChatsEndpoint          = "/v2/businesses/{businessId}/chats"
	SendFileToChatEndpoint    = "/v2/businesses/{businessId}/chats/file/send"
	SendMessageToChatEndpoint = "/v2/businesses/{businessId}/chats/message"
)
