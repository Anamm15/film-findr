package entity

import (
	"gorm.io/gorm"
)

type ReaksiReview struct {
	gorm.Model
	ID       int    `json:"id"`
	Reaksi   string `json:"reaksi" binding:"required"`
	UserID   int    `json:"user_id" gorm:"index:idx_user_review"`
	ReviewID int    `json:"komentar_id" gorm:"index:idx_user_review"`
	User     User   `gorm:"foreignKey:UserID"`
	Review   Review `gorm:"foreignKey:ReviewID"`
}
