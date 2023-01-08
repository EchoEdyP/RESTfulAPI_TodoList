package api

type TodolistUpdateTitleDescriptionRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,min=5,max=50"`
	Description string `json:"description" validate:"required,min=5"`
}
