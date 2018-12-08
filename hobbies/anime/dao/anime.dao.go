package dao

import (
	."golang-api-sandbox/hobbies/anime/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AnimeDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "animes"
)

func (dao *AnimeDAO) Connect() {
	session, err := mgo.Dial(dao.Server)

	if err != nil {
		log.Fatal(err)
	}

	session.DB(dao.Database)
}

func (dao *AnimeDAO) GetAll() ([]Anime, error) {
	var animes []Anime
	err := db.C(COLLECTION).Find(bson.M{}).All(&animes)
	return animes, err
}

func (dao *AnimeDAO) GetById(id string) (Anime, error) {
	var anime Anime
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&anime)
	return anime, err
}

func (dao *AnimeDAO) GetByName(name string) (Anime, error) {
	var anime Anime
	err := db.C(COLLECTION).Find(bson.M{name: name}).One(&anime)
	return anime, err
}

func (dao *AnimeDAO) GetAllByName(name string) ([]Anime, error) {
	var animes []Anime
	err := db.C(COLLECTION).Find(bson.M{name: name}).All(&animes)
	return animes, err
}