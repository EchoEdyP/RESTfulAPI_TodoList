package main

import (
	"PZN_RESTfulAPI/app"
	"PZN_RESTfulAPI/controller"
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/middleware"
	"PZN_RESTfulAPI/repository"
	"PZN_RESTfulAPI/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := app.NewDB()
	validate := validator.New()
	todolistRepository := repository.NewTodolistRepository()
	todolistService := service.NewTodolistService(todolistRepository, db, validate)
	todolistController := controller.NewTodolistController(todolistService)
	router := app.NewRouter(todolistController)

	server := http.Server{
		Addr:    "localhost:1234",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)

}
