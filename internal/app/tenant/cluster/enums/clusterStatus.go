package enums

// AccountStatementStatus status
type ClusterStatus string

// All status
var (
	ClusterActive      ClusterStatus = "active"
	ClusterSealed      ClusterStatus = "sealed"
	ClusterUnsealed    ClusterStatus = "unsealed"
	ClusterInactive    ClusterStatus = "inactive"
	ClusterPending     ClusterStatus = "pending"
	ClusterFailed      ClusterStatus = "failed"
	ClusterDeleted     ClusterStatus = "deleted"
	ClusterInitiated   ClusterStatus = "initiated"
	ClusterReInitiated ClusterStatus = "reinitiated"
)

func (status ClusterStatus) String() string {
	return string(status)
}
