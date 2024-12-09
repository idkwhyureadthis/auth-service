package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	conn *sql.DB
}

func New(connUrl string) *DB {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{conn: conn}
}

func setupMigrations() {

}
