package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	internalError "github.com/igilgyrg/gin-todo/internal/error"
)

func (mw *MiddlewareManager) ResponseErrorToJSON() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()

		if len(ctx.Errors) == 0 {
			ctx.Status(http.StatusOK)
			return
		}

		for i := range ctx.Errors {
			var appError *internalError.HttpError
			if errors.As(ctx.Errors[i], &appError) {
				ctx.Writer.WriteHeader(appError.Status())
				ctx.Writer.Write(appError.Marshal())
				continue
			}

			ctx.Writer.WriteHeader(http.StatusInternalServerError)
			errBytes, _ := json.Marshal(ctx.Errors[i])
			ctx.Writer.Write(errBytes)
		}
	}
}
