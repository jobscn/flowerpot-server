package errcode

import "jobscn/ai-flower-pot/loki/common"

var (
	SESSION_PasswordIncorrectErr    = common.NewError(20000001, "登录账户密码不正确")
	SESSION_UserDaoGetByUsernameErr = common.NewError(20000002, "检查用户密码数据错误")
	SESSION_GenJWTTokenErr          = common.NewError(20000003, "生成登录鉴权token失败")
	SESSION_ValidRefreshTokenFailed = common.NewError(20000004, "RefreshToken不正确或已过期")
)
