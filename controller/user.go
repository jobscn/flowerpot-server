package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/dao/impl"
	"jobscn/ai-flower-pot/model/do"
	"jobscn/ai-flower-pot/model/dto"
	"jobscn/ai-flower-pot/service"
	"jobscn/ai-flower-pot/ymer"
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
	userDao := &dao_impl.UserDao{}
	err := userDao.Insert(ymer.DBEngine, &do.TUser{
		Username:   "test1",
		Password:   fmt.Sprintf("%x", md5.Sum([]byte("123456"))),
		Phone:      "13358292565",
		Avatar:     "",
		Gender:     0,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"ok": "something right!",
	})
}