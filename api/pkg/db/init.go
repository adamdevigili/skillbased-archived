package db

import (
	"crypto/tls"
	"time"

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
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Host:     dbConfig.Host,
			Port:     dbConfig.Port,
			User:     dbConfig.User,
			Password: dbConfig.Password,
			Database: dbConfig.Database,
		},
	}

	log.Infof("Attempting to connect to database '%s' at %s:%d as user '%s'",
		dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User,
	)

	var connPool *pgx.ConnPool
	for i := 1; i <= dbConnRetryLimit; i++ {
		connPool, err = pgx.NewConnPool(pgxConfig)
		if err != nil {
			log.Warnf("Unable to connect to database %v. Retries remaining: %d", err, dbConnRetryLimit-i)
		} else {
			break
		}

		if i == dbConnRetryLimit {
			log.Fatal("Maximum number of retries to establish database connection reached, exiting")
		}

		time.Sleep(15 * time.Second)
	}

	if connPool != nil {
		log.Infof("Successfully connected to database '%s' at %s:%d as user '%s'",
			dbConfig.Database, dbConfig.Host, dbConfig.Port, dbConfig.User,
		)
	} else {
		log.Fatal("Could not connect to database, exiting")
	}

	//db, err := gorm.Open(
	//	"postgres",
	//	fmt.Sprintf(
	//		"host=%s port=%s user=%s dbname=%s password=%s",
	//		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password,
	//	),
	//)
	//defer db.Close()

	CreateSportsTable(connPool)

	log.Info("Populating sports database with initial values..")
	for _, s := range models.InitialSports {
		if err := InsertSport(connPool, &s); err != nil {
			log.Warn(err)
		}
	}

	return connPool
}
