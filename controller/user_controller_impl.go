package controller

import (
	"github.com/julienschmidt/httprouter"
	"go_blogger/helper"
	"go_blogger/model/web"
	"go_blogger/service"
	"net/http"
	"strconv"
)

type UserControllerImplementation struct {
	UserService service.UserService
}

func NewUserControllerImplementation(userService service.UserService) *UserControllerImplementation {
	return &UserControllerImplementation{
		UserService: userService,
	} 
}

func (controller UserControllerImplementation) RegisterUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.CreateUserRequest{}
	helper.ReadFromRequestBody(request,&userCreateRequest)

	userResponse := controller.UserService.RegisterUser(request.Context(),userCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteResponBody(writer,webResponse)
}

func (controller UserControllerImplementation) LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.LoginUserRequest{}
	helper.ReadFromRequestBody(request,&userLoginRequest)

	userResponse := controller.UserService.LoginUser(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteResponBody(writer,webResponse)
}

func (controller UserControllerImplementation) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UpdateUserRequest{}
	helper.ReadFromRequestBody(request,&userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(),userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteResponBody(writer,webResponse)
}

func (controller UserControllerImplementation) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.FindById(request.Context(),id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteResponBody(writer,webResponse)
}

func (controller UserControllerImplementation) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id,err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteResponBody(writer,webResponse)
}
