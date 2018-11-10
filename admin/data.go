package admin

import (
	"github.com/codersgarage/black-marlin-api/app"
)

type AdminDao struct {
}

func NewAdminDao() *AdminDao {
	return &AdminDao{}
}

func (hd *AdminDao) Save(scope *app.Scope, admin *Admin) error {
	return scope.DB.Table(admin.TableName()).Create(admin).Error
}

func (hd *AdminDao) Get(scope *app.Scope, UUID string, v *Admin) error {
	return scope.DB.
		Table(v.TableName()).
		Find(v, "uuid = ?", UUID).Error
}

func (hd *AdminDao) List(scope *app.Scope, v interface{}) error {
	admin := Admin{}
	return scope.DB.
		Table(admin.TableName()).
		Find(v).Error
}
