package models

type Video struct {
	// xml:"title" form:"title" "validate":"email" binding:"required"
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
