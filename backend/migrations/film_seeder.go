package migrations

import (
	"errors"
	"time"

	"FilmFindr/entity"
	"FilmFindr/helpers"

	"gorm.io/gorm"
)

func ListFilmSeeder(db *gorm.DB) error {
	listFilm := []entity.Film{
		{
			Judul:        "Spiderman 1",
			Status:       helpers.ENUM_FILM_AIRING,
			Sinopsis:     "random",
			Durasi:       120,
			TotalEpisode: 1,
			Sutradara:    "Anam",
			TanggalRilis: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Judul:        "Spiderman 2",
			Status:       helpers.ENUM_FILM_AIRING,
			Sinopsis:     "random",
			Durasi:       120,
			TotalEpisode: 10,
			Sutradara:    "Anam",
			TanggalRilis: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	if !db.Migrator().HasTable(&entity.Film{}) {
		if err := db.Migrator().CreateTable(&entity.Film{}); err != nil {
			return err
		}
	}

	for _, data := range listFilm {
		var film entity.Film
		err := db.Where("judul = ?", data.Judul).First(&film).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}
