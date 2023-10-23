package postgres

//nolint:revive
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func MustConnect(uri string) *sql.DB {
	db, err := Connect(uri)
	if err != nil {
		panic("can't connect to postgres: " + err.Error())
	}

	return db
}

func Connect(uri string) (*sql.DB, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
