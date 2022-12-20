package models

import "time"

type Person struct {
	ID        uint64 `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstName" binding:"required" gorm:"type:varhar(32)"`
	LastName  string `json:"lastName" binding:"required" gorm:"type:varhar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varhar(256)"`
}

type Video struct {
	// xml:"title" form:"title" validate:"email" binding:"required"
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=20" validate:"is-cool" gorm:"type:varhar(100)"`
	Description string    `json:"description" binding:"max=20" gorm:"type:varhar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varhar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
