package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/igilgyrg/gin-todo/internal/middleware"
	userHttp "github.com/igilgyrg/gin-todo/internal/user/delivery/http"
	userRepository "github.com/igilgyrg/gin-todo/internal/user/repository"
	userUseCase "github.com/igilgyrg/gin-todo/internal/user/usecase"
)

func (a *App) mapHandlers() error {
	a.router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	middlewareManager := middleware.NewMiddlewareManager()

	v1 := a.router.Group("/api/v1")
	v1.Use(middlewareManager.Logger(), middlewareManager.ResponseErrorToJSON())

	userGroup := v1.Group("/users")

	userRepo := userRepository.NewUserMongoRepository(a.mongoDB)
	userUseCase := userUseCase.NewUserCase(userRepo)
	userHandler := userHttp.NewUserHttpHandler(userUseCase)

	userHttp.MapUserRoutes(userGroup, userHandler, middlewareManager)

	return nil
}
