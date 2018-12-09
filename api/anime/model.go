package anime

import "gopkg.in/mgo.v2/bson"

type Anime struct {
	Id bson.ObjectId `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Genre string `json:"genre, omitempty"`
	Episodes int `json:"episodes, omitempty"`
	Duration int `json:"duration, omitempty"`
}