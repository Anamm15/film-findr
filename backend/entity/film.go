package entity

import (
	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Judul      string    `json:"judul"`
	Status     string    `json:"status"`
	Sinopsis   string    `json:"sinopsis"`
	Durasi     string    `json:"durasi"`
	Total_Episode int       `json:"total_episode"`
	Sutradara  string    `json:"sutradara"`
	TanggalRilis    string    `json:"tanggal"`
	FilmGambar []FilmGambar `json:"film_gambar" gorm:"foreignKey:FilmID"`
	FilmGenre  []FilmGenre  `json:"film_genre" gorm:"foreignKey:FilmID"`
}