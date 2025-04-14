package entity

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Nama		int 		 `json:"nama"`
	FilmGenre  []FilmGenre  `json:"film_genre" gorm:"foreignKey:GenreID"`
}