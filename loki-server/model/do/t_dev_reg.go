package do

type TDevReg struct {
	Base         `xorm:"extends"`
	DevSerialize string
	DevToken     string
	Region       string
	ProductId    int64
	DevOname     string
}
