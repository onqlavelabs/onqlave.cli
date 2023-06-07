package contracts

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/acl/contracts"

type GetApplicationsResponseWrapper struct {
	ACL          contracts.ACL                    `json:"acl"`
	Applications []ExistingApplicationWithDetails `json:"applications"`
	Models       ApplicationModelWrapper          `json:"model"`
	Statistics   ApplicationStatistics            `json:"statistics"`
}
