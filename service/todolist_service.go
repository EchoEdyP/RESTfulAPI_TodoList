package service

import (
	"PZN_RESTfulAPI/model/api"
	"context"
)

type TodolistService interface {
	Create(ctx context.Context, r api.TodolistCreateRequest) api.TodolistResponse
	UpdateTitleDescription(ctx context.Context, r api.TodolistUpdateTitleDescriptionRequest) api.TodolistResponse
	UpdateStatus(ctx context.Context, r api.TodolistUpdateStatusRequest) api.TodolistResponse
	Delete(ctx context.Context, todolistId int)
	ReadById(ctx context.Context, todolistId int) api.TodolistResponse
	ReadAll(ctx context.Context) []api.TodolistResponse
}
