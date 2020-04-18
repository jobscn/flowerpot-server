package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/service"
	"jobscn/ai-flower-pot/loki/common"
)

type SessionController struct {
	SessionService service.ISessionService `inject:""`
}

func (p *SessionController) Login(c *gin.Context) {
	ctx := common.ContextWithGin(context.Background(), c)

	in := &dto.SessionLoginParam{}
	if c.ShouldBind(in) != nil {
		common.GinUnvalidated(c, nil)
		return
	}

	data, err := p.SessionService.Login(ctx, in)
	if err != nil {
		common.GinFailedWithError(c, err)
		return
	}

	common.GinSuccess(c, data)
}

func (p *SessionController) RefreshToken(c *gin.Context) {

}
