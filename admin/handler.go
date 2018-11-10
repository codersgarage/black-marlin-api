package admin

import (
	"net/http"

	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/errors"
)

type Methods interface {
	Save(s *app.Scope) (*Model, error)
	Get(s *app.Scope) (*Model, error)
	Update(s *app.Scope) (*Model, error)
	List(s *app.Scope) ([]Model, error)
	Delete(s *app.Scope) error
}

type Handler struct {
	Repo Methods
}

func NewRoutes() *Handler {
	return &Handler{
		Repo: NewRepo(),
	}
}

func (hr *Handler) save(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Model, err := hr.Repo.Save(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusCreated
	resp.Data = Model
	resp.Title = "Successfully created"
	resp.ServerJSON(w)
}

func (hr *Handler) get(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Model, err := hr.Repo.Get(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Model
	resp.ServerJSON(w)
}

func (hr *Handler) update(w http.ResponseWriter, r *http.Request) {
	resp := response{}
	Model, err := hr.Repo.Update(app.NewScope(app.DB(), r))
	if err != nil {
		resp.Status = err.(*errors.APIError).Status
		resp.Code = err.(*errors.APIError).Code
		resp.Title = err.(*errors.APIError).Title
		resp.Errors = err.(*errors.APIError).Errors
		resp.ServerJSON(w)
		return
	}

	resp.Status = http.StatusOK
	resp.Data = Model
	resp.ServerJSON(w)
}

func (hr *Handler) list(w http.ResponseWriter, r *http.Request) {
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

func (hr *Handler) delete(w http.ResponseWriter, r *http.Request) {

}
