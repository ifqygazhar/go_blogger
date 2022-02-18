package service

import (
	"context"
	"go_blogger/model/web"
)

type UserService interface {
	RegisterUser(ctx context.Context, request web.CreateUserRequest) web.UserResponse
	LoginUser(ctx context.Context, request web.LoginUserRequest) web.UserResponse
	Update(ctx context.Context,request web.UpdateUserRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
}
