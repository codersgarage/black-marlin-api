package admin

import (
	"net/http"

	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/errors"
)

type ApiRepo interface {
	SaveAdmin(s *app.Scope) (*Admin, error)
	GetAdmin(s *app.Scope) (*Admin, error)
	UpdateAdmin(s *app.Scope) (*Admin, error)
	ListAdmin(s *app.Scope) ([]Admin, error)
	DeleteAdmin(s *app.Scope) error
}

type ApiRoutes struct {
	Repo ApiRepo
}

func NewAdminRoutes() *ApiRoutes {
	return &ApiRoutes{
		Repo: NewAdminRepo(),
	}
}

func (hr *ApiRoutes) save(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Admin, err := hr.Repo.SaveAdmin(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusCreated
	resp.Data = Admin
	resp.Title = "Admin has been created"
	resp.ServerJSON(w)
}

func (hr *ApiRoutes) get(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Admin, err := hr.Repo.GetAdmin(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Admin
	resp.ServerJSON(w)
}

func (hr *ApiRoutes) update(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Admin, err := hr.Repo.UpdateAdmin(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Admin
	resp.ServerJSON(w)
}

func (hr *ApiRoutes) list(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Admins, err := hr.Repo.ListAdmin(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Admins
	resp.ServerJSON(w)
}

func (hr *ApiRoutes) delete(w http.ResponseWriter, r *http.Request) {

}
