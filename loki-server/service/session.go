package service

import (
	"context"
	"jobscn/ai-flower-pot/loki-server/model/dto"
)

type ISessionService interface {
	RefreshToken(ctx context.Context, in *dto.RefreshTokenParam) (*dto.SessionToken, error)
	Login(ctx context.Context, in *dto.SessionLoginParam) (*dto.SessionToken, error)
}
