package db

import (
	"database/sql"
	"fetch/management-console/internal/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Db() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_NAME))

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
