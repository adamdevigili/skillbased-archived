package db

import (
	"fmt"

	"github.com/adamdevigili/skillbased/api/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

const (
	dbConnRetryLimit = 5
)

// Environment variables to configure target DB. All are required. Will be looked for with the "PG_" prefix
type Config struct {
	Host     string `required:"true"`
	Port     uint16 `required:"true"`
	User     string `required:"true"`
	Database string `required:"true"`
	Default  string `default:"postgres"`
	Password string `required:"true"`
	DevMode  bool   `default:"false"`
	Disabled bool   `default:"false"`
}

// InitDB connects to the Postgres database, and initializes it where required
func InitDB(dbConfig Config) *gorm.DB {
	if dbConfig.Disabled {
		return nil
	}

	baseConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
	)

	log.Info(fmt.Sprintf("%+v", dbConfig))

	// If we're using a development Postgres, disable TLS
	if dbConfig.DevMode {
		baseConnStr += " sslmode=disable"
	}

	connStr := fmt.Sprintf("dbname=%s ", dbConfig.Database) + baseConnStr

	defaultConnStr := fmt.Sprintf("dbname=%s ", dbConfig.Default) + baseConnStr

	var db *gorm.DB

	// Connect to the default "postgres" database first
	log.Infof("Attempting initial connection to database", defaultConnStr)
	db, err := gorm.Open("postgres", defaultConnStr)
	if err != nil {
		log.Fatalf("Unable to connect to database with default settings: %v", err)
	}

	// Create the skillbased database. If the database already exists, this command won't have any effect
	log.Infof("Attempting to create main database", dbConfig.Database)
	db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbConfig.Database))

	// Connect to the "skillbased" database
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to main database: %v", err)
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
	initSportsTable(db)
	deleteAllSeedPlayers(db)
	initPlayersTable(db)
}

func initPlayersTable(db *gorm.DB) {
	log.Info("Populating players database with initial values..")

	db.AutoMigrate(&models.Player{})
	for _, p := range generateSeedPlayers() {
		if err := InsertPlayer(db, p); err != nil {
			log.Warn(err)
		}
	}
}

func deleteAllSeedPlayers(db *gorm.DB) {
	log.Info("deleteing all seed players from existing DB")
	db.Where("is_seed = ?", "t").Delete(&models.Player{})
}

func initSportsTable(db *gorm.DB) {
	log.Info("Populating sports database with initial values..")

	db.AutoMigrate(&models.Sport{})
	for _, s := range initialSports {
		if err := InsertSport(db, &s); err != nil {
			log.Warn(err)
		}
	}
}
