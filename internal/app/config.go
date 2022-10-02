package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type config struct {
	Port                  string `env:"PORT"`
	MongoPort             string `env:"MONGO_PORT"`
	MongoHost             string `env:"MONGO_HOST"`
	MongoDatabase         string `env:"MONGO_DATABASE"`
	MongoUsername         string `env:"MONGO_USERNAME"`
	MongoPassword         string `env:"MONGO_PASSWORD"`
	AccessTokenSignature  string `env:"ACCESS_TOKEN_SIGNATURE"`
	RefreshTokenSignature string `env:"REFRESH_TOKEN_SIGNATURE"`
}

var instance *config
var once sync.Once

func newConfig() *config {
	once.Do(configEnv)
	return instance
}

func configEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	instance = &config{}
	if err := env.Parse(instance); err != nil {
		log.Fatal(err)
	}
}