package repository

import (
	"context"

	"FilmFindr/entity"

	"gorm.io/gorm"
)

type FilmGambarRepository interface {
	Save(ctx context.Context, tx *gorm.DB, filmGambar entity.FilmGambar) error
}

type filmGambarRepository struct {
	db *gorm.DB
}

func NewFilmGambarRepository(db *gorm.DB) FilmGambarRepository {
	return &filmGambarRepository{db: db}
}

func (r *filmGambarRepository) Save(ctx context.Context, tx *gorm.DB, filmGambar entity.FilmGambar) error {
	if err := tx.Create(&filmGambar).Error; err != nil {
		return err
	}
	return nil
}
