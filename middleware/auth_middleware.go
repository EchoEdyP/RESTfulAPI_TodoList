package middleware

import (
	"PZN_RESTfulAPI/helper"
	"PZN_RESTfulAPI/model/api"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		apiResponse := api.ApiResponse{
			Status:  http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, apiResponse)
	}
}
