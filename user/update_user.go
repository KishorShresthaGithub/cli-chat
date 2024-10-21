package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	shared "kslabs/chat-app-cli/shared"
	"log"
	"reflect"
	"strings"
)

type UpdateUserEntity struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateUserInterface interface {
	Controller(u *UpdateUserEntity, w io.Writer, id int)
}

type UpdateUser struct {
	db     *sql.DB
	writer *io.Writer
}

func (user *UpdateUser) SetWriter(write io.Writer) {
	user.writer = &write
}

func (user *UpdateUser) Presentation(u *UpdateUserEntity) {
	json.NewEncoder(*user.writer).Encode(&u)
}

func (user *UpdateUser) Controller(u *UpdateUserEntity, w io.Writer, id int) {
	user.SetWriter(w)
	user.DataAccess(u, id)
	user.Presentation(u)
}

func (user *UpdateUser) DataAccess(u *UpdateUserEntity, id int) {

	values := reflect.ValueOf(*u)
	types := values.Type()

	var queryArray []string
	var vals []interface{}

	for i := 0; i < values.NumField(); i++ {
		if !values.Field(i).IsZero() {
			fieldName := types.Field(i).Tag.Get("json")
			fieldName = strings.Replace(fieldName, ",omitempty", "", -1)
			queryArray = append(queryArray, fieldName+" = ?")
			vals = append(vals, values.Field(i).Interface())
		}
	}

	query := strings.Join(queryArray, ", ")

	user.db = shared.GetDbConnection()
	defer user.db.Close()

	tx, err := user.db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	stmt, err := tx.Prepare(fmt.Sprintf("UPDATE users SET %v WHERE id = ?", query))

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	vals = append(vals, id)

	// Convert []int to []interface{}
	args := make([]interface{}, len(vals))
	for i, v := range vals {
		args[i] = v
	}

	// Execute the statement
	result, err := stmt.Exec(args...)
	if err != nil {
		panic(err)
	}

	// Process the result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("failed to commit transaction: %w", err)
	}

	fmt.Printf("Update successful, rows affected: %d\n", rowsAffected)

}
