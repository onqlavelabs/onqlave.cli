package api_key

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	application "github.com/onqlavelabs/onqlave.cli/core/contracts/application"
	arx "github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type APIKeyDetail struct {
	ID            string                `json:"id"`
	AccessKey     string                `json:"access_key"`
	CreatedAt     string                `json:"created_at"`
	Status        string                `json:"status"`
	CreatedBy     *common.ShortUserInfo `json:"created_by"`
	ApplicationID string                `json:"application_id,omitempty"`
	Application   ShortInfoApplication  `json:"application"`
	ClusterID     string                `json:"cluster_id,omitempty"`
	Cluster       ShortInfoCluster      `json:"cluster"`
	ACL           acl.ACL               `json:"acl"`
	ArxUrl        string                `json:"arx_url"`
}

type Application struct {
	ShortInfoApplication  `json:",inline"`
	Label                 string                            `json:"label"`
	ApplicationTechnology application.ApplicationTechnology `json:"application_technology"`
}

type Cluster struct {
	ShortInfoCluster `json:",inline"`
	Purpose          arx.ArxPurpose                 `json:"purpose"`
	Plan             arx.ArxPlan                    `json:"plan"`
	Provider         arx.ArxProvider                `json:"provider"`
	Regions          []arx.ArxRegion                `json:"regions"`
	Encryption       arx.ArxEncryptionMethod        `json:"encryption"`
	RotationCycle    arx.ArxEncryptionRotationCycle `json:"rotation_cycle"`
	CreatedBy        CreatedBy                      `json:"created_by"`
}

type APIKeyModelsWrapper struct {
	Applications []Application `json:"applications"`
	Clusters     []Cluster     `json:"clusters"`
}

type GetAPIKeysResponseWrapper struct {
	ACL      acl.ACL             `json:"acl"`
	APIKeys  []APIKeyDetail      `json:"api_keys"`
	Model    APIKeyModelsWrapper `json:"model"`
	Insights Insights            `json:"insights"`
}
