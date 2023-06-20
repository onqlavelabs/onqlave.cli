package acl

type ACL struct {
	Can    map[string]bool   `json:"can"`
	CanNot map[string]string `json:"can_not"`
}
