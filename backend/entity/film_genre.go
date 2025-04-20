package entity

import (
	"gorm.io/gorm"
)

type FilmGenre struct {
	gorm.Model
	FilmID  int   `json:"film_id"`
	GenreID int   `json:"genre_id"`
	Film    Film  `gorm:"foreignKey:FilmID"`
	Genre   Genre `gorm:"foreignKey:GenreID"`
}
