package common

import (
	"regexp"
)

type BaseErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code          int           `json:"code" yaml:"code"`                     // Code represent codes in response
	Status        string        `json:"status" yaml:"status"`                 // Status represent string value of code
	Message       string        `json:"message" yaml:"message"`               // Message represent detail message
	CorrelationID string        `json:"correlation_id" yaml:"correlation_id"` // The RequestId that's also set in the header
	Details       []interface{} `json:"details" yaml:"details"`               // Details is a list of details in any types in string
}

type ArxType int64

type ArxProvider int64

func (c ArxProvider) String() string {
	return ArxProviders[c]
}

func (c ArxRegion) String() string {
	return ArxRegions[c]
}

type ArxId string
type ApplicationId string
type TenantId string
type RegistrationID string
type UserId string
type DomainId string
type UserEmail string
type ArxPurpose int64
type ArxRegion int64

func (id *ArxId) Valid() bool {
	// pattern := regexp.MustCompile(`##`)
	// result := pattern.Split(string(*id), -1)
	// return len(result) == 5
	return true
}

func (id *ApplicationId) Valid() bool {
	// pattern := regexp.MustCompile(`##`)
	// result := pattern.Split(string(*id), -1)
	// return len(result) == 5
	return true
}

func (id *TenantId) Valid() bool {
	pattern := regexp.MustCompile(`--`)
	result := pattern.Split(string(*id), -1)
	return len(result) == 2
}

var (
	ArxRegions = map[ArxRegion]string{REGION_INVALID: "Invalid", REGION_AU: "au",
		REGION_SG: "sg", REGION_US: "us", REGION_UK: "gb"}

	ArxTypes = map[ArxType]string{InvalidCluster: "Invalid", ServerlessCluster: "serverless",
		DedicatedCluster: "dedicated", OnPremCluster: "on-premise"}

	ArxProviders = map[ArxProvider]string{ProviderInvalid: "Invalid", ProviderAWS: "aws",
		ProviderAzure: "azure", ProviderGCP: "gcp"}

	ArxPurposes = map[ArxPurpose]string{PurposeInvalid: "Invalid", PurposeDevelopment: "development", PurposeTesting: "testing",
		PurposeProduction: "production", PurposeStaging: "staging"}

	ArxOperations = map[ArxOperation]string{OperationNone: "None", OperationFailed: "Failed",
		OperationReinitialise: "Retry", OperationChanged: "Changed", OperationCreated: "Created",
		OperationDeleted: "Deleted", OperationDisabled: "Disabled", OperationEnabled: "Enabled", OperationUpdated: "Updated"}

	ArxProvisioningStates = map[ArxProvisioningState]string{ClusterInvalid: "invalid", ClusterActive: "active",
		ClusterSealed: "sealed", ClusterUnseal: "unsealed", ClusterDelete: "deleted", ArxUpdate: "update", ClusterInactive: "inactive", ClusterFailed: "failed",
		ClusterReinitiated: "reinitiated", ClusterInitiated: "initiated", ClusterPending: "pending",
		ClusterTimedout: "timeout"}
)

func (o ArxPurpose) String() string {
	return ArxPurposes[o]
}

func (o ArxType) String() string {
	return ArxTypes[o]
}

func (o ArxProvisioningState) String() string {
	return ArxProvisioningStates[o]
}

func (o ArxOperation) String() string {
	return ArxOperations[o]
}

type ArxProvisioningState int64
type ArxOperation int64
type ArxEndpoint string

func (e *ArxEndpoint) Decode() (ArxEndpoint, error) {
	return ArxEndpoint(*e), nil
}

const (
	ProviderInvalid ArxProvider = iota
	ProviderAzure
	ProviderAWS
	ProviderGCP
)

const (
	REGION_INVALID ArxRegion = iota
	REGION_AU
	REGION_SG
	REGION_US
	REGION_UK
	REGION_AUS_EAST
	REGION_AUS_WEST
	// US_EAST
	// US_WEST
	// EU_EAST
	// EU_WEST
	// APAC_EAST
	// APAC_WEST
	// APAC_NORTH
)
const (
	OperationNone ArxOperation = iota
	OperationChanged
	OperationReinitialise
	OperationCreated
	OperationUpdated
	OperationDeleted
	OperationDisabled
	OperationEnabled
	OperationFailed
)

const (
	ClusterInvalid ArxProvisioningState = iota
	ClusterActive
	ClusterSealed
	ClusterUnseal
	ClusterDelete
	ArxUpdate
	ClusterInactive
	ClusterFailed
	ClusterReinitiated
	ClusterInitiated
	ClusterPending
	ClusterTimedout
)

const (
	InvalidCluster ArxType = iota
	ServerlessCluster
	DedicatedCluster
	OnPremCluster
)

const (
	PurposeInvalid ArxPurpose = iota
	PurposeDevelopment
	PurposeTesting
	PurposeProduction
	PurposeStaging
)
