package middleware

import (
	"fmt"
	errors "intelligent-investor/internal/pkg/errors"
	"intelligent-investor/internal/pkg/response"
	"intelligent-investor/internal/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

var IgnorePaths = []string{}

func IgnorePathsInit(paths []string) {
	fmt.Println("IgnorePaths:", paths)
	IgnorePaths = paths
}

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查路径是否在忽略列表中
		for _, path := range IgnorePaths {
			if c.Request.URL.Path == path {
				c.Next()
				return
			}
			if strings.HasSuffix(path, "/*") {
				if strings.HasPrefix(c.Request.URL.Path, path[:len(path)-2]) {
					c.Next()
					return
				}
			}
		}
		// 从请求头中获取 Authorization 字段
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(response.CodeFailed, errors.ErrAuthorizationFailed.WithMessage("Authorization header is required"))
			c.Abort()
			return
		}
		// 验证 JWT 令牌
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := token.ParseToken(tokenString)
		if err != nil {
			c.JSON(response.CodeFailed, errors.ErrAuthorizationFailed.WithMessage("Invalid token"))
			c.Abort()
			return
		}
		// 验证通过，将用户名设置到上下文
		c.Next()
	}
}
