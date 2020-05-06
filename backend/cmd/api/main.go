package main

import (
	"fmt"
	"os"

	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/server"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println(os.Args[1])

	dbPassword := os.Args[1]
	dbConn := db.InitDB(dbPassword)

	e := echo.New()

	server.InitRoutes(e, dbConn)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
