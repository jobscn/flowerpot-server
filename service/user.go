package service

import (
	"context"
	"jobscn/ai-flower-pot/model/dto"
)

type IUserService interface {
	Login(ctx context.Context, in *dto.UserLoginIn) (*dto.UserLoginOut, error)
}
