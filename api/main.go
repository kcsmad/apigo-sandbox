package api

import (
	"github.com/gorilla/mux"
	."golang-api-sandbox/config"
	."golang-api-sandbox/hobbies/anime"
)

var config = Config{}
var animeDAO = AnimeDAO {}
var animes []Anime

func main() {
	routerConfig()

}


func routerConfig() {
	router := mux.NewRouter()
	router.HandleFunc()

}
