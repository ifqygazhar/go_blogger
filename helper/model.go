package helper

import (
	"go_blogger/model/domain"
	"go_blogger/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id: user.Id,
		Name: user.Name,
	}
}
