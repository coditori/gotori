package models

type Video struct {
	Title       string `json:"title" xml:"title" form:"title" "validate:"email""`
	Description string `json:"description"`
	URL         string `json:"url"`
}
