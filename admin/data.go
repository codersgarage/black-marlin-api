package admin

import (
	"github.com/codersgarage/black-marlin-api/app"
)

type data struct {
}

func newData() *data {
	return &data{}
}

func (hd *data) Save(scope *app.Scope, admin *model) error {
	return scope.DB.Table(admin.TableName()).Create(admin).Error
}

func (hd *data) Get(scope *app.Scope, UUID string, v *model) error {
	return scope.DB.
		Table(v.TableName()).
		Find(v, "uuid = ?", UUID).Error
}

func (hd *data) List(scope *app.Scope, v interface{}) error {
	admin := model{}
	return scope.DB.
		Table(admin.TableName()).
		Find(v).Error
}
