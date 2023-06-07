package contracts

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/common"

type APIKeyDetail struct {
	ID            string                `json:"id"`
	ClusterID     string                `json:"cluster_id,omitempty"`
	ApplicationID string                `json:"application_id,omitempty"`
	CreatedAt     string                `json:"created_at"`
	Status        string                `json:"status"`
	AccessKey     string                `json:"access_key,omitempty"`
	CreatedBy     *common.ShortUserInfo `json:"created_by,omitempty"`
	ArxUrl        string                `json:"arx_url"`
}
