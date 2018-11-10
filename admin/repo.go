package admin

import (
	"net/http"
	"time"

	"github.com/codersgarage/black-marlin-api/app"
	"github.com/codersgarage/black-marlin-api/errors"
	"github.com/codersgarage/black-marlin-api/utils"
	"github.com/go-chi/chi"
)

type Repo struct {
	data *data
}

func NewRepo() *Repo {
	return &Repo{
		data: newData(),
	}
}

func (hr *Repo) Save(s *app.Scope) (*Model, error) {
	admin := &Model{}
	if err := utils.ParseBody(s.Request, admin); err != nil {
		return nil, errors.NewAPIError(http.StatusBadRequest, "400001", "Unable to parse request body", nil)
	}
	if err := admin.Validate(); err != nil {
		return nil, errors.NewAPIError(http.StatusUnprocessableEntity, "422001", "Invalid data", err)
	}

	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	if err := hr.data.Save(s, admin); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to save admin", err)
	}
	return admin, nil
}

func (hr *Repo) Get(s *app.Scope) (*Model, error) {
	admin := &Model{}

	ID := chi.URLParam(s.Request, "id")
	err := hr.data.Get(s, ID, admin)

	if err != nil && errors.IsRecordNotFoundError(err) {
		return nil, errors.NewAPIError(http.StatusNotFound, "404001", "admin not found", err)
	}
	if err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to get admin", err)
	}
	return admin, nil
}

func (hr *Repo) Update(s *app.Scope) (*Model, error) {
	return &Model{}, nil
}

func (hr *Repo) List(s *app.Scope) ([]Model, error) {
	var data []Model
	if err := hr.data.List(s, &data); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to list Admins", err)
	}
	if len(data) > 0 {
		return data, nil
	}
	return []Model{}, nil
}

func (hr *Repo) Delete(s *app.Scope) error {
	return nil
}
