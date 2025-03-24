package middlewares

import (
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

func RoleVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
