package dao

import (
	"jobscn/ai-flower-pot/loki-server/model/do"
	"xorm.io/xorm"
)

type IUserDao interface {
	GetByUsername(db *xorm.Engine, username string) (*do.TUser, error)
	GetByPhone(db *xorm.Engine, phone string) (*do.TUser, error)
	Insert(db *xorm.Engine, do *do.TUser) error
}
