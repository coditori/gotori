package repository

import (
	"errors"
	"rest/models"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video models.Video) (*models.Video, error)
	Update(video models.Video) models.Video
	Delete(video models.Video)
	FindById(id uint64) models.Video
	FindAll() []models.Video
	CLoseDB()
}

type database struct {
	connection *gorm.DB
}

func New() VideoRepository {
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

func (db *database) Save(video models.Video) (*models.Video, error) {
	tx := db.connection.Create(&video)
	if tx.Error != nil {
		return nil, errors.New("could not create the resurce")
	}
	return &video, nil
}

func (db *database) Update(video models.Video) models.Video {
	db.connection.Save(&video)
	return video
}

func (db *database) Delete(video models.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindById(id uint64) models.Video {
	var videoResult models.Video
	db.connection.First(&videoResult, id)
	return videoResult
}

func (db *database) FindAll() []models.Video {
	var videos []models.Video
	db.connection.Preload("Author").Find(&videos)
	return videos
}
