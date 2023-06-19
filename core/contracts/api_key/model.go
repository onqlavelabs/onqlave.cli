package api_key

import "time"

type ShortInfoApplication struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShortInfoCluster struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type CreatedBy struct {
	ID           string `json:"id"`
	FullName     string `json:"name"`
	EmailAddress string `json:"email_address"`
	Avatar       string `json:"avatar"`
}

type Insights struct {
	TotalKeys    int64 `json:"total_keys"`
	TotalActive  int64 `json:"total_active"`
	TotalDeleted int64 `json:"total_deleted"`
}

type SensitiveData struct {
	ID                    string     `json:"id"`
	AccessKey             string     `json:"access_key"`
	Status                string     `json:"status"`
	ClientKey             string     `json:"client_key"`
	ServerSigningKey      string     `json:"server_signing_key"`
	ServerCryptoAccessKey string     `json:"server_crypto_access_key"`
	ArxUrl                string     `json:"arx_url"`
	ProvidedAt            *time.Time `json:"provided_at"`
}
