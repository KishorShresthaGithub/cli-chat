package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	userUseCase "kslabs/chat-app-cli/user"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var itemMap = map[int]*Item{}

func RestApiExample() {

	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	router.HandleFunc("/users", handleGetAllUsers).Methods("GET")
	router.HandleFunc("/users", handleCreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handleUpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handleGetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handleDeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {

	var createUserEntity userUseCase.CreateUserEntity
	var createUserInterface userUseCase.CreateUserInterface

	err := json.NewDecoder(r.Body).Decode(&createUserEntity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUserInterface = &userUseCase.CreateUser{}
	createUserInterface.Controller(&createUserEntity, w)

}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	var updateUserEntity userUseCase.UpdateUserEntity
	var updateUserInterface userUseCase.UpdateUserInterface

	err = json.NewDecoder(r.Body).Decode(&updateUserEntity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateUserInterface = &userUseCase.UpdateUser{}
	updateUserInterface.Controller(&updateUserEntity, w, id)

}

func handleGetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	var getUserInterface userUseCase.GetUserInterface

	getUserInterface = &userUseCase.GetUser{}
	getUserInterface.GetSingleController(id, w)

}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request) {

	var getUserInterface userUseCase.GetUserInterface

	getUserInterface = &userUseCase.GetUser{}
	getUserInterface.GetAllController(w)

}

func handleDeleteUser(w http.ResponseWriter, r *http.Request){
  
  params:= mux.Vars(r)

  id,_:= strconv.Atoi(params["id"])


  var deleteUserInterface userUseCase.DeleteUserInterface
  deleteUserInterface = &userUseCase.DeleteUser{}
  deleteUserInterface.Controller(id, w)
}

func getItems(w http.ResponseWriter, r *http.Request) {

	items := make([]*Item, 0)

	for _, val := range itemMap {
		items = append(items, val)
	}
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if val, ok := itemMap[id]; ok {
		json.NewEncoder(w).Encode(val)
		return
	}

	http.NotFound(w, r)
}

func createItem(w http.ResponseWriter, r *http.Request) {

	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.ID = len(itemMap)

	itemMap[item.ID] = &item

	json.NewEncoder(w).Encode(item)

}

func updateItem(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedItem Item

	err = json.NewDecoder(r.Body).Decode(&updatedItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if val, ok := itemMap[id]; ok {
		itemMap[id] = &updatedItem
		json.NewEncoder(w).Encode(val)
		return
	}

	http.NotFound(w, r)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := itemMap[id]; ok {

		delete(itemMap, id)

		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.NotFound(w, r)

}
