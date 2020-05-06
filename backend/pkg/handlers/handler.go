package handlers

import "github.com/jackc/pgx"

type (
	Handler struct {
		DBConn *pgx.ConnPool
	}
)
