package models

type Person struct {
	FirstName string `json:"finstName"`
	LastName  string `json:"lastName"`
	Age       int8   `json:"age"`
	Email     string `json:"email" validate:"email"`
}

type Video struct {
	// xml:"title" form:"title" validate:"email" binding:"required"
	Title       string `json:"title" binding:"min=2 max=10"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
