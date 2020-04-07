package dao_impl

import (
	"fmt"
	"jobscn/ai-flower-pot/model/do"
	"xorm.io/xorm"
)

type UserDao struct {}

func (p *UserDao)  GetByUsername(db *xorm.Engine, username string) (*do.TUser, error) {
	user := &do.TUser{
		Username:   username,
	}

	has, err := db.Where(user).Get(user)
	if err != nil {
		return nil, err
	}
	if has != true {
		return nil, fmt.Errorf("no found result")
	}

	return user, nil
}

func (p *UserDao)  Insert(db *xorm.Engine, do *do.TUser) error {
	n, err := db.InsertOne(do)
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("insert fail")
	}

	return nil
}
