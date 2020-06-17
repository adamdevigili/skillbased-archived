package main

import (
	"github.com/adamdevigili/skillbased.io/pkg/db"
	"github.com/adamdevigili/skillbased.io/pkg/server"
	"github.com/labstack/echo/v4"
)

func main() {
	// Create new Echo server
	e := echo.New()

	// Initialize DB and routes
	server.InitRoutes(e, db.InitDB())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
