package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	u := uuid.NewRandom()
	return scope.SetColumn("Id", u.String())
}
