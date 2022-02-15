package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const DNS = "root:password@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Tweet struct {
	gorm.Model

	Title string `json:"title"`
	Body  string `json:"body"`

	User string `json:"user"`
}

func InitialMigration() {
	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to db")
	}
	db.AutoMigrate(&Tweet{})
}

func Gettweets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var tweets []Tweet
	db.Find(&tweets)
	json.NewEncoder(w).Encode(tweets)

}

func Showtweet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Tweet
	db.First(&tweet, params["id"])
	json.NewEncoder(w).Encode(tweet)

}

func Deletetweets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Tweet
	db.Delete(&tweet, params["id"])
	json.NewEncoder(w).Encode("The tweet is Deleted Successfully!")

}

func Addtweets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var tweet Tweet
	json.NewDecoder(r.Body).Decode(&tweet)
	db.Create(&tweet)
	json.NewEncoder(w).Encode(tweet)

}

func Updatetweets(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tweet Tweet
	db.First(&tweet, params["id"])
	json.NewDecoder(r.Body).Decode(&tweet)
	db.Save(&tweet)
	json.NewEncoder(w).Encode(tweet)

}
