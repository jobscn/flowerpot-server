package errcode

import "jobscn/ai-flower-pot/loki/common"

var (
	USER_PasswordIncorrectErr    = common.NewError(10000001, "登录账户密码不正确")
	USER_UserDaoGetByUsernameErr = common.NewError(10000002, "检查用户密码数据错误")
	USER_GenJWTTokenErr          = common.NewError(10000003, "生成登录鉴权token失败")
	USER_InsertDBErr             = common.NewError(10000004, "用户注册写数据库失败")
)
