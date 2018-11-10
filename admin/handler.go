package admin

import (
	"net/http"

	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/errors"
)

type ApiRepo interface {
	Save(s *app.Scope) (*Admin, error)
	Get(s *app.Scope) (*Admin, error)
	Update(s *app.Scope) (*Admin, error)
	List(s *app.Scope) ([]Admin, error)
	Delete(s *app.Scope) error
}

type ApiRoutes struct {
	Repo ApiRepo
}

func NewRoutes() *ApiRoutes {
	return &ApiRoutes{
		Repo: NewRepo(),
	}
}

func (hr *ApiRoutes) save(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Admin, err := hr.Repo.Save(app.NewScope(app.DB(), r))
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
	Admin, err := hr.Repo.Get(app.NewScope(app.DB(), r))
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
	Admin, err := hr.Repo.Update(app.NewScope(app.DB(), r))
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
	Admins, err := hr.Repo.List(app.NewScope(app.DB(), r))
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
