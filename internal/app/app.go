package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/igilgyrg/gin-todo/pkg/logging"
	mongoClient "github.com/igilgyrg/gin-todo/pkg/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const ctxTimeout = 5
const timeoutServer = 10

type App struct {
	cfg     *config
	mongoDB *mongo.Database
	router  *gin.Engine
}

func New() *App {
	cfg := newConfig()
	mongoConfig := mongoClient.NewMongoConfig(cfg.MongoHost, cfg.MongoPort, cfg.MongoDatabase, cfg.MongoUsername, cfg.MongoPassword)

	mongoClient, err := mongoClient.Init(mongoConfig)
	if err != nil {
		log.Fatal("mongo database have not initialized")
	}

	return &App{
		cfg:     cfg,
		mongoDB: mongoClient,
		router:  gin.Default(),
	}
}

func (a *App) StartServer() {
	if err := a.mapHandlers(); err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:           ":3000",
		Handler:        a.router,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   timeoutServer,
		BaseContext: func(listener net.Listener) context.Context {
			ctx, shutdown := context.WithTimeout(logging.ContextWithLogger(context.Background()), ctxTimeout*time.Second)
			defer shutdown()
			return ctx
		},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
