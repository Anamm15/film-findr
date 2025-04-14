package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama       string    `json:"nama"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	UserFilm   []UserFilm `json:"user_film" gorm:"foreignKey:UserID"`
	Review     []Review   `json:"review" gorm:"foreignKey:UserID"`
}

// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	var err error
// 	u.Password, err = helpers.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }