package db

import (
	"fmt"
	"time"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	dotenv "github.com/joho/godotenv"
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
	DevMode  bool   `required:"true"`
}

// InitDB connects to the Postgres database, and initializes it where required
func InitDB() *gorm.DB {
	var dbConfig dbConfig

	dotenv.Load("../.env")
	err := envconfig.Process("pg", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Database,
		dbConfig.Password,
	)

	// If we're using a development Postgres, disable TLS
	if dbConfig.DevMode {
		connStr += " sslmode=disable"
	}

	log.Infof("Attempting to connect to database '%s' at %s:%d as user '%s'. DevMode=%t",
		dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DevMode,
	)

	var db *gorm.DB
	for tries := 1; tries <= dbConnRetryLimit; tries++ {
		db, err = gorm.Open("postgres", connStr)
		if err != nil {
			log.Warnf("Unable to connect to database %v. Retries remaining: %d", err, dbConnRetryLimit-tries)
		} else {
			break
		}

		if tries == dbConnRetryLimit {
			log.Fatal("Maximum number of retries to establish database connection reached, exiting")
		}

		time.Sleep(15 * time.Second)
	}

	if db != nil {
		log.Infof("Successfully connected to database '%s' at %s:%d as user '%s'",
			dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User,
		)
	} else {
		log.Fatal("Could not connect to database, exiting")
	}

	// Initialize the required tables
	initTables(db)

	return db
}

func initTables(db *gorm.DB) {
	//initSportsTable(db)
}

func initSportsTable(db *gorm.DB) {
	log.Info("Populating sports database with initial values..")

	db.AutoMigrate(&models.Sport{})
	for _, s := range models.InitialSports {
		if err := InsertSport(db, &s); err != nil {
			log.Warn(err)
		}
	}
}
