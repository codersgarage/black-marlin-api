package models

import "time"

type App struct {
	ID          int       `sql:"id;primary key;auto_increment" json:"id"`
	UUID        string    `sql:"uuid;index;unique;not null" json:"uuid"`
	Name        string    `sql:"name;not null" json:"name"`
	Description string    `sql:"description;not null" json:"description"`
	Key         string    `sql:"key;index;unique;not null" json:"key"`
	CreatedAt   time.Time `sql:"created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at;not null" json:"updated_at"`
}

func (a *App) TableName() string {
	return "apps"
}
