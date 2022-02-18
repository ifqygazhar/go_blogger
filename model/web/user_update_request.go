package web

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,max=200,min=1" json:"name"`
	Password string `validate:"required" json:"password"`
}
