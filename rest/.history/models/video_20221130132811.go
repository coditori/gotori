package models

type Person struct {
	FirstName string `json:"finstName"`
	LastName  string `json:"lastName"`
	Age       int8   `json:"age"`
	Email     string `json:"email" validate:"email"`
}

type Video struct {
	// xml:"title" form:"title" validate:"email" binding:"required"
	Title       string `json:"title" min:"2" max:"10"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Author      Person `json:"author"`
}
