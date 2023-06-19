package common

type OnqlaveOperationContext interface {
	SetApiKey(apiKey OnqlaveApiKeyHeader)
	GetApiKey() OnqlaveApiKeyHeader
	GetUser() *Token
	SetUser(token *Token)
}

type onqlaveOperationContext struct {
	apiKey OnqlaveApiKeyHeader
	user   *Token
}

func (c *onqlaveOperationContext) SetApiKey(apiKey OnqlaveApiKeyHeader) {
	c.apiKey = apiKey
}

func (c *onqlaveOperationContext) GetApiKey() OnqlaveApiKeyHeader {
	return c.apiKey
}

func (c *onqlaveOperationContext) GetUser() *Token {
	return c.user
}

func (c *onqlaveOperationContext) SetUser(token *Token) {
	c.user = token
}
