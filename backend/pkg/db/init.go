package db

import (
	"fmt"
	"os"

	"github.com/adamdevigili/skillbased.io/pkg/models"

	"github.com/jackc/pgx"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost" // "skillbased-pg-postgresql"
	port     = 5432
	user     = "postgres"
	password = "elite360" //"Z17YBnZk9I"
	dbname   = "skillbased"
)

func InitDB(password string) *pgx.ConnPool {
	fmt.Print(password)

	pgxConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Database: dbname,
		},
	}

	connPool, err := pgx.NewConnPool(pgxConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer connPool.Close()

	//psqlInfo := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(fmt.Sprintf("Successfully connected to database '%s' at %s:%d as user '%s'",
		dbname, host, port, user))

	var id string
	var name string

	err = connPool.QueryRow(SelectSportQuery("12345")).Scan(&id, &name)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	fmt.Println(name)
	fmt.Println("Populating database with initial values..")

	//connPool.Exec(InsertSportQuery(s.ID, s.Name))
	for _, s := range models.InitialSports {
		if err := InsertSport(connPool, &s); err != nil {
			panic(err)
		}
	}

	return connPool
}
