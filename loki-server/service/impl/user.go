package service_impl

import (
	"context"
	"jobscn/ai-flower-pot/loki"
	"jobscn/ai-flower-pot/loki-server/dao"
	"jobscn/ai-flower-pot/loki-server/model/do"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/pkg/errcode"
)

type UserService struct {
	UserDao dao.IUserDao `inject:""`
}

func (p *UserService) Register(ctx context.Context, in *dto.UserRegisterParam) error {
	err := p.UserDao.Insert(loki.DBEngine, &do.TUser{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return errcode.USER_InsertDBErr
	}

	return nil
}
