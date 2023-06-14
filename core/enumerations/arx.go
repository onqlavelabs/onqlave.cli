package enumerations

type ArxStatus string

var (
	ArxActive      ArxStatus = "active"
	ArxSealed      ArxStatus = "sealed"
	ArxUnsealed    ArxStatus = "unsealed"
	ArxInactive    ArxStatus = "inactive"
	ArxPending     ArxStatus = "pending"
	ArxFailed      ArxStatus = "failed"
	ArxDeleted     ArxStatus = "deleted"
	ArxInitiated   ArxStatus = "initiated"
	ArxReInitiated ArxStatus = "reinitiated"
)

func (status ArxStatus) String() string {
	return string(status)
}
