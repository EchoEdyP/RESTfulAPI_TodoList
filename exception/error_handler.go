package exception

import (
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/model/api"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {

	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		apiResponse := api.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    exception.Error(),
		}

		helper.WriteToResponseBody(w, apiResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		apiResponse := api.ApiResponse{
			Status:  http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    exception.Error,
		}

		helper.WriteToResponseBody(w, apiResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	apiResponse := api.ApiResponse{
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data:    err,
	}

	helper.WriteToResponseBody(w, apiResponse)
}
