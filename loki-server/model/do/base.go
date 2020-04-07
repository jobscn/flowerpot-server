package do

import "time"

type Base struct {
	Id         int64 `xorm:"autoincr pk"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
	DeleteTime time.Time `xorm:"deleted"`
}