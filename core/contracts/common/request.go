package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
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

func NewBaseErrorResponse(code int, status, message, cId string, details []interface{}) BaseErrorResponse {
	return BaseErrorResponse{
		Error: Error{
			Code:          code,
			Message:       message,
			Status:        status,
			CorrelationID: cId,
			Details:       details,
		},
	}
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

func EmptyRegistrationID() RegistrationID {
	return RegistrationID("")
}

func EmptyDomainId() DomainId {
	return DomainId("")
}

func EmptyApplicationId() ApplicationId {
	return ApplicationId("")
}

func EmptyUserId() UserId {
	return UserId("")
}

func EmptyTenantId() TenantId {
	return TenantId("")
}

func EmptyArxId() ArxId {
	return ArxId("")
}

func NewArxProvider(t string) ArxProvider {
	for k, v := range ArxProviders {
		if strings.EqualFold(strings.ToLower(v), strings.ToLower(t)) {
			return k
		}
	}
	return ProviderInvalid
}

func NewArxType(t string) ArxType {
	for k, v := range ArxTypes {
		if strings.EqualFold(strings.ToLower(v), strings.ToLower(t)) {
			return k
		}
	}
	return InvalidCluster
}

func NewArxPurpose(t string) ArxPurpose {
	for k, v := range ArxPurposes {
		if strings.EqualFold(strings.ToLower(v), strings.ToLower(t)) {
			return k
		}
	}
	return PurposeInvalid
}

func NewArxProvisioningState(t string) ArxProvisioningState {
	for k, v := range ArxProvisioningStates {
		if strings.EqualFold(strings.ToLower(v), strings.ToLower(t)) {
			return k
		}
	}
	return ClusterInvalid
}

func NewArxRegion(t string) ArxRegion {
	for k, v := range ArxRegions {
		if strings.EqualFold(strings.ToLower(v), strings.ToLower(t)) {
			return k
		}
	}
	return REGION_INVALID
}

func NewArxId(tenantId TenantId, clusterProvider ArxProvider, clusterType ArxType) ArxId {
	//id := fmt.Sprintf("cluster##%s##%s##%s##%s", tenantId, clusterProvider.String(), clusterType.String(), uuid.New().String())
	id := fmt.Sprintf("cluster--%s", uuid.New().String())

	return ArxId(id)
}

func NewTenantId() TenantId {
	id := fmt.Sprintf("tenant--%s", uuid.New().String())
	return TenantId(id)
}

func ExisingTenantId(id string) TenantId {
	return TenantId(id)
}

func ExisingUserId(id string) UserId {
	return UserId(id)
}

func ExisingArxId(id string) ArxId {
	return ArxId(id)
}

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
