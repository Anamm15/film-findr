package entity

import (
	"FilmFindr/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int        `json:"id"`
	Nama        string     `json:"nama" binding:"required"`
	Username    string     `json:"username" gorm:"unique" binding:"required"`
	Password    string     `json:"password" binding:"required"`
	Role        string     `json:"role"`
	Bio         string     `json:"bio" binding:"required"`
	PhotoProfil string     `json:"photo_profil"`
	UserFilm    []UserFilm `json:"user_film" gorm:"foreignKey:UserID"`
	Review      []Review   `json:"review" gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
