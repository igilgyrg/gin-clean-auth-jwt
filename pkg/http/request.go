package http

import "github.com/gin-gonic/gin"

func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.BindJSON(request); err != nil {
		return err
	}

	return ValidateStruct(ctx.Request.Context(), request)
}