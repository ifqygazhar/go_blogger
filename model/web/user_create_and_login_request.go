package web

type CreateUserRequest struct {
	Name     string `json:"name" `
	Password string `json:"password" `
}

type LoginUserRequest struct {
	Name     string `json:"name" `
	Password string `json:"password"`
}
