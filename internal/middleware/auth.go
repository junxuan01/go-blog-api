package middleware

import (
	"strings"

	"go-blog-api/pkg/util"

	"github.com/gin-gonic/gin"
)

// JWT 认证中间件，验证 JWT Token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization Header
		token := c.GetHeader("Authorization")
		if token == "" {
			util.HandleError(c, util.ErrUnauthorized)
			c.Abort()
			return
		}

		// 2. 校验 Token 格式 "Bearer <token>"
		parts := strings.SplitN(token, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.HandleError(c, util.ErrUnauthorized.WithMsg("Token 格式错误"))
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			util.HandleError(c, util.ErrTokenExpired)
			c.Abort()
			return
		}

		// 4. 将用户信息存入 Context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
