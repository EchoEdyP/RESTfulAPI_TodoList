package api

type TodolistUpdateStatusRequest struct {
	Id     int    `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}
