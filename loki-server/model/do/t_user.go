package do

type TUser struct {
	Base     `xorm:"extends"`
	Username string
	Password string
	Phone    string
	Avatar   string
	Gender   int
}
