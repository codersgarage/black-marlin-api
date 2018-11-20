package models

import "time"

type Admin struct {
	ID        int64     `sql:"id;primary key;auto_increment" json:"id"`
	UUID      string    `sql:"uuid;not null;unique" json:"uuid"`
	Name      string    `sql:"name;not null;" json:"name"`
	Password  string    `sql:"password;not null" json:"-"`
	Email     string    `sql:"email;not null;unique" json:"email"`
	IsActive  bool      `sql:"is_active;not null;default:true" json:"is_active"`
	CreatedAt time.Time `sql:"created_at;not null" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at;not null" json:"updated_at"`
}

func (a *Admin) TableName() string {
	return "admins"
}
