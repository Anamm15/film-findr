package repository

import (
	"context"
	"ReviewPiLem/entity"
	"gorm.io/gorm"
)


type FilmRepository interface {
	GetAllFilm(ctx context.Context) ([]entity.Film, error)
	CreateFilm(ctx context.Context, film entity.Film) (entity.Film, error)
	UpdateFilm(ctx context.Context, film entity.Film) (entity.Film, error)
	DeleteFilm(ctx context.Context, id uint64) error
	GetFilmByID(ctx context.Context, id uint64) (entity.Film, error)
}

type filmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepository{db: db}
}

func (r *filmRepository) GetAllFilm(ctx context.Context) ([]entity.Film, error) {
	var films []entity.Film
	if err := r.db.Find(&films).Error; err != nil {
		return nil, err
	}

	return films, nil
}

func (r *filmRepository) GetFilmByID(ctx context.Context, id uint64) (entity.Film, error) {
	var film entity.Film
	if err := r.db.First((&film), id).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) CreateFilm(ctx context.Context, film entity.Film) (entity.Film, error) {
	err := r.db.Create(&film).Error
	return film, err
}

func (r *filmRepository) UpdateFilm(ctx context.Context, film entity.Film) (entity.Film, error) {
	err := r.db.Save(&film).Error
	return film, err
}

func (r *filmRepository) DeleteFilm(ctx context.Context, id uint64) error {
	err := r.db.Delete(&entity.Film{}, id).Error
	return err
}