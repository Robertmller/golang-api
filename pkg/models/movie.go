package models

import (
	"golang-api/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Title    string `gorm:"" json:"title"`
	Imdb     string `json:"imdb"`
	Director string `json:"director"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (b *Movie) CreateMovie() *Movie {
	db.NewRecord(b)
	db.Create(&b)
	return b

}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMovieById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(Id int64) Movie {
	var movie Movie
	db.Where("ID=?", Id).Delete(movie)
	return movie
}
