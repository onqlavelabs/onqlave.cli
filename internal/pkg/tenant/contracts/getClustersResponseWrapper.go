package contracts

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/acl/contracts"

type GetArxsResponseWrapper struct {
	ACL      contracts.ACL                `json:"acl"`
	Clusters []ExistingClusterWithDetails `json:"clusters"`
	Models   ArxModelWrapper              `json:"model"`
	Insights ArxInsights                  `json:"insights"`
}

type ArxModelWrapper struct {
	Purposes          []ArxPurpose                 `json:"purposes" validate:"required"`
	Plans             []ArxPlan                    `json:"plans" validate:"required"`
	EncryptionMethods []ArxEncryptionMethod        `json:"encryption_methods" validate:"required"`
	RotationCycles    []ArxEncryptionRotationCycle `json:"rotation_cycles" validate:"required"`
	Providers         []ArxProvider                `json:"providers" validate:"required"`
}
