package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseOptions struct {
	Host   string
	Port   string
	User   string
	Pass   string
	Dbname string
}

var database *sql.DB

func GetDB() *sql.DB {
	return database
}

func InitDB(cfg DatabaseOptions) error {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	database = db

	return nil
}
