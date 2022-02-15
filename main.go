package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Initializer() {
	r := mux.NewRouter()

	r.HandleFunc("/tweets", Gettweets).Methods("get")
	r.HandleFunc("/tweets/{id}", Showtweet).Methods("get")
	r.HandleFunc("/tweets", Addtweets).Methods("post")
	r.HandleFunc("/tweets/{id}", Updatetweets).Methods("put")
	r.HandleFunc("/tweets/{id}", Deletetweets).Methods("delete")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	InitialMigration()

	Initializer()

}
