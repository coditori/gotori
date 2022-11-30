package models

type Video struct {
	Title       string `json:"title" xml:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
