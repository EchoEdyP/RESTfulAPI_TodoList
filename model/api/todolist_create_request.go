package api

type TodolistCreateRequest struct {
	Title       string `json:"title" validate:"required,min=5,max=50"`
	Description string `json:"description" validate:"required,min=5"`
}
