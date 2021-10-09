package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Instagram struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	//Type     string `json:"type"`
}

var users []Instagram

type Post struct {
	Type string `json:"type"`
	User string `json:"user"`
	Id   string `json:"id"`
	//Type     string `json:"type"`
}

var posts []Post

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Instagram{})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range posts {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}
func main() {

	users = append(users, Instagram{Username: "Sai", Id: "1"})
	users = append(users, Instagram{Username: "Vignesh", Id: "2"})
	posts = append(posts, Post{Type: "Photo", User: "1234", Id: "3"})
	posts = append(posts, Post{Type: "Photo", User: "1235", Id: "4"})

	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")

	//http.HandleFunc("/users", CreateUser)
	//http.HandleFunc("/users/{id}", getUser)

	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8000", r))

	


}
