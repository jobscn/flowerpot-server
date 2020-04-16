package service_impl

import (
	"context"
	"jobscn/ai-flower-pot/loki"
	"jobscn/ai-flower-pot/loki-server/dao"
	"jobscn/ai-flower-pot/loki-server/model/do"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/pkg/errcode"
	"jobscn/ai-flower-pot/loki-server/pkg/util"
	"strings"
)

type UserService struct {
	UserDao dao.IUserDao `inject:""`
}

func (p *UserService) Login(ctx context.Context, in *dto.UserLoginIn) (string, error) {
	userInfo, err := p.UserDao.GetByUsername(loki.DBEngine, in.Username)
	if err != nil {
		panic(err)
		return "", errcode.USER_UserDaoGetByUsernameErr
	}

	if strings.Compare(userInfo.Password, in.Password) != 0 {
		return "", errcode.USER_PasswordIncorrectErr
	}

	token, err := util.GenAuthTokenByUserInfo(&dto.UserInfo{
		UID:      userInfo.Id,
		Username: userInfo.Username,
		Avatar:   userInfo.Avatar,
		Gender:   userInfo.Gender,
	})
	if err != nil {
		return "", errcode.USER_GenJWTTokenErr
	}

	return token, nil
}

func (p *UserService) Register(ctx context.Context, in *dto.UserRegisterIn) error {
	err := p.UserDao.Insert(loki.DBEngine, &do.TUser{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return errcode.USER_InsertDBErr
	}

	return nil
}
