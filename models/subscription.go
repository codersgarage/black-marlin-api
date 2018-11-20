package models

import "time"

type Subscription struct {
	ID        int64     `sql:"id;primary key;auto_increment" json:"id"`
	AppID     int64     `sql:"app_id;not null;index" json:"app_id"`
	UserUUID  string    `sql:"user_uuid;not null;index" json:"user_uuid"`
	CreatedAt time.Time `sql:"created_at;not null" json:"created_at"`
}

func (s *Subscription) TableName() string {
	return "subscriptions"
}
