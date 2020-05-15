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

/* Create/Insert */

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

/* List/Get */

func GetSport(conn *pgx.ConnPool, id string) (*models.Sport, error) {
	row := conn.QueryRow(GetSportQuery(id))

	var (
		name string
	)

	err := row.Scan(&id, &name)
	if err != nil {
		log.Error(fmt.Sprintf("sport with id '%s' does not exist in the database", id))
		return nil, err
	}

	return &models.Sport{ID: id, Name: name}, nil

}

func GetSportQuery(id string) (string, string) {
	return fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, sportTableName), id
}

/* List/Get */

func ListSports(conn *pgx.ConnPool) (*models.SportList, error) {
	rows, err := conn.Query(ListSportQuery())
	if err != nil {
		return nil, err
	}

	sportList := &models.SportList{}

	var (
		id   string
		name string
	)

	for rows.Next() {
		rows.Scan(&id, &name)

		sportList.Items = append(sportList.Items, models.Sport{
			ID:   id,
			Name: name,
		})

		sportList.NumItems += 1
	}

	return sportList, nil

}

func ListSportQuery() string {
	return fmt.Sprintf(`SELECT * FROM %s`, sportTableName)
}

/* Update */

/* Delete */
