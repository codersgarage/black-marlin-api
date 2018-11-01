package api

import (
	"net/http"
)

type AuthRoutes struct {
}

func NewAuthRoutes() *AuthRoutes {
	return &AuthRoutes{}
}

func (hr *AuthRoutes) registerUser(w http.ResponseWriter, r *http.Request) {
	resp := response{}

	resp.Status = http.StatusCreated
	resp.Title = "User signed up!"
	resp.ServerJSON(w)
}

func (hr *AuthRoutes) loginUser(w http.ResponseWriter, r *http.Request) {
	resp := response{}

	resp.Status = http.StatusOK
	resp.Title = "Login successfully!"
	resp.ServerJSON(w)
}
