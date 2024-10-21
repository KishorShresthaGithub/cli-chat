package user

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io"
	shared "kslabs/chat-app-cli/shared"
	"log"
  "fmt"
)

type GetUserEntity struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type GetUserDb struct {
	id    int
	name  string
	email string
}

type GetUserInterface interface {
	GetAllController(w io.Writer)
	GetSingleController(id int, w io.Writer)
}

type GetUser struct {
	db     *sql.DB
	writer *io.Writer
}

func (user *GetUser) SetWriter(writer io.Writer) {
	user.writer = &writer
}

func (user *GetUser) GetSingleController(id int, writer io.Writer) {
	u := user.DataAccessSingle(id)
	user.SinglePresentation(u, &writer)
}

func (user *GetUser) SinglePresentation(u *GetUserEntity, writer *io.Writer) {
	json.NewEncoder(*writer).Encode(u)
}

func (user *GetUser) Presentation(u []*GetUserEntity, writer *io.Writer) {
	json.NewEncoder(*writer).Encode(u)
}

func (user *GetUser) GetAllController(w io.Writer) {
	u := user.DataAccess()
  fmt.Println(u)
	user.Presentation(u, &w)
}

func (user *GetUser) DataAccessSingle(id int) *GetUserEntity {
	user.db = shared.GetDbConnection()

	defer user.db.Close()

	stmt, err := user.db.Prepare("SELECT * FROM users WHERE id=?")

	if err != nil {
		log.Fatal(err)
	}

	var userFromDb GetUserDb

	err = stmt.QueryRow(id).Scan(&userFromDb.id, &userFromDb.name, &userFromDb.email)

	if err != nil {
		log.Fatal(err)
	}

	return &GetUserEntity{
		ID:    userFromDb.id,
		Name:  userFromDb.name,
		Email: userFromDb.email,
	}

}

func (user *GetUser) DataAccess() []*GetUserEntity {
	user.db = shared.GetDbConnection()

	defer user.db.Close()

	rows, err := user.db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var collect []*GetUserEntity

	for rows.Next() {
		var userFromDb GetUserDb
		err = rows.Scan(&userFromDb.id, &userFromDb.name, &userFromDb.email)

		if err != nil {
			log.Fatal(err)
		}

		collect = append(collect, &GetUserEntity{
			ID:    userFromDb.id,
			Name:  userFromDb.name,
			Email: userFromDb.email,
		})
	}

	return collect 
}
