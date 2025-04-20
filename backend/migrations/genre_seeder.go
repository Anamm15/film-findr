package migrations

import (
	"errors"

	"ReviewPiLem/entity"

	"gorm.io/gorm"
)

func ListGenreSeeder(db *gorm.DB) error {
	listGenre := []entity.Genre{
		{
			Nama: "Action",
		},
		{
			Nama: "Romance",
		},
		{
			Nama: "Sci-Fi",
		},
		{
			Nama: "Drama",
		},
	}

	hasTable := db.Migrator().HasTable(&entity.Genre{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Genre{}); err != nil {
			return err
		}
	}

	for _, data := range listGenre {
		var genre entity.Genre
		err := db.Where(&entity.Genre{Nama: data.Nama}).First(&genre).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&genre, "nama = ?", data.Nama).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
