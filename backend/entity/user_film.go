package entity

import (
	"gorm.io/gorm"
)

type UserFilm struct {
	gorm.Model
	Status string `json:"status" binding:"required"`
	UserID int    `json:"user_id" gorm:"not null index:idx_user_id"`
	FilmID int    `json:"film_id" gorm:"not null index:idx_film_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	Film   Film   `gorm:"foreignKey:FilmID" json:"film"`
}
