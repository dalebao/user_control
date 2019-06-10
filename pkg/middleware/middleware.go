package middleware

import (
	"github.com/dalebao/user_control/pkg/sign"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	验证 guard 参数 是否合法
 */
func GuardMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		guardSet := make(map[string]bool)
		guardSet["wyche"] = true
		guardSet["partner"] = true
		guardSet["xuejia"] = true

		guard := c.PostForm("guard")
		if guardSet[guard] == false {
			c.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  "guard 填写错误",
				"data": "",
			})
			return
		}
		c.Next()
	}
}

func ValidateSignSKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		guard := c.PostForm("guard")
		sKey := c.PostForm("sKey")
		if !sign.ValidateSign(guard,sKey) {
			c.JSON(http.StatusOK, gin.H{
				"code": 422,
				"msg":  "签名 填写错误",
				"data": "",
			})
			return
		}
		c.Next()
	}
}
