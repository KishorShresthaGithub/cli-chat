package user

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	shared "kslabs/chat-app-cli/shared"
	"log"
  "net/http"
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
	GetAllController(w http.ResponseWriter)
	GetSingleController(id int, w http.ResponseWriter)
}

type GetUser struct {
	db     *sql.DB
	writer *http.ResponseWriter
}

func (user *GetUser) SetWriter(writer *http.ResponseWriter) {
	user.writer = writer
}

func (user *GetUser) GetSingleController(id int, writer http.ResponseWriter) {
  user.SetWriter(&writer)
	u,err := user.DataAccessSingle(id)

  if err !=nil{
    http.Error(writer, "some error happened", http.StatusInternalServerError)   
    return
  }

	user.SinglePresentation(u)
}

func (user *GetUser) SinglePresentation(u *GetUserEntity) {
	json.NewEncoder(*user.writer).Encode(u)
}

func (user *GetUser) Presentation(u []*GetUserEntity) {
	json.NewEncoder(*user.writer).Encode(u)
}

func (user *GetUser) GetAllController(w http.ResponseWriter) {
	user.SetWriter(&w)
	u := user.DataAccess()
	user.Presentation(u)
}

func (user *GetUser) DataAccessSingle(id int) (*GetUserEntity, any) {
	user.db = shared.GetDbConnection()

	defer user.db.Close()

	stmt, err := user.db.Prepare("SELECT id, name, email FROM users WHERE id=?")

	if err != nil {
		log.Fatal(err)
	}

	var userFromDb GetUserDb

	err = stmt.QueryRow(id).Scan(&userFromDb.id, &userFromDb.name, &userFromDb.email)

	if err != nil {
    return nil, err
	}

	return &GetUserEntity{
		ID:    userFromDb.id,
		Name:  userFromDb.name,
		Email: userFromDb.email,
	},nil

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
