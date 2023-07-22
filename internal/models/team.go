package models

type Team struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Position  string `json:"position"`
	Biography string `json:"biography"`
	Img       string `json:"img"`
}
