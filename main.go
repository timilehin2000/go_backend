package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type user struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

var users = []user{}

func (usr *user) saveUser() {
	usr.ID = len(users) + 1
	users = append(users, *usr)
}

func main() {
	http.HandleFunc("/users/new", AddNewUserHandler)
	http.HandleFunc("/users/", GetAllUsers)
	http.HandleFunc("/user/single", GetAUser)
	http.HandleFunc("/user/update", UpdateUser)

}

func AddNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var userData user

	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data received"))
		return
	}

	userData.saveUser()
	json.NewEncoder(w).Encode(userData)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetAUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	if userId > len(users) || userId < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user ID is out of range"))
		return
	}

	user := users[userId-1]
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	if userId > len(users) || userId < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user ID is out of range"))
		return
	}

	user := users[userId-1]
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data submited"))
		return
	}

	json.NewEncoder(w).Encode(user)
}
