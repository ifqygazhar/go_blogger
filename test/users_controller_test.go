package test

import (
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"go_blogger/app"
	"go_blogger/controller"
	"go_blogger/repository"
	"go_blogger/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupRouter() http.Handler {
	//err := helper.Load()
	//if err != nil {
	//	log.Fatalf("error load .env")
	//}
	var db = app.Dbtest()
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserControllerImplementation(userService)
	router := app.NewRouter(userController)

	return router
}

func TestCreateUsersSuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name" : "Gadget","password" : "anjay"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
