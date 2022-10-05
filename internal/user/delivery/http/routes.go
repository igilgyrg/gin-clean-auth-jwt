package httphandler

import (
	"github.com/gin-gonic/gin"
	"github.com/igilgyrg/gin-todo/internal/middleware"
)

func MapUserRoutes(group *gin.RouterGroup, userHandler *UserHandler, middlewareManager *middleware.MiddlewareManager) {
	group.GET("/current", userHandler.Get())
	group.GET("", userHandler.GetByEmail())
}