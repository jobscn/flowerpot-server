package service

import (
	"context"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki/common"
)

type IUserService interface {
	Login(ctx context.Context, in *dto.UserLoginIn) (*dto.UserLoginOut, error)
	Register(ctx context.Context, in *dto.UserRegisterByPhoneIn) (*dto.UserRegisterByPhoneOut, common.Error)
}
