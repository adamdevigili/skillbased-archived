package db

import (
	"fmt"

	"github.com/jackc/pgx"

	"github.com/adamdevigili/skillbased.io/pkg/models"
)

func InsertSport(conn *pgx.ConnPool, sport *models.Sport) error {
	if _, err := conn.Exec(InsertSportQuery(sport.ID, sport.Name)); err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("succesfully stored sport '%s' in the database", sport.Name))

	return nil
}

func ConvertSportToRow(sport *models.Sport) {

}

func SelectSportQuery(id string) string {
	return fmt.Sprintf("SELECT * FROM sports WHERE id='%s';", id)
}

func InsertSportQuery(id, name string) (string, string, string) {
	return `insert into sports(id, name) values ($1, $2)`, id, name
}
