package contracts

import "github.com/onqlavelabs/onqlave.cli/core/contracts/acl"

type GetApplicationsResponseWrapper struct {
	ACL          acl.ACL                          `json:"acl"`
	Applications []ExistingApplicationWithDetails `json:"applications"`
	Models       ApplicationModelWrapper          `json:"model"`
	Statistics   ApplicationStatistics            `json:"statistics"`
}
