package entity

import (
	"time"

	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Judul        string       `json:"judul" binding:"required"`
	Status       string       `json:"status" binding:"required" gorm:"index:idx_film_status"`
	Sinopsis     string       `json:"sinopsis" binding:"required"`
	Durasi       int          `json:"durasi" binding:"required"`
	TotalEpisode int          `json:"total_episode" binding:"required"`
	Sutradara    string       `json:"sutradara" binding:"required"`
	TanggalRilis time.Time    `gorm:"type:date" json:"tanggal_rilis" time_format:"2006-01-02" binding:"required"`
	FilmGambar   []FilmGambar `json:"film_gambar" gorm:"foreignKey:FilmID"`
	FilmGenre    []FilmGenre  `json:"film_genre" gorm:"foreignKey:FilmID"`
	Review       []Review     `json:"review" gorm:"foreignKey:FilmID"`
}
