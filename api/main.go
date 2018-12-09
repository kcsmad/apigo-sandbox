package main

import (
	"github.com/gorilla/mux"
	. "golang-api-sandbox/api/anime"
	. "golang-api-sandbox/infra/config"
	"log"
	"net/http"
)

var config = Config{}
var dao = AnimeDAO{}

func init() {
	config.ReadConfigFile()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	initServe()
}

func initServe() {
	router := mux.NewRouter()
	router.HandleFunc("/anime", GetAnimes).Methods("GET")
	router.HandleFunc("/planet", CreateAnime).Methods("POST")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
