package web

type CreateUserRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}

type LoginUserRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
