package db

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/adamdevigili/skillbased.io/pkg/models"

	"github.com/jackc/pgx"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // "skillbased-pg-postgresql"
	port     = 5432
	user     = "postgres"
	password = "elite360" //"Z17YBnZk9I"
	dbname   = "skillbased"
)

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

	fmt.Println(fmt.Sprintf("%+v", dbConfig))

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

	log.Info(fmt.Sprintf("Successfully connected to database '%s' at %s:%d as user '%s'",
		dbname, host, port, user))

	log.Info("Populating database with initial values..")

	//connPool.Exec(InsertSportQuery(s.ID, s.Name))
	for _, s := range models.InitialSports {
		if err := InsertSport(connPool, &s); err != nil {
			panic(err)
		}
	}

	return connPool
}
