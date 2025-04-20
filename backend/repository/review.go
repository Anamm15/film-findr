package repository

import (
	"context"
	"math"

	"ReviewPiLem/dto"
	"ReviewPiLem/entity"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewByFilmId(ctx context.Context, filmId int, page int) ([]entity.Review, int64, error)
	GetReviewByUserId(ctx context.Context, id int, page int) ([]entity.Review, int64, error)
	GetRatingByFilmID(ctx context.Context, filmId int) (float64, error)
	CreateReview(ctx context.Context, review entity.Review) (entity.Review, error)
	UpdateReview(ctx context.Context, review dto.UpdateReviewRequest) error
	UpdateReaksiReview(ctx context.Context, review dto.UpdateReaksiReviewRequest) error
	DeleteReview(ctx context.Context, id int) error
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetReviewByFilmId(ctx context.Context, filmId int, page int) ([]entity.Review, int64, error) {
	var reviews []entity.Review
	var countReview int64

	const limit = 5
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Model(&entity.Review{}).
		Where("film_id = ?", filmId).
		Count(&countReview).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Preload("User", func(db *gorm.DB) *gorm.DB { return db.Select("id", "username") }).
		Select("id", "komentar", "rating", "created_at", "user_id").
		Where("film_id = ?", filmId).Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(countReview) / float64(limit))
	return reviews, int64(totalPage), nil
}

func (r *reviewRepository) GetReviewByUserId(ctx context.Context, id int, page int) ([]entity.Review, int64, error) {
	var review []entity.Review
	var countReview int64

	const limit = 5
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Model(&entity.Review{}).
		Where("user_id = ?", id).
		Count(&countReview).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Select("id", "komentar", "rating", "created_at").
		Where("user_id = ?", id).
		Find(&review).Error; err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(countReview) / float64(limit))
	return review, int64(totalPage), nil
}

func (r *reviewRepository) GetRatingByFilmID(ctx context.Context, filmId int) (float64, error) {
	var rating float64

	if err := r.db.WithContext(ctx).
		Model(&entity.Review{}).
		Select("AVG(rating)").
		Where("film_id = ?", filmId).
		Scan(&rating).Error; err != nil {
		return 0, err
	}

	return rating, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review entity.Review) (entity.Review, error) {
	if err := r.db.WithContext(ctx).Create(&review).Error; err != nil {
		return entity.Review{}, err
	}
	return review, nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, req dto.UpdateReviewRequest) error {
	var review entity.Review
	if err := r.db.First(&review, req.ID).Error; err != nil {
		return err
	}

	if req.Komentar != "" {
		review.Komentar = req.Komentar
	}
	if req.Rating != 0 {
		review.Rating = req.Rating
	}

	if err := r.db.WithContext(ctx).Save(&review).Error; err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) UpdateReaksiReview(ctx context.Context, review dto.UpdateReaksiReviewRequest) error {
	var updatedReview entity.Review

	if err := r.db.WithContext(ctx).
		Model(&updatedReview).
		Where("id = ?", review.ID).
		Update("reaksi", review.Reaksi).Error; err != nil {
		return err
	}

	return nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Review{}).Error; err != nil {
		return err
	}
	return nil
}
