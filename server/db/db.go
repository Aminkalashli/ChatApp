package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/chatApp?sslmode=disalbe")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil

}

func (d *Database) close() {
	d.db.Close()
}

func (d *Database) getDB() *sql.DB {
	return d.db
}
