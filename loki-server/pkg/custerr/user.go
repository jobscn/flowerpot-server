package custerr

import "jobscn/ai-flower-pot/loki/common"

var (
	UserLogin_PasswordIncorrectErr = common.NewError(10000001, "登录账户密码不正确")
)
