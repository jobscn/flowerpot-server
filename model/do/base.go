package do

import "time"

type Base struct {
	Id         int64 `xorm:"autoincr pk"`
	CreateTime time.Time `xorm:"created default(0)"`
	UpdateTime time.Time `xorm:"updated default(0)"`
	DeleteTime time.Time `xorm:"deleted default(0)"`
}