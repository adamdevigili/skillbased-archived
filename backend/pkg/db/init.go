package db

import (
	"fmt"
	"os"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/jackc/pgx"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

// Environment variables to configure target DB. All are required. Will be looked for with the "PG_" prefix
type dbConfig struct {
	Host     string `required:"true"`
	Port     uint16 `required:"true"`
	User     string `required:"true"`
	Database string `required:"true"`
	Password string `required:"true"`
}

func InitDB() *pgx.ConnPool {
	var dbConfig dbConfig

	err := envconfig.Process("pg", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	pgxConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			Database: dbConfig.Database,
		},
	}

	connPool, err := pgx.NewConnPool(pgxConfig)
	if err != nil {
		log.Error("Unable to connect to database", err)
		os.Exit(1)
	}

	log.Info(fmt.Sprintf(
		"successfully connected to database '%s' at %s:%d as user '%s'",
		dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User,
	))

	CreateSportsTable(connPool)

	log.Info("populating sports database with initial values..")
	for _, s := range models.InitialSports {
		if err := InsertSport(connPool, &s); err != nil {
			log.Warn(err)
		}
	}

	return connPool
}
