package api_key

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	application "github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	arx "github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type APIKey struct {
	ID            string                `json:"id"`
	AccessKey     string                `json:"access_key"`
	CreatedAt     string                `json:"created_at"`
	Status        string                `json:"status"`
	CreatedBy     *common.ShortUserInfo `json:"created_by"`
	ApplicationID string                `json:"application_id,omitempty"`
	Application   ShortInfoApplication  `json:"application"`
	ArxID         string                `json:"cluster_id,omitempty"`
	Arx           ShortInfoArx          `json:"cluster"`
	ACL           acl.ACL               `json:"acl"`
	ArxUrl        string                `json:"arx_url"`
}

type APIKeys struct {
	ACL      acl.ACL      `json:"acl"`
	APIKeys  []APIKey     `json:"api_keys"`
	Model    APIKeyModels `json:"model"`
	Insights Insights     `json:"insights"`
}

type APIKeySensitiveData struct {
	ID                    string     `json:"id"`
	AccessKey             string     `json:"access_key"`
	Status                string     `json:"status"`
	ClientKey             string     `json:"client_key"`
	ServerSigningKey      string     `json:"server_signing_key"`
	ServerCryptoAccessKey string     `json:"server_crypto_access_key"`
	ArxUrl                string     `json:"arx_url"`
	ProvidedAt            *time.Time `json:"provided_at"`
}

type APIKeyModels struct {
	Applications []APIKeyApplication `json:"applications"`
	Arx          []APIKeyArx         `json:"clusters"`
}

type APIKeyApplication struct {
	ShortInfoApplication  `json:",inline"`
	Label                 string                            `json:"label"`
	ApplicationTechnology application.ApplicationTechnology `json:"application_technology"`
}

type APIKeyArx struct {
	ShortInfoArx  `json:",inline"`
	Purpose       arx.ArxPurpose                 `json:"purpose"`
	Plan          arx.ArxPlan                    `json:"plan"`
	Provider      arx.ArxProvider                `json:"provider"`
	Regions       []arx.ArxRegion                `json:"regions"`
	Encryption    arx.ArxEncryptionMethod        `json:"encryption"`
	RotationCycle arx.ArxEncryptionRotationCycle `json:"rotation_cycle"`
	CreatedBy     *common.ShortUserInfo          `json:"created_by"`
}

type ShortInfoApplication struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShortInfoArx struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}

type Insights struct {
	TotalKeys    int64 `json:"total_keys"`
	TotalActive  int64 `json:"total_active"`
	TotalDeleted int64 `json:"total_deleted"`
}

type CreateAPIKey struct {
	ApplicationID         string `json:"application_id" validate:"required"`
	ClusterID             string `json:"cluster_id" validate:"required"`
	ApplicationTechnology string `json:"application_technology" validate:"required"`
}
