package httphandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/igilgyrg/gin-todo/internal/entity"
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
		time.Sleep(6 * time.Second)
		ctx.JSON(http.StatusOK, user)
	}
}
