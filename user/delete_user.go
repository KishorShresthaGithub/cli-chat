package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	shared "kslabs/chat-app-cli/shared"
	"log"
)

type DeleteUserInterface interface {
	Controller(id int, w http.ResponseWriter)
}

type DeleteUser struct {
	db     *sql.DB
	writer *http.ResponseWriter
}

func (user *DeleteUser) SetWriter(w *http.ResponseWriter) {
	user.writer = w
}

func (user *DeleteUser) Controller(id int, w http.ResponseWriter) {
	user.SetWriter(&w)
	user.DataAccess(id)
	user.Presentation()
}

func (user *DeleteUser) DataAccess(id int) {
	user.db = shared.GetDbConnection()
	defer user.db.Close()

	tx, err := user.db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Delete user %d", id)
}

func (user *DeleteUser) Presentation() {

	response := &shared.Response{
		Code:    200,
		Message: "User has been deleted"}

	json.NewEncoder(*user.writer).Encode(response)
}
