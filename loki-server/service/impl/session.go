package service_impl

import (
	"context"
	"jobscn/ai-flower-pot/loki"
	"jobscn/ai-flower-pot/loki-server/dao"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/pkg/errcode"
	"jobscn/ai-flower-pot/loki-server/pkg/util"
	"strings"
)

type SessionService struct {
	UserDao dao.IUserDao `inject:""`
}

func (p *SessionService) genSessionToken(userInfo *dto.UserInfo) (*dto.SessionToken, error) {
	// Generate Token
	token, expireIn, err := util.GenAuthTokenByUserInfo(
		util.AUTH_TYPE_TOKEN,
		&dto.UserInfo{
			UID:      userInfo.UID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			Gender:   userInfo.Gender,
		},
	)
	if err != nil {
		return nil, errcode.SESSION_GenJWTTokenErr
	}

	// Generate RefreshToken
	refreshToken, _, err := util.GenAuthTokenByUserInfo(
		util.AUTH_TYPE_REFRESH_TOKEN,
		&dto.UserInfo{
			UID:      userInfo.UID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			Gender:   userInfo.Gender,
		},
	)
	if err != nil {
		return nil, errcode.SESSION_GenJWTTokenErr
	}

	return &dto.SessionToken{
		Token:        token,
		ExpiresIn:    expireIn,
		RefreshToken: refreshToken,
	}, nil
}

func (p *SessionService) Login(ctx context.Context, in *dto.SessionLoginParam) (*dto.SessionToken, error) {
	// get user info by username
	userInfo, err := p.UserDao.GetByUsername(loki.DBEngine, in.Username)
	if err != nil {
		return nil, errcode.SESSION_UserDaoGetByUsernameErr
	}

	if strings.Compare(userInfo.Password, in.Password) != 0 {
		return nil, errcode.SESSION_PasswordIncorrectErr
	}

	return p.genSessionToken(&dto.UserInfo{
		UID:      userInfo.Id,
		Username: userInfo.Username,
		Avatar:   userInfo.Avatar,
		Gender:   userInfo.Gender,
	})
}

func (p *SessionService) RefreshToken(ctx context.Context, in *dto.RefreshTokenParam) (*dto.SessionToken, error) {
	userInfo, err := util.ValidateAuthTokenByUserInfo(in.Token, util.AUTH_TYPE_REFRESH_TOKEN)
	if err != nil {
		return nil, errcode.SESSION_ValidRefreshTokenFailed
	}

	return p.genSessionToken(userInfo)
}
