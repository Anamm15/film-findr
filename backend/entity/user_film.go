package entity

import (
	"gorm.io/gorm"
)

type UserFilm struct {
	gorm.Model
	Status      string    `json:"status"`
	UserID		int 		 `json:"user_id"`
	FilmID		int 		 `json:"film_id"`
	User User `gorm:"foreignKey:UserID" json:"user"`
	Film Film `gorm:"foreignKey:FilmID" json:"film"`
	Review     []Review   `json:"review" gorm:"foreignKey:FilmID"`
}