package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/service"
	"jobscn/ai-flower-pot/loki/common"
)

type UserController struct {
	UserService service.IUserService `inject:""`
}

func (p *UserController) Login(c *gin.Context) {
	ctx := common.ContextWithGin(context.Background(), c)

	userLoginIn := &dto.UserLoginIn{}
	err := c.BindQuery(userLoginIn)
	if err != nil {
		common.GinUnvalidated(c, nil)
		return
	}

	data, err := p.UserService.Login(ctx, userLoginIn)
	if err != nil {
		common.GinFailedWithError(c, err)
		return
	}

	common.GinSuccess(c, data)
}

func (p *UserController) Register(c *gin.Context) {
}
