package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/jackc/pgx"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload" // Import autoload package to setup env vars for dotenv
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

const (
	dbConnRetryLimit = 5
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

	var connPool *pgx.ConnPool
	for i := 0; i <= dbConnRetryLimit; i++ {
		connPool, err = pgx.NewConnPool(pgxConfig)
		if err != nil {
			log.Warnf("Unable to connect to database", err)
		}
		if i == dbConnRetryLimit {
			log.Panic("Maximum number of retries to establish database connection reached")
		}
		time.Sleep(time.Duration(1 * time.Minute))
	}

	log.Info(fmt.Sprintf(
		"successfully connected to database '%s' at %s:%d as user '%s'",
		dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User,
	))

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password,
		),
	)
	defer db.Close()

	CreateSportsTable(connPool)

	log.Info("populating sports database with initial values..")
	for _, s := range models.InitialSports {
		if err := InsertSport(connPool, &s); err != nil {
			log.Warn(err)
		}
	}

	return connPool
}
