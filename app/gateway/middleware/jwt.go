package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_shopping/pkg/utils"
	"net/http"
	"time"
)

func MiddleJWT(c *gin.Context) {
	var code uint
	code = http.StatusOK

	// 获取token
	tokenStr := c.GetHeader("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "鉴权失败,token为空",
		})
		c.Abort()
		return
	}
	// 获取claims
	Myclaims, err := utils.ParseToken(tokenStr, utils.MySecret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "鉴权失败,claims获取失败",
		})
		c.Abort()
		return
	}
	// 验证失效时间
	if time.Now().Unix() > Myclaims.ExpiresAt {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "鉴权失败,时间过期,请重新登录",
		})
		c.Abort()
		return
	}

	c.Request = c.Request.WithContext(context.WithValue(c, "userid", Myclaims.UserID))
	c.Next()
}
