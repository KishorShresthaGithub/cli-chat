package shared 

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
  "fmt"
)

func GetDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./chat.db")

	if err != nil {
    fmt.Println(err)
		log.Fatal(err)
	}

	return db
}
