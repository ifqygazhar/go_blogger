package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"go_blogger/app"
	"go_blogger/controller"
	"go_blogger/helper"
	"go_blogger/middleware"
	"go_blogger/repository"
	"go_blogger/service"
	"log"
	"net/http"
)

func main() {
	err := helper.Load()
	if err != nil {
		log.Fatalf("error load .env")
	}
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserControllerImplementation(userService)

	router := app.NewRouter(userController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	errServ := server.ListenAndServe()
	helper.PanicIfError(errServ)

}
