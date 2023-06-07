package contracts

type NewCluster struct {
	Name             string   `json:"name" validate:"required"`
	SpendLimit       *uint64  `json:"spend_limit" validate:"required"`
	Purpose          string   `json:"purpose" validate:"required,max=50"`
	Plan             string   `json:"plan" validate:"required,max=50"`
	Provider         string   `json:"provider" validate:"required,max=50"`
	Regions          []string `json:"regions" validate:"required,max=5"`
	EncryptionMethod string   `json:"encryption_method" validate:"required,max=50"`
	RotationCycle    string   `json:"rotation_cycle" validate:"required,max=50"`
	Owner            string   `json:"owner" validate:"required,max=150"`
	IsDefault        bool     `json:"is_default"`
}
