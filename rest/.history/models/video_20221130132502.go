package models

type Video struct {
	Title       string `json:"title" xml:"title" form:"title" "validate":"email" binding:"required"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
