package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 输出访问ip
		authToken := c.GetHeader("auth_token")
		proxy := c.GetHeader("proxy")

		// 检查 auth_token 是否存在，如果不存在则返回错误
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing auth_token,Please add auth_token to request header",
			})
			c.Abort()
			return
		}

		// 在上下文中设置 auth_token
		c.Set("auth_token", authToken)

		// 检查 proxy 是否存在，如果存在则在上下文中设置 proxy
		if proxy != "" {
			c.Set("proxy", proxy)
		} else {
			c.Set("proxy", "")
		}

		// 继续处理请求
		c.Next()
	}
}
