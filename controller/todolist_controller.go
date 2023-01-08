package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TodolistController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTitleDescription(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ReadById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ReadAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
