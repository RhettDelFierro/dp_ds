package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlHandler struct {
	*SQLHandler
}

func NewMySqlHandler(connection string) (*MySqlHandler, error) {
	db, err := sql.Open("postgres", connection)
	return &MySqlHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
