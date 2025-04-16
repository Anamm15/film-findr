package entity

import (
	"gorm.io/gorm"
	"time"
)

type Film struct {
	gorm.Model
	Judul      string    `json:"judul"`
	Status     string    `json:"status"`
	Sinopsis   string    `json:"sinopsis"`
	Durasi     int    `json:"durasi"`
	TotalEpisode int       `json:"total_episode"`
	Sutradara  string    `json:"sutradara"`
	TanggalRilis time.Time `gorm:"type:date" json:"tanggal_rilis"`
	FilmGambar []FilmGambar `json:"film_gambar" gorm:"foreignKey:FilmID"`
	FilmGenre  []FilmGenre  `json:"film_genre" gorm:"foreignKey:FilmID"`
}