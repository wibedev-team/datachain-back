package models

type Solution struct {
	Title    string    `json:"title"`
	Features []Feature `json:"features"`
	Link     string    `json:"link"`
	File     string    `json:"file"`
}

type Feature struct {
	Text string `json:"text"`
}
