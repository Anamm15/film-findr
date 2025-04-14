package entity

import (
	"gorm.io/gorm"
)

type UserFilm struct {
	gorm.Model
	Status      string    `json:"status"`
	UserID		int 		 `json:"user_id"`
	FilmID		int 		 `json:"film_id"`
	Review     []Review   `json:"review" gorm:"foreignKey:FilmID"`
}