package controller

import (
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki-server/model/dto"
	"jobscn/ai-flower-pot/loki-server/service"
)

type UserController struct {
	service.IUserService
}

func (p *UserController) Login(c *gin.Context)  {
	loginInfo := &dto.UserLoginIn{}

	err := c.BindJSON(loginInfo)
	if err != nil {
		c.JSON(200, Response{
			Code:    10000001,
			Message: "参数错误",
		})
		return
	}



	c.JSON(200, Response{
		Code:    0,
		Message: "123213",
		Data:    nil,
	})
}

func (p *UserController) Register(c *gin.Context) {


	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": "something right!",
	})
}