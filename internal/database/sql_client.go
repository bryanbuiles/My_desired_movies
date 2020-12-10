package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //conection to postgres
)

// PostgresSQL struct conection
type PostgresSQL struct {
	*sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "movie_dev"
	password = "movie_dev_pwd"
	dbname   = "movie_wish_db"
)

// NewPostgresSQLClient enable conection with postgres
func NewPostgresSQLClient() *PostgresSQL {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s "+
		"sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &PostgresSQL{db}
}
