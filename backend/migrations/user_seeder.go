package migrations

import (
	"errors"

	"FilmFindr/entity"
	"FilmFindr/helpers"

	"gorm.io/gorm"
)

func ListUserSeeder(db *gorm.DB) error {
	listUser := []entity.User{
		{
			Nama:     "Admin1",
			Password: "admin123",
			Username: "admin",
			Role:     helpers.ENUM_ROLE_ADMIN,
			Bio:      "Admin 1",
		},
		{
			Nama:     "Admin2",
			Password: "admin123",
			Username: "admin2",
			Role:     helpers.ENUM_ROLE_ADMIN,
			Bio:      "Admin 2",
		},
		{
			Nama:     "User1",
			Password: "user123",
			Username: "user",
			Role:     helpers.ENUM_ROLE_USER,
			Bio:      "User 1",
		},
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var user entity.User
		err := db.Where(&entity.User{Username: data.Username}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "username = ?", data.Username).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
