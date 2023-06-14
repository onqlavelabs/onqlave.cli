package contracts

type UpdateApplication struct {
	Name        string            `json:"name" validate:"required,max=150"`
	Description string            `json:"description" validate:"max=500"`
	Technology  string            `json:"technology" validate:"required,max=20"`
	Owner       string            `json:"owner" validate:"required,max=150"`
	Cors        []ApplicationCors `json:"cors" validate:"max=10"`
}
