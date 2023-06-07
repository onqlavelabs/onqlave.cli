package contracts

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/acl/contracts"

type GetAPIKeysResponseWrapper struct {
	ACL      contracts.ACL       `json:"acl"`
	APIKeys  []APIKey            `json:"api_keys"`
	Model    APIKeyModelsWrapper `json:"model"`
	Insights Insights            `json:"insights"`
}

type GetAPIKeyBaseResponse struct {
	ACL   contracts.ACL       `json:"acl"`
	Model APIKeyModelsWrapper `json:"model"`
}

func (wrapper *GetAPIKeysResponseWrapper) SetACL(acl contracts.ACL) {
	wrapper.ACL = acl
}

func (wrapper *GetAPIKeysResponseWrapper) SetAPIKey(apiKey APIKey) {
	wrapper.APIKeys = append(wrapper.APIKeys, apiKey)
}

func (wrapper *GetAPIKeysResponseWrapper) SetModel(apps []Application, clusters []Cluster) {
	if apps != nil {
		wrapper.Model.Applications = make([]Application, len(apps))
		copy(wrapper.Model.Applications, apps)
	}

	if clusters != nil {
		wrapper.Model.Clusters = make([]Cluster, len(clusters))
		copy(wrapper.Model.Clusters, clusters)
	}
}

func (wrapper *GetAPIKeysResponseWrapper) SetInsights(insights Insights) {
	wrapper.Insights = insights
}

type APIKey struct {
	ID          string               `json:"id"`
	AccessKey   string               `json:"access_key"`
	CreatedAt   string               `json:"created_at"`
	Status      string               `json:"status"`
	CreatedBy   CreatedBy            `json:"created_by"`
	Application ShortInfoApplication `json:"application"`
	Cluster     ShortInfoCluster     `json:"cluster"`
	ACL         contracts.ACL        `json:"acl"`
	ArxUrl      string               `json:"arx_url"`
}

type CreatedBy struct {
	ID           string `json:"id"`
	FullName     string `json:"name"`
	EmailAddress string `json:"email_address"`
	Avatar       string `json:"avatar"`
}

type APIKeyModelsWrapper struct {
	Applications []Application `json:"applications"`
	Clusters     []Cluster     `json:"clusters"`
}

type ShortInfoApplication struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Application struct {
	ShortInfoApplication  `json:",inline"`
	Label                 string                `json:"label"`
	ApplicationTechnology ApplicationTechnology `json:"application_technology"`
}

func (app *Application) SetTechnology(applicationTechnology ApplicationTechnology) {
	app.ApplicationTechnology = applicationTechnology
}

type ShortInfoCluster struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}
type Cluster struct {
	ShortInfoCluster `json:",inline"`
	Purpose          Purpose       `json:"purpose"`
	Plan             Plan          `json:"plan"`
	Provider         Provider      `json:"provider"`
	Regions          []Region      `json:"regions"`
	Encryption       Encryption    `json:"encryption"`
	RotationCycle    RotationCycle `json:"rotation_cycle"`
	CreatedBy        CreatedBy     `json:"created_by"`
}

func (c *Cluster) SetPurpose(purpose Purpose) {
	c.Purpose = purpose
}

func (c *Cluster) SetPlan(plan Plan) {
	c.Plan = plan
}

func (c *Cluster) SetProvider(provider Provider) {
	c.Provider = provider
}

func (c *Cluster) SetRegions(regions []Region) {
	if regions != nil {
		c.Regions = make([]Region, len(regions))
		copy(c.Regions, regions)
	}
}

func (c *Cluster) SetEncryption(encryption Encryption) {
	c.Encryption = encryption
}

func (c *Cluster) SetRotationCycle(rotationCycle RotationCycle) {
	c.RotationCycle = rotationCycle
}

func (c *Cluster) SetCreatedBy(owner CreatedBy) {
	c.CreatedBy = owner
}

type Purpose struct {
	Name string `json:"name"`
}

type Plan struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Provider struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Region struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Encryption struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type RotationCycle struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Owner struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Avatar       string `json:"avatar"`
}

type Insights struct {
	TotalKeys    int64 `json:"total_keys"`
	TotalActive  int64 `json:"total_active"`
	TotalDeleted int64 `json:"total_deleted"`
}
