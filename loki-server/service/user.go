package service

import (
	"context"
	"jobscn/ai-flower-pot/loki-server/model/dto"
)

type IUserService interface {
	Register(ctx context.Context, in *dto.UserRegisterParam) error
}
