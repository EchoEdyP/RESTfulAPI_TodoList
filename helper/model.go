package helper

import (
	"PZN_RESTfulAPI/model/api"
	"PZN_RESTfulAPI/model/domain"
)

func ToTodolistResponse(todolist domain.Todolist) api.TodolistResponse {
	return api.TodolistResponse{
		Id:          todolist.Id,
		Title:       todolist.Title,
		Description: todolist.Description,
		Status:      todolist.Status,
	}
}

func ToTodolistResponses(todoliests []domain.Todolist) []api.TodolistResponse {
	var todolistResponses []api.TodolistResponse
	for _, todolist := range todoliests {
		todolistResponses = append(todolistResponses, ToTodolistResponse(todolist))
	}

	return todolistResponses
}
