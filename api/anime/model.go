package anime

type Anime struct {
	Id string `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	Genre string `json:"genre, omitempty"`
	Episodes int `json:"episodes, omitempty"`
	Duration int `json:"duration, omitempty"`
}