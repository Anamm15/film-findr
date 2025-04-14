package entity

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating		int 		 `json:"rating"`
	Komentar	string		 `json:"komentar"`
	Reaksi		string		 `json:"reaksi"`
	UserID		int 		 `json:"user_id"`
	FilmID		int 		 `json:"film_id"`
	User User `gorm:"foreignKey:UserID"`
	Film Film `gorm:"foreignKey:FilmID"`
}