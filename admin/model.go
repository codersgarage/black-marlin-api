package admin

import (
	"strings"
	"time"

	"github.com/codersgarage/black-marlin-api/errors"
)

type Admin struct {
	ID        int       `sql:"id;primary key;auto_increment" json:"-"`
	UUID      string    `sql:"uuid;not null;unique" json:"uuid"`
	Name      string    `sql:"name" json:"name"`
	Email     string    `sql:"email" json:"email"`
	Password  string    `sql:"password" json:"password"`
	CreatedAt time.Time `sql:"created_at;not null" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at;not null" json:"updated_at"`
}

func (h *Admin) TableName() string {
	return "admins"
}

func (h *Admin) Validate() *errors.ValidationError {
	h.Name = strings.TrimSpace(h.Name)

	ve := errors.ValidationError{}

	if h.Name == "" {
		ve.Add("name", "is required")
	}

	if len(ve) == 0 {
		return nil
	}

	return &ve
}
