package anime

import (
	"encoding/json"
	."golang-api-sandbox/api/rest"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var dao = AnimeDAO{}

func CreateAnime(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var anime Anime

	if err := json.NewDecoder(req.Body).Decode(&anime); err != nil {
		ResponseBadRequest(resp)
	}

	anime.Id = bson.NewObjectId().String()

	if err := dao.Insert(anime); err != nil {
		log.Fatal(err.Error())
		ResponseInternalServerError(resp)
		return
	}

	ResponseCreated(resp, anime)
}

func GetAnimes(resp http.ResponseWriter, req *http.Request) {
	animes, err := dao.GetAll()

	if err != nil {
		ResponseInternalServerError(resp)
		return
	}

	ResponseSuccess(resp, animes)
}