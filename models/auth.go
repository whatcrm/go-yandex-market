package models

const (
	ApiKeyScopes apiKeyContextKey = "ApiKey.Scopes"
	OAuthScopes  oAuthContextKey  = "OAuth.Scopes"
)

type ApiKeyDTO struct {
	AuthScopes []ApiKeyScopeType `json:"authScopes"`

	Name string `json:"name"`
}

type ApiKeyScopeType string

type GetTokenInfoResponse struct {
	Result *TokenDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type TokenDTO struct {
	ApiKey ApiKeyDTO `json:"apiKey"`
}
