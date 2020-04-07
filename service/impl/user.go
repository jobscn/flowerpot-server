package impl

import (
	"context"
	"jobscn/ai-flower-pot/model/dto"
)

type UserService struct{

}

func (u *UserService) Login(ctx context.Context, in *dto.UserLoginIn) (*dto.UserLoginOut, error) {
	return nil, nil
}

