package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_shopping/pkg/utils"
	"net/http"
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
	Myclaims := utils.ParseToken(tokenStr, utils.MySecret)
	if Myclaims == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": "鉴权失败,claims获取失败",
		})
		c.Abort()
		return
	}
	c.Request = c.Request.WithContext(context.WithValue(c, "userid", Myclaims.UserID))
	c.Next()
}
