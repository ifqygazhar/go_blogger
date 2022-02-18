package app

import (
	"github.com/julienschmidt/httprouter"
	"go_blogger/controller"
	"go_blogger/exception"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users/:userId", userController.FindById )
	router.POST("/api/users", userController.RegisterUser)
	router.PUT("/api/users/:userId", userController.Update)
	router.DELETE("/api/users/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandle

	return router

}