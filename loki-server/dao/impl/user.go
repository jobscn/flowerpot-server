package dao_impl

import (
	"fmt"
	"jobscn/ai-flower-pot/loki-server/model/do"
	"xorm.io/xorm"
)

type UserDao struct{}

func (p *UserDao) GetByUsername(db *xorm.Engine, username string) (*do.TUser, error) {
	where := map[string]interface{}{
		"username": username,
	}

	var user do.TUser

	_, err := db.Where(where).Get(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserDao) Insert(db *xorm.Engine, do *do.TUser) error {
	n, err := db.InsertOne(do)
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("insert fail")
	}

	return nil
}

func (p *UserDao) GetByPhone(db *xorm.Engine, phone string) (*do.TUser, error) {
	where := map[string]interface{}{
		"phone": phone,
	}

	var user do.TUser

	_, err := db.Where(where).Get(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
