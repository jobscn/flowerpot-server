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

func (p *UserController) Register(c *gin.Context) {
	ctx := common.ContextWithGin(context.Background(), c)

	in := &dto.UserRegisterParam{}
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
