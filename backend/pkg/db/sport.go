package db

import (
	"fmt"
	"strings"

	"github.com/adamdevigili/skillbased.io/pkg/models"
	"github.com/jackc/pgx"
	"github.com/labstack/gommon/log"
)

const (
	sportTableName = "sports"
)

const (
	duplicateKeyError   = "duplicate key value violates unique constraint"
	duplicateTableError = "already exists"
)

func CreateSportsTable(conn *pgx.ConnPool) error {
	if _, err := conn.Exec(
		fmt.Sprintf(`CREATE TABLE %s(id VARCHAR(50) UNIQUE PRIMARY KEY, name VARCHAR(50));`, sportTableName),
	); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Info("sports table already exists")
			return nil
		} else {
			return err
		}
	}

	log.Info(fmt.Sprintf("successfully created the %s table", sportTableName))

	return nil
}

func InsertSport(conn *pgx.ConnPool, sport *models.Sport) error {
	if _, err := conn.Exec(InsertSportQuery(sport.ID, sport.Name)); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("succesfully stored sport '%s' in the database", sport.Name))

	return nil
}

func InsertSportQuery(id, name string) (string, string, string) {
	return fmt.Sprintf(`INSERT INTO %s(id, name) VALUES ($1, $2)`, sportTableName), id, name
}

func ConvertSportToRow(sport *models.Sport) {

}
