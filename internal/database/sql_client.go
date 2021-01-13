package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //conection to postgres, The _ starts the init
)

// PostgresSQL struct conection
type PostgresSQL struct {
	*sql.DB
}

const (
	host = "localhost"
	//host     = "database"
	port     = 5432
	user     = "movie_dev"
	password = "movie_dev_pwd"
	dbname   = "movie_dev"
)

// NewPostgresSQLClient enable conection with postgres
func NewPostgresSQLClient() *PostgresSQL {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
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
