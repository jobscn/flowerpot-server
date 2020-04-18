package common

var LokiCode = &struct {
	SUCCESS      Response
	FAILED       Response
	UNVALIDATED  Response
	UNAUTHORIZED Response
	FORBIDDEN    Response
}{
	SUCCESS: Response{
		Code:    200,
		Message: "请求成功",
	},
	FAILED: Response{
		Code:    500,
		Message: "请求失败",
	},
	UNVALIDATED: Response{
		Code:    400,
		Message: "请求参数不正确",
	},
	UNAUTHORIZED: Response{
		Code:    401,
		Message: "用户未登录",
	},
	FORBIDDEN: Response{
		Code:    403,
		Message: "没有权限访问",
	},
}
