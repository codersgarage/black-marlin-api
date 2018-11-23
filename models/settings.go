package models

type Settings struct {
	ID                 int    `sql:"id;primary key" json:"id"`
	ProjectName        string `sql:"project_name;not null" json:"project_name"`
	ProjectDescription string `sql:"project_description;not null" json:"project_description"`
	SystemKey          string `sql:"system_key;not null" json:"system_key"`
	ServiceKey         string `sql:"service_key;not null" json:"service_key"`
}

func (s *Settings) TableName() string {
	return "settings"
}
