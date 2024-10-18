package shared 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./chat.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
