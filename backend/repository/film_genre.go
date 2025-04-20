package repository

import (
	"context"

	"ReviewPiLem/entity"

	"gorm.io/gorm"
)

type FilmGenreRepository interface {
	CreateFilmGenre(ctx context.Context, tx *gorm.DB, genre entity.FilmGenre) (entity.FilmGenre, error)
	DeleteFilmGenre(ctx context.Context, filmGenre entity.FilmGenre) error
}

type filmGenreRepository struct {
	db *gorm.DB
}

func NewFilmGenreRepository(db *gorm.DB) FilmGenreRepository {
	return &filmGenreRepository{db: db}
}

func (r *filmGenreRepository) CreateFilmGenre(ctx context.Context, tx *gorm.DB, genre entity.FilmGenre) (entity.FilmGenre, error) {
	if err := tx.Create(&genre).Error; err != nil {
		return entity.FilmGenre{}, err
	}

	return genre, nil
}

func (r *filmGenreRepository) DeleteFilmGenre(ctx context.Context, genre entity.FilmGenre) error {
	var newFilmGenre entity.FilmGenre
	if err := r.db.WithContext(ctx).
		Delete(&newFilmGenre, "film_id = ? AND genre_id = ?", genre.FilmID, genre.GenreID).Error; err != nil {
		return err
	}

	return nil
}
