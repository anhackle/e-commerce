package middlewares

import (
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

func RoleVerifyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, err := c.Value("role").(string)
		if !err {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		if role != "admin" {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
