package entity

import (
	"gorm.io/gorm"
)

type FilmGambar struct {
	gorm.Model
	Url    string `json:"url"`
	FilmID int    `json:"film_id" gorm:"index:idx_film_id"`
	Film   Film   `gorm:"foreignKey:FilmID"`
}
