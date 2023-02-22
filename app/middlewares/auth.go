package middlewares

import (
	"go-yao/pkg/jwt"
	"go-yao/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.BadRequest(c, err.Error())
			return
		}

		c.Set("current_user_id", claims.UserID)
		c.Next()
	}
}
