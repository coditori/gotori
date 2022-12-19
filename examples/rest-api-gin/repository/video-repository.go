package repository

import (
	"rest/models"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video models.Video)
	Update(video models.Video)
	Delete(video models.Video)
	FindAll() []models.Video
	CLoseDB()
}

type database struct {
	connection *gorm.DB
}

func NewRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.Video{}, &models.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CLoseDB() {
	myDb, err := db.connection.DB()
	if err != nil {
		panic("Failed to connect database")
	}
	myDb.Close()
}

func (db *database) Save(video models.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video models.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video models.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []models.Video {
	var videos []models.Video
	db.connection.Find(&videos)
	return videos
}
