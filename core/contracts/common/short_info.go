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

type ShortResourceInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
}
