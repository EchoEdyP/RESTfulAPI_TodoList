package service

import (
	"PZN_RESTfulAPI/exception"
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/model/api"
	"PZN_RESTfulAPI/model/domain"
	"PZN_RESTfulAPI/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type TodolistServiceImpl struct {
	TodolistRepository repository.TodolistRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewTodolistService(todolistRepository repository.TodolistRepository, DB *sql.DB, validate *validator.Validate) TodolistService {
	return &TodolistServiceImpl{
		TodolistRepository: todolistRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *TodolistServiceImpl) Create(ctx context.Context, r api.TodolistCreateRequest) api.TodolistResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist := domain.Todolist{
		Title:       r.Title,
		Description: r.Description,
	}

	todolist = service.TodolistRepository.Save(ctx, tx, todolist)

	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) UpdateTitleDescription(ctx context.Context, request api.TodolistUpdateTitleDescriptionRequest) api.TodolistResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.ReadById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todolist.Title = request.Title
	todolist.Description = request.Description

	todolist = service.TodolistRepository.UpdateTitleDescription(ctx, tx, todolist)

	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) UpdateStatus(ctx context.Context, request api.TodolistUpdateStatusRequest) api.TodolistResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.ReadById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todolist.Status = request.Status

	todolist = service.TodolistRepository.UpdateStatus(ctx, tx, todolist)

	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) Delete(ctx context.Context, todolistId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.ReadById(ctx, tx, todolistId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodolistRepository.Delete(ctx, tx, todolist)
}

func (service *TodolistServiceImpl) ReadById(ctx context.Context, todolistId int) api.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todolist, err := service.TodolistRepository.ReadById(ctx, tx, todolistId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodolistResponse(todolist)
}

func (service *TodolistServiceImpl) ReadAll(ctx context.Context) []api.TodolistResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todoliests := service.TodolistRepository.ReadAll(ctx, tx)

	return helper.ToTodolistResponses(todoliests)
}
