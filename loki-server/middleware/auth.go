package middleware

import (
	"github.com/gin-gonic/gin"
	"jobscn/ai-flower-pot/loki-server/pkg/util"
	"jobscn/ai-flower-pot/loki/common"
	"strings"
)

func JwtAuthMiddleware(c *gin.Context) {
	auth := c.GetHeader("authorization")

	splits := strings.Split(auth, " ")

	if len(splits) == 0 || strings.Compare(splits[0], "Bearer") != 0 {
		common.GinUnauthorized(c, nil)
		c.Abort()
		return
	}

	userInfo, err := util.ValidateAuthTokenByUserInfo(splits[1], util.AUTH_TYPE_TOKEN)
	if err != nil {
		common.GinUnauthorized(c, nil)
		c.Abort()
		return
	}

	c.Set("user", userInfo)

	c.Next()
}
