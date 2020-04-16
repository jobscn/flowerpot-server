package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"time"
)

var JWTSigningMethod = jwt.SigningMethodHS256

const JWTSecret = "I3hJdWVLaSk5uchI"

type JWTPayload struct {
	dto.UserInfo
	jwt.StandardClaims
}

func GenAuthTokenByUserInfo(userInfo *dto.UserInfo) (string, error) {
	t := time.Now()
	var payload = &JWTPayload{
		UserInfo: dto.UserInfo{
			UID:      userInfo.UID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			Gender:   userInfo.Gender,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Add(time.Hour).Unix(),
			IssuedAt:  t.Unix(),
			NotBefore: t.Unix(),
		},
	}

	j := jwt.NewWithClaims(JWTSigningMethod, payload)

	token, err := j.SignedString([]byte(JWTSecret))

	return "Bearer " + token, err
}

func ValidateAuthTokenByUserInfo(t string) (*dto.UserInfo, error) {
	token, err := jwt.ParseWithClaims(t, &JWTPayload{}, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(token.Valid)
	if token.Valid {
		if payload, ok := token.Claims.(*JWTPayload); ok {
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
