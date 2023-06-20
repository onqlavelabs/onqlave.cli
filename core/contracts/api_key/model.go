package api_key

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	arx "github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type APIKey struct {
	ID            string                   `json:"id"`
	AccessKey     string                   `json:"access_key"`
	CreatedAt     string                   `json:"created_at"`
	Status        string                   `json:"status"`
	CreatedBy     *common.ShortUserInfo    `json:"created_by"`
	ApplicationID string                   `json:"application_id,omitempty"`
	Application   common.ShortResourceInfo `json:"application"`
	ArxID         string                   `json:"cluster_id,omitempty"`
	Arx           common.ShortResourceInfo `json:"cluster"`
	ACL           acl.ACL                  `json:"acl"`
	ArxUrl        string                   `json:"arx_url"`
}

type APIKeys struct {
	ACL      acl.ACL  `json:"acl"`
	APIKeys  []APIKey `json:"api_keys"`
	Model    Models   `json:"model"`
	Insights Insights `json:"insights"`
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

type Models struct {
	Applications []Application `json:"applications"`
	Arx          []Arx         `json:"clusters"`
}

type Application struct {
	common.ShortResourceInfo `json:",inline"`
	Label                    string                 `json:"label"`
	ApplicationTechnology    application.Technology `json:"application_technology"`
}

type Arx struct {
	common.ShortResourceInfo `json:",inline"`
	Purpose                  arx.Purpose                 `json:"purpose"`
	Plan                     arx.Plan                    `json:"plan"`
	Provider                 arx.Provider                `json:"provider"`
	Regions                  []arx.Region                `json:"regions"`
	Encryption               arx.EncryptionMethod        `json:"encryption"`
	RotationCycle            arx.EncryptionRotationCycle `json:"rotation_cycle"`
	CreatedBy                *common.ShortUserInfo       `json:"created_by"`
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

type APIKeyVersion struct {
	Number    string    `json:"version_number" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
}
