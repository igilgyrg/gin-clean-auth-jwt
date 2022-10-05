package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	mongoClient "github.com/igilgyrg/gin-todo/pkg/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

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
		Addr:           fmt.Sprintf(":%s", a.cfg.Port),
		Handler:        a.router,
		MaxHeaderBytes: 1 << 10,
		WriteTimeout:   timeoutServer * time.Second,
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
