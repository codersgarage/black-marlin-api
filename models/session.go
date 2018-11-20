package models

import "time"

type Session struct {
	ID           int64     `sql:"id;primary key;auto_increment" json:"id"`
	UserID       int64     `sql:"user_id;not null;index" json:"user_id"`
	AccessToken  string    `sql:"access_token;not null;index;unique" json:"access_token"`
	RefreshToken string    `sql:"refresh_token;not null;index;unique" json:"refresh_token"`
	CreatedAt    time.Time `sql:"created_at;not null" json:"created_at"`
	ExpiresIn    int64     `sql:"expires_in;not null" json:"expires_in"` // In hours
}

func (ss *Session) TableName() string {
	return "sessions"
}
