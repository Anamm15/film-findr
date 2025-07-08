package repository

import (
	"context"
	"math"

	"FilmFindr/dto"
	"FilmFindr/entity"

	"gorm.io/gorm"
)

type FilmRepository interface {
	GetAllFilm(ctx context.Context, page int) ([]entity.Film, int64, error)
	CreateFilm(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error)
	UpdateFilm(ctx context.Context, film dto.UpdateFilmRequest) (entity.Film, error)
	DeleteFilm(ctx context.Context, id int) error
	GetFilmByID(ctx context.Context, id int) (entity.Film, error)
	UpdateStatus(ctx context.Context, id int, status string) error
	CheckStatusFilm(ctx context.Context, id int) (entity.Film, error)
	SearchFilm(ctx context.Context, req dto.SearchFilmRequest) ([]entity.Film, error)
}

type filmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepository{db: db}
}

func (r *filmRepository) GetAllFilm(ctx context.Context, page int) ([]entity.Film, int64, error) {
	var films []entity.Film
	var count int64

	const limit = 10
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Model(&entity.Film{}).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&entity.Film{}).
		Select("id", "judul", "tanggal_rilis", "durasi", "status").
		Preload("FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).
		Preload("FilmGenre.Genre").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&films).Error; err != nil {
		return nil, 0, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(limit)))
	return films, totalPage, nil
}

func (r *filmRepository) GetFilmByID(ctx context.Context, id int) (entity.Film, error) {
	var film entity.Film
	if err := r.db.WithContext(ctx).
		Select("id", "judul", "tanggal_rilis", "durasi", "status", "total_episode", "sutradara", "sinopsis").
		Preload("FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).
		Preload("FilmGenre.Genre").
		First((&film), id).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) CreateFilm(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error) {
	err := tx.WithContext(ctx).Create(&film).Error
	return film, err
}

func (r *filmRepository) UpdateFilm(ctx context.Context, reqFilm dto.UpdateFilmRequest) (entity.Film, error) {
	var film entity.Film

	err := r.db.WithContext(ctx).Model(&entity.Film{}).Where("id = ?", reqFilm.ID).First(&film).Error

	if reqFilm.Judul != "" {
		film.Judul = reqFilm.Judul
	}

	if reqFilm.Sinopsis != "" {
		film.Sinopsis = reqFilm.Sinopsis
	}

	if reqFilm.Sutradara != "" {
		film.Sutradara = reqFilm.Sutradara
	}

	if reqFilm.Status != "" {
		film.Status = reqFilm.Status
	}

	if !reqFilm.TanggalRilis.IsZero() {
		film.TanggalRilis = reqFilm.TanggalRilis
	}

	if reqFilm.Durasi != 0 {
		film.Durasi = reqFilm.Durasi
	}

	if reqFilm.TotalEpisode != 0 {
		film.TotalEpisode = reqFilm.TotalEpisode
	}

	err = r.db.WithContext(ctx).Save(&film).Error
	return entity.Film{}, err
}

func (r *filmRepository) DeleteFilm(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Delete(&entity.Film{}, id).Error
	return err
}

func (r *filmRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	err := r.db.WithContext(ctx).Model(&entity.Film{}).Where("id = ?", id).Update("status", status).Error
	return err
}

func (r *filmRepository) CheckStatusFilm(ctx context.Context, id int) (entity.Film, error) {
	var film entity.Film
	if err := r.db.WithContext(ctx).Select("id", "status").Where("id = ?", id).First(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) SearchFilm(ctx context.Context, req dto.SearchFilmRequest) ([]entity.Film, error) {
	var films []entity.Film

	if req.Keyword != "" {
		r.db.Where("judul LIKE ? OR sinopsis LIKE ? OR sutradara LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	if req.Status != "" {
		r.db.Where("status = ?", req.Status)
	}

	if len(req.Genres) > 0 {
		r.db.Joins("JOIN film_genre ON film_genre.film_id = films.id").Where("film_genre.genre_id IN (?)", req.Genres)
	}

	if err := r.db.WithContext(ctx).
		Select("id", "judul", "tanggal_rilis", "durasi", "status").
		Preload("FilmGambar", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "url", "film_id")
		}).Preload("FilmGenre.Genre").Find(&films).Error; err != nil {
		return nil, err
	}

	return films, nil
}
