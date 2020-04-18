package common

import (
	"context"
	"github.com/gin-gonic/gin"
)

// context处理
// context.Context用的是指针this，因此context.Context一定是指针传值，这里效率不会低
// TODO: >> gin.Context也是指针传值，理论上这里不需要再使用指针，研究下
func ContextWithGin(ctx context.Context, ginContext *gin.Context) context.Context {
	return context.WithValue(ctx, "gin-context", ginContext)
}

func GinSuccess(c *gin.Context, Data interface{}) {
	c.JSON(LokiCode.SUCCESS.GetCode(), Response{
		Code:    LokiCode.SUCCESS.GetCode(),
		Message: LokiCode.SUCCESS.GetMessage(),
		Data:    Data,
	})
}

func GinFailed(c *gin.Context) {
	c.JSON(LokiCode.FAILED.GetCode(), Response{
		Code:    LokiCode.FAILED.GetCode(),
		Message: LokiCode.FAILED.GetMessage(),
	})
}

func GinFailedWithError(c *gin.Context, err error) {
	_err, failed := err.(Error)
	if failed != true {
		c.JSON(LokiCode.FAILED.GetCode(), Response{
			Code:    LokiCode.FAILED.GetCode(),
			Message: "错误不能被正确解析为common.Error类型: " + err.Error(),
		})
		return
	}

	c.JSON(LokiCode.FAILED.GetCode(), Response{
		Code:    _err.GetCode(),
		Message: _err.GetMessage(),
	})
}

func GinFailedWithData(c *gin.Context, Data interface{}) {
	c.JSON(LokiCode.FAILED.GetCode(), Response{
		Code:    LokiCode.FAILED.GetCode(),
		Message: LokiCode.FAILED.GetMessage(),
		Data:    Data,
	})
}

func GinForbidden(c *gin.Context, Data interface{}) {
	c.JSON(LokiCode.FORBIDDEN.GetCode(), Response{
		Code:    LokiCode.FORBIDDEN.GetCode(),
		Message: LokiCode.FORBIDDEN.GetMessage(),
		Data:    Data,
	})
}

func GinUnauthorized(c *gin.Context, Data interface{}) {
	c.JSON(LokiCode.UNAUTHORIZED.GetCode(), Response{
		Code:    LokiCode.UNAUTHORIZED.GetCode(),
		Message: LokiCode.UNAUTHORIZED.GetMessage(),
		Data:    Data,
	})
}

func GinUnvalidated(c *gin.Context, Data interface{}) {
	c.JSON(LokiCode.UNVALIDATED.GetCode(), Response{
		Code:    LokiCode.UNVALIDATED.GetCode(),
		Message: LokiCode.UNVALIDATED.GetMessage(),
		Data:    Data,
	})
}
