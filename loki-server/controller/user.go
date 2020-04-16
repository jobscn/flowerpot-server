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

func (p *UserController) SessionLogin(c *gin.Context) {
	ctx := common.ContextWithGin(context.Background(), c)

	in := &dto.UserLoginIn{}
	if c.ShouldBind(in) != nil {
		common.GinUnvalidated(c, nil)
		return
	}

	data, err := p.UserService.Login(ctx, in)
	if err != nil {
		common.GinFailedWithError(c, err)
		return
	}

	c.Header("Authorization", data)
	common.GinSuccess(c, nil)
}

func (p *UserController) AccountRegister(c *gin.Context) {
	ctx := common.ContextWithGin(context.Background(), c)

	in := &dto.UserRegisterIn{}
	if err := c.ShouldBind(in); err != nil {
		common.GinUnvalidated(c, nil)
		return
	}

	err := p.UserService.Register(ctx, in)
	if err != nil {
		common.GinFailedWithError(c, err)
		return
	}

	common.GinSuccess(c, nil)
}
