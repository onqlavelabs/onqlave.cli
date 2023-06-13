package enumerations

type ApiKeyStatus string

const (
	Pending  ApiKeyStatus = "pending"
	Active   ApiKeyStatus = "active"
	Disabled ApiKeyStatus = "disabled"
	Deleted  ApiKeyStatus = "deleted"
	Failed   ApiKeyStatus = "failed"
)

func (status ApiKeyStatus) String() string {
	return string(status)
}
