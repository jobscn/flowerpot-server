package do

type TUser struct {
	Base `xorm:"extends"`
	Username   string `xorm:"default('')"`
	Password   string `xorm:"default('')"`
	Phone      string `xorm:"default('')"`
	Avatar     string `xorm:"default('')"`
	Gender     int `xorm:"default(0)"`
}
