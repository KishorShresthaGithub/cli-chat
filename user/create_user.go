package user

import (
  "io"
	"database/sql"
  "encoding/json"
	_ "github.com/mattn/go-sqlite3"
	shared "kslabs/chat-app-cli/shared"
	"log"
  "fmt"
)

type CreateUserEntity struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserInterface interface {
	Controller(u *CreateUserEntity, w io.Writer)
}

type CreateUser struct {
	db *sql.DB
  writer *io.Writer
}


func (user *CreateUser) SetWriter(write io.Writer)  {
  user.writer = &write
}

func (user *CreateUser) Presentation(u *CreateUserEntity)  {
  fmt.Println(u)
  json.NewEncoder(*user.writer).Encode(&u)
}

func (user *CreateUser) Controller(u *CreateUserEntity, w io.Writer)  {
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
