package repository

import (
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/model/domain"
	"context"
	"database/sql"
	"errors"
)

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() TodolistRepository {
	return &TodolistRepositoryImpl{}
}

func (repository *TodolistRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQL := "INSERT INTO TodoList(title, description) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQL, todolist.Title, todolist.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todolist.Id = int(id)
	return todolist
}

func (repository *TodolistRepositoryImpl) UpdateTitleDescription(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQL := "UPDATE TodoList SET title=?, description=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, todolist.Id, todolist.Title, todolist.Description)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodolistRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQL := "UPDATE TodoList SET status=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, todolist.Id, todolist.Status)
	helper.PanicIfError(err)

	return todolist
}

func (repository *TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) {
	SQL := "DELETE FROM TodoList WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, todolist.Id)
	helper.PanicIfError(err)
}

func (repository *TodolistRepositoryImpl) ReadById(ctx context.Context, tx *sql.Tx, todolistId int) (domain.Todolist, error) {
	SQL := "SELECT * FROM Todolist WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, todolistId)
	helper.PanicIfError(err)
	defer rows.Close()

	todolist := domain.Todolist{}
	if rows.Next() {
		err := rows.Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.Status)
		helper.PanicIfError(err)
		return todolist, nil
	} else {
		return todolist, errors.New("todolist is not found")
	}
}

func (repository *TodolistRepositoryImpl) ReadAll(ctx context.Context, tx *sql.Tx) []domain.Todolist {
	SQL := "SELECT * FROM Todolist"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todoliests []domain.Todolist
	for rows.Next() {
		todolist := domain.Todolist{}
		err := rows.Scan(&todolist.Id, &todolist.Title, &todolist.Description, &todolist.Status)
		helper.PanicIfError(err)
		todoliests = append(todoliests, todolist)
	}
	return todoliests
}
