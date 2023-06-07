package common

type ShortUserInfo struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	Avatar       string `json:"avatar"`
}

func (ShortUserInfo) TableName() string {
	return "users"
}
