package service

import (
	"context"
	"jobscn/ai-flower-pot/loki-server/model/dto"
)

type IUserService interface {
	Login(ctx context.Context, in *dto.UserLoginIn) (string, error)
	Register(ctx context.Context, in *dto.UserRegisterIn) error
}
