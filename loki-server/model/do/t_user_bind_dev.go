package do

type TUserBindDev struct {
	Base   `xorm:"extends"`
	UserId int64
	DevId  int64
}
