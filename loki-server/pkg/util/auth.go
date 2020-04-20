package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"time"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

const (
	AUTH_TYPE_TOKEN                    = 0
	AUTH_SECRET_TOKEN                  = "bf3a174b-fa87-48c2-941e-50dd2d336b6a"
	AUTH_SECRET_TOKEN_EXPIRE_IN        = 2 * time.Hour
	AUTH_SECRET_TOKEN_EXPIRE_IN_SECOND = 2 * 60 * 60

	AUTH_TYPE_REFRESH_TOKEN                    = 1
	AUTH_SECRET_REFRESH_TOKEN                  = "5cfb6797-483c-4cfb-a374-57b65e0e89bb"
	AUTH_SECRET_REFRESH_TOKEN_EXPIRE_IN        = 15 * 24 * time.Hour
	AUTH_SECRET_REFRESH_TOKEN_EXPIRE_IN_SECOND = 15 * 24 * 60 * 60
)

type JWTPayload struct {
	dto.UserInfo
	jwt.StandardClaims
}

func getParamByAuthType(authType int64) (string, time.Duration, int) {
	if authType == AUTH_TYPE_REFRESH_TOKEN {
		return AUTH_SECRET_REFRESH_TOKEN, AUTH_SECRET_REFRESH_TOKEN_EXPIRE_IN, AUTH_SECRET_TOKEN_EXPIRE_IN_SECOND
	}

	return AUTH_SECRET_TOKEN, AUTH_SECRET_TOKEN_EXPIRE_IN, AUTH_SECRET_REFRESH_TOKEN_EXPIRE_IN_SECOND
}

func GenAuthTokenByUserInfo(authType int64, userInfo *dto.UserInfo) (string, int, error) {
	secret, timeAppend, expireIn := getParamByAuthType(authType)

	t := time.Now()
	var payload = &JWTPayload{
		UserInfo: dto.UserInfo{
			UID:      userInfo.UID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			Gender:   userInfo.Gender,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Add(timeAppend).Unix(),
			IssuedAt:  t.Unix(),
			NotBefore: t.Unix(),
		},
	}

	j := jwt.NewWithClaims(JWT_SIGNING_METHOD, payload)

	token, err := j.SignedString([]byte(secret))

	return "Bearer " + token, expireIn, err
}

func ValidateAuthTokenByUserInfo(t string, authType int64) (*dto.UserInfo, error) {
	secret, _, _ := getParamByAuthType(authType)

	token, err := jwt.ParseWithClaims(t, &JWTPayload{}, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		if payload, ok := token.Claims.(*JWTPayload); ok && payload != nil {
			return &dto.UserInfo{
				UID:      payload.UID,
				Username: payload.Username,
				Avatar:   payload.Avatar,
				Gender:   payload.Gender,
			}, nil
		}
	}

	return nil, fmt.Errorf("Authorization Token Failed)")
}
