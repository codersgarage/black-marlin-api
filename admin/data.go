package admin

import (
	"github.com/codersgarage/black-marlin-api/app"
)

type Data struct {
}

func NewData() *Data {
	return &Data{}
}

func (hd *Data) Save(scope *app.Scope, admin *Model) error {
	return scope.DB.Table(admin.TableName()).Create(admin).Error
}

func (hd *Data) Get(scope *app.Scope, UUID string, v *Model) error {
	return scope.DB.
		Table(v.TableName()).
		Find(v, "uuid = ?", UUID).Error
}

func (hd *Data) List(scope *app.Scope, v interface{}) error {
	admin := Model{}
	return scope.DB.
		Table(admin.TableName()).
		Find(v).Error
}
