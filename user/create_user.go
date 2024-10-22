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

type CreateUserEntity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserInterface interface {
	Controller(u *CreateUserEntity, w http.ResponseWriter)
}

type CreateUser struct {
	db     *sql.DB
	writer *http.ResponseWriter
}

func (user *CreateUser) SetWriter(write http.ResponseWriter) {
	user.writer = &write
}

func (user *CreateUser) Presentation(u *CreateUserEntity) {
	fmt.Println(u)
	json.NewEncoder(*user.writer).Encode(&u)
}

func (user *CreateUser) Controller(u *CreateUserEntity, w http.ResponseWriter) {
	user.SetWriter(w)
	user.DataAccess(u)
	user.Presentation(u)
}

func (user *CreateUser) DataAccess(u *CreateUserEntity) {

	user.db = shared.GetDbConnection()

	tx, err := user.db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO users (name,email) values (?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email)

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	user.db.Close()

}
