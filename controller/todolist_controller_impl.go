package controller

import (
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/model/api"
	"PZN_RESTfulAPI/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type TodolistControllerImpl struct {
	TodolistService service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) TodolistController {
	return &TodolistControllerImpl{
		todolistService,
	}
}

func (controller *TodolistControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistCreateRequest := api.TodolistCreateRequest{}
	helper.ReadfromRequestBody(r, &todolistCreateRequest)

	todolistResponse := controller.TodolistService.Create(r.Context(), todolistCreateRequest)
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "Create todo SuccesFully",
		Data:    todolistResponse,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *TodolistControllerImpl) UpdateTitleDescription(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistUpdateTitleDescriptionRequest := api.TodolistUpdateTitleDescriptionRequest{}
	helper.ReadfromRequestBody(r, &todolistUpdateTitleDescriptionRequest)

	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistUpdateTitleDescriptionRequest.Id = id

	todolistResponse := controller.TodolistService.UpdateTitleDescription(r.Context(), todolistUpdateTitleDescriptionRequest)
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "Update title and description in todo SuccesFully",
		Data:    todolistResponse,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *TodolistControllerImpl) UpdateStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistUpdateStatusRequest := api.TodolistUpdateStatusRequest{}
	helper.ReadfromRequestBody(r, &todolistUpdateStatusRequest)

	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistUpdateStatusRequest.Id = id

	todolistResponse := controller.TodolistService.UpdateStatus(r.Context(), todolistUpdateStatusRequest)
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "update status in todo SuccesFully",
		Data:    todolistResponse,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *TodolistControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	controller.TodolistService.Delete(r.Context(), id)
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "Delete todo SuccesFully",
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *TodolistControllerImpl) ReadById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistId := params.ByName("todolistId")
	id, err := strconv.Atoi(todolistId)
	helper.PanicIfError(err)

	todolistResponse := controller.TodolistService.ReadById(r.Context(), id)
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "Read By Id todo SuccesFully",
		Data:    todolistResponse,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

func (controller *TodolistControllerImpl) ReadAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolistResponses := controller.TodolistService.ReadAll(r.Context())
	apiResponse := api.ApiResponse{
		Status:  http.StatusOK,
		Message: "Read All todo SuccesFully",
		Data:    todolistResponses,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
