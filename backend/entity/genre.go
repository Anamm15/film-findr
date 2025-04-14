package entity

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Nama		string 		 `json:"nama"`
	FilmGenre  []FilmGenre  `json:"film_genre" gorm:"foreignKey:GenreID"`
}