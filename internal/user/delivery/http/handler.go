package httphandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/igilgyrg/gin-todo/internal/entity"
	internalError "github.com/igilgyrg/gin-todo/internal/error"
)

type UserHandler struct {
	useCase entity.UserUseCase
}

func NewUserHttpHandler(useCase entity.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (h *UserHandler) Get() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := &entity.User{
			ID:        entity.NewID(),
			Email:     "ig.pomazkov@gmail.com",
			CreatedAt: time.Now(),
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) GetByEmail() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		email, ok := ctx.GetQuery("email")
		if !ok {
			ctx.Error(internalError.NewBadRequestError(nil, "email is required"))
		}
		ctx.JSON(http.StatusOK, email)
	}
}
