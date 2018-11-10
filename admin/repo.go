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
	Dao *AdminDao
}

func NewRepo() *Repo {
	return &Repo{
		Dao: NewAdminDao(),
	}
}

func (hr *Repo) Save(s *app.Scope) (*Admin, error) {
	admin := &Admin{}
	if err := utils.ParseBody(s.Request, admin); err != nil {
		return nil, errors.NewAPIError(http.StatusBadRequest, "400001", "Unable to parse request body", nil)
	}
	if err := admin.Validate(); err != nil {
		return nil, errors.NewAPIError(http.StatusUnprocessableEntity, "422001", "Invalid data", err)
	}

	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	if err := hr.Dao.SaveAdmin(s, admin); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to save admin", err)
	}
	return admin, nil
}

func (hr *Repo) Get(s *app.Scope) (*Admin, error) {
	admin := &Admin{}

	ID := chi.URLParam(s.Request, "id")
	err := hr.Dao.GetAdmin(s, ID, admin)

	if err != nil && errors.IsRecordNotFoundError(err) {
		return nil, errors.NewAPIError(http.StatusNotFound, "404001", "admin not found", err)
	}
	if err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to get admin", err)
	}
	return admin, nil
}

func (hr *Repo) Update(s *app.Scope) (*Admin, error) {
	return &Admin{}, nil
}

func (hr *Repo) List(s *app.Scope) ([]Admin, error) {
	var data []Admin
	if err := hr.Dao.ListAdmins(s, &data); err != nil {
		return nil, errors.NewAPIError(http.StatusInternalServerError, "500001", "Failed to list Admins", err)
	}
	if len(data) > 0 {
		return data, nil
	}
	return []Admin{}, nil
}

func (hr *Repo) Delete(s *app.Scope) error {
	return nil
}
