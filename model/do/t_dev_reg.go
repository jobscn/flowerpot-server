package do

type TDevReg struct {
	Base `xorm:"extends"`
	DevSerialize string `xorm:"default('')"`
	DevToken     string `xorm:"default('')"`
	Region       string `xorm:"default('')"`
	ProductId    int64 `xorm:"default(0)"`
	DevOname     string `xorm:"default('')"`
}
