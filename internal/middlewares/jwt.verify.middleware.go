package middlewares

import (
	"strings"

	"github.com/anle/codebase/internal/utils/token"
	"github.com/anle/codebase/response"
	"github.com/gin-gonic/gin"
)

var (
	headerName = "Authorization"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.GetHeader(headerName)
		if headerValue == "" {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		arrayHeaderValues := strings.Split(headerValue, " ")
		if len(arrayHeaderValues) != 2 || arrayHeaderValues[0] != "Bearer" {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		accessToken := arrayHeaderValues[1]
		userID, err := token.VerifyJWTToken(accessToken)
		if err != nil {
			response.ErrorResponseNotAuthorize(c, response.ErrCodeNotAuthorize, nil)
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
