package middleware

import "github.com/gin-gonic/gin"

func (mw *MiddlewareManager) AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			ctx.AbortWithStatusJSON(401, "Bearer token is required")
		}
	}
}