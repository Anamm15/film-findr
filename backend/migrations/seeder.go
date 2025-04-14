package migrations

import (
	"errors"

	"ReviewPiLem/entity"
	"ReviewPiLem/helpers"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListUserSeeder(db); err != nil {
		return err
	}

	return nil
}

func ListUserSeeder(db *gorm.DB) error {
	var listUser = []entity.User{
		{
			Nama:       "Admin1",
			Password:   "admin123",
			Username:   "admin",
			Role:       helpers.ENUM_ROLE_ADMIN,
		},
		{
			Nama:       "Admin2",
			Password:   "admin123",
			Username:   "admin23",
			Role:       helpers.ENUM_ROLE_USER,
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