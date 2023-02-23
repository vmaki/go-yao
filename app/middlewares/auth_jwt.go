package middlewares

import (
	"go-yao/common/response"
	"go-yao/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(ctx)
		if err != nil {
			response.BadRequest(ctx, err.Error())
			return
		}

		ctx.Set("current_user_id", claims.UserID)

		ctx.Next()
	}
}
