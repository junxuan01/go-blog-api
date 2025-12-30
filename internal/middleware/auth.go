package middleware

// AuthMiddleware 认证中间件，验证 JWT Token
import (
	"go-blog-api/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 0
		token := c.GetHeader("Authorization")

		if token == "" {
			code = 401
		} else {
			// 通常 Token 格式为 "Bearer <token>"
			parts := strings.SplitN(token, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				code = 401
			} else {
				claims, err := util.ParseToken(parts[1])
				if err != nil {
					code = 401
				} else {
					// 将当前用户信息存在 Context 中，后续 Handler 可以直接获取
					c.Set("userID", claims.UserID)
					c.Set("username", claims.Username)
				}
			}
		}

		if code != 0 {
			util.Error(c, http.StatusUnauthorized, 401, "Unauthorized")
			c.Abort() // 阻止后续 Handler 执行
			return
		}

		c.Next() // 继续执行后续 Handler
	}
}
