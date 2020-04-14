package service_impl

import (
	"context"
	"jobscn/ai-flower-pot/loki-server/dao"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/pkg/custerr"
)

type UserService struct {
	UserDao dao.IUserDao `inject:""`
}

func (u *UserService) Login(ctx context.Context, in *dto.UserLoginIn) (*dto.UserLoginOut, error) {
	return nil, custerr.UserLogin_PasswordIncorrectErr
}

func (u *UserService) Register(ctx context.Context, in *dto.UserRegisterByPhoneIn) (*dto.UserRegisterByPhoneOut, error) {
	return nil, nil
}
