package repository

import (
	"context"

	"ReviewPiLem/entity"

	"gorm.io/gorm"
)

type UserFilmRepository interface {
	GetUserFilmByUserId(ctx context.Context, userId int) ([]entity.UserFilm, error)
	CreateUserFilm(ctx context.Context, userFilm entity.UserFilm) (entity.UserFilm, error)
	UpdateStatusUserFilm(ctx context.Context, userFilmId int, status string) error
	CheckUserFilm(ctx context.Context, userId int, filmId int) (bool, error)
}

type userFilmRepository struct {
	db *gorm.DB
}

func NewUserFilmRepository(db *gorm.DB) UserFilmRepository {
	return &userFilmRepository{db: db}
}

func (r *userFilmRepository) GetUserFilmByUserId(ctx context.Context, userId int) ([]entity.UserFilm, error) {
	var userFilms []entity.UserFilm

	if err := r.db.WithContext(ctx).
		Model(&entity.UserFilm{}).
		Select("id", "status", "user_id", "film_id").
		Where("user_id = ?", userId).
		Preload("Film", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "judul", "tanggal_rilis", "durasi", "status")
		}).
		Preload("Film.FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).
		Preload("Film.FilmGenre.Genre", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nama")
		}).
		Find(&userFilms).Error; err != nil {
		return nil, err
	}

	return userFilms, nil
}

func (r *userFilmRepository) CreateUserFilm(ctx context.Context, userFilm entity.UserFilm) (entity.UserFilm, error) {
	if err := r.db.Create(&userFilm).Error; err != nil {
		return entity.UserFilm{}, err
	}

	return userFilm, nil
}

func (r *userFilmRepository) UpdateStatusUserFilm(ctx context.Context, userFilmId int, status string) error {
	if err := r.db.WithContext(ctx).Table("user_films").Where("id = ?", userFilmId).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}

func (r *userFilmRepository) CheckUserFilm(ctx context.Context, userId int, filmId int) (bool, error) {
	var userFilm entity.UserFilm
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND film_id = ? AND status <> 'plan to watch'", userId, filmId).
		First(&userFilm).Error; err != nil {
		return false, err
	}

	if userFilm.ID == 0 {
		return false, nil
	}

	return true, nil
}
