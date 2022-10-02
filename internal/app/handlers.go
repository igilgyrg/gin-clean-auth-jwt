package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/igilgyrg/gin-todo/internal/middleware"
	userHttp "github.com/igilgyrg/gin-todo/internal/user/delivery/http"
	userRepository "github.com/igilgyrg/gin-todo/internal/user/repository"
	userUseCase "github.com/igilgyrg/gin-todo/internal/user/usecase"
	mongoClient "github.com/igilgyrg/gin-todo/pkg/repository/mongo"
)

func (a *App) mapHandlers() error {
	mongoConfig := mongoClient.NewMongoConfig("localhost", "27017", "todo", "root", "12345")
	mongoDatabase, err := mongoClient.Init(mongoConfig)
	if err != nil {
		log.Fatal("error of connect to mongo database")
	}

	a.router.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		ctx.Next()
	})

	middlewareManager := middleware.NewMiddlewareManager()

	v1 := a.router.Group("/api/v1")
	v1.Use(middlewareManager.Logger())

	userGroup := v1.Group("/users")

	userRepo := userRepository.NewUserMongoRepository(mongoDatabase)
	userUseCase := userUseCase.NewUserCase(userRepo)
	userHandler := userHttp.NewUserHttpHandler(userUseCase)

	userHttp.MapUserRoutes(userGroup, userHandler, middlewareManager)

	return nil
}
