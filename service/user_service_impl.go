package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"go_blogger/exception"
	"go_blogger/helper"
	"go_blogger/model/domain"
	"go_blogger/model/web"
	"go_blogger/repository"
)

type UserServiceImplementation struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImplementation{UserRepository: userRepository, DB: DB, validate: validator}
}

func (service *UserServiceImplementation) RegisterUser(ctx context.Context, request web.CreateUserRequest) web.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userInput := domain.User{
		Name:     request.Name,
		Password: request.Password,
	}

	userSignUp := service.UserRepository.SignUp(ctx, tx, userInput)
	return helper.ToUserResponse(userSignUp)

}

func (service *UserServiceImplementation) LoginUser(ctx context.Context, request web.LoginUserRequest) web.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userInput := domain.User{
		Name:     request.Name,
		Password: request.Password,
	}

	userLogin, err := service.UserRepository.Login(ctx, tx, userInput)
	if err != nil {
		errors.New("tidak ada data nama dan password")
	}
	return helper.ToUserResponse(userLogin)
}

func (service *UserServiceImplementation) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)

}

func (service *UserServiceImplementation) Update(ctx context.Context, request web.UpdateUserRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = request.Name
	user.Password = request.Password

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImplementation) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}
