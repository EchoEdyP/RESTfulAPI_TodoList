package app

import (
	"PZN_RESTfulAPI/controller"
	"PZN_RESTfulAPI/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(todolistController controller.TodolistController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", todolistController.ReadAll)
	router.GET("/api/categories/:categoryId", todolistController.ReadById)
	router.POST("/api/categories", todolistController.Create)
	router.PUT("/api/todoliests/:categoryId", todolistController.UpdateTitleDescription)
	router.PUT("/api/categories/:categoryId", todolistController.UpdateStatus)
	router.DELETE("/api/categories/:categoryId", todolistController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
