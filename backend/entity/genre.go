package entity

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	ID        int         `json:"id"`
	Nama      string      `json:"nama" binding:"required"`
	FilmGenre []FilmGenre `json:"film_genre" gorm:"foreignKey:GenreID"`
}
