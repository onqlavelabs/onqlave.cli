package common

type OnqlaveOperationContext interface {
	SetApiKey(apiKey OnqlaveApiKeyHeader)
	GetApiKey() OnqlaveApiKeyHeader
	GetUser() *Token
	SetUser(token *Token)
}
