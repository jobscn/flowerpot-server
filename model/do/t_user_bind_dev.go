package do

type TUserBindDev struct {
	Base `xorm:"extends"`
	UserId     int64 `xorm:"default(0)"`
	DevId      int64 `xorm:"default(0)"`
}
