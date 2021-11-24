package main

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/db"
	"github.com/adamdevigili/skillbased/api/pkg/server"
	dotenv "github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Port     int    `default:"8080"`
	TokenKey string `envconfig:"TOKEN_KEY"`
	DevMode  bool   `envconfig:"DEV_MODE" default:"false"`
}

func main() {
	// Create new Echo server
	e := echo.New()

	// Load env vars for configuration
	dotenv.Load(".env")

	var dbConfig db.Config
	err := envconfig.Process("db", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	var config Config
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.DevMode {
		log.Info("starting server in development mode")
		dbConfig.DevMode = config.DevMode
	}

	// Initialize DB and routes
	server.InitRoutes(e, db.InitDB(dbConfig))

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
