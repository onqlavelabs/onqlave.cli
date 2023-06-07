package contracts

type ApplicationCors struct {
	Address string `json:"address" validate:"required,max=300,url"`
}
