package errcode

import "jobscn/ai-flower-pot/loki/common"

var (
	USER_InsertDBErr = common.NewError(10000001, "用户注册写数据库失败")
)
