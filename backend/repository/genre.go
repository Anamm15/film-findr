package repository

import (
	"context"

	"ReviewPiLem/entity"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAllGenre(ctx context.Context) ([]entity.Genre, error)
	CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error)
	UpdateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error)
}

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{db: db}
}

func (r *genreRepository) GetAllGenre(ctx context.Context) ([]entity.Genre, error) {
	var genres []entity.Genre
	if err := r.db.Find(&genres).Error; err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *genreRepository) CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error) {
	err := r.db.Create(&genre).Error
	return genre, err
}

func (r *genreRepository) UpdateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error) {
	err := r.db.Save(&genre).Error
	return genre, err
}
