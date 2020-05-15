package main

import (
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/server"
	"github.com/labstack/echo/v4"
)

func main() {
	dbConn := db.InitDB()

	e := echo.New()
	server.InitRoutes(e, dbConn)

	// Serve static files for non-operational API calls (browser visits to api subdomain
	e.Static("/", "./static")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
