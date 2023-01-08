package repository

import (
	"PZN_RESTfulAPI/model/domain"
	"context"
	"database/sql"
)

type TodolistRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist
	UpdateTitleDescription(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist
	UpdateStatus(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist
	Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist)
	ReadById(ctx context.Context, tx *sql.Tx, todolistId int) (domain.Todolist, error)
	ReadAll(ctx context.Context, tx *sql.Tx) []domain.Todolist
}
