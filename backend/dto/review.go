package dto

import "errors"

const (
	// failed
	MESSAGE_FAILED_GET_REVIEW     = "Failed get review"
	MESSAGE_FAILED_CREATED_REVIEW = "Failed created review"
	MESSAGE_FAILED_UPDATED_REVIEW = "Failed updated review"
	MESSAGE_FAILED_DELETED_REVIEW = "Failed deleted review"

	// success
	MESSAGE_SUCCESS_REVIEW_NOT_FOUND = "Review not found"
	MESSAGE_SUCCESS_CREATED_REVIEW   = "Review created successfully"
	MESSAGE_SUCCESS_UPDATED_REVIEW   = "Review updated successfully"
	MESSAGE_SUCCESS_DELETED_REVIEW   = "Review deleted successfully"
	MESSAGE_SUCCESS_GET_LIST_REVIEW  = "Success get list review"
	MESSAGE_SUCCESS_GET_REVIEW       = "Success get review"
)

var (
	ErrGetReviewByUserId = errors.New("Failed to get review")
	ErrGetReviewFilmById = errors.New("Failed to get review in this film")
	ErrGetReviewByID     = errors.New("Failed to get review")
	ErrCreateReview      = errors.New("Failed to create review")
	ErrUpdateReview      = errors.New("Failed to update review")
	ErrDeleteReview      = errors.New("Failed to delete review")
)

type (
	UserReview struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	UserReaksiReview struct {
		ID     int    `json:"id"`
		Reaksi string `json:"reaksi"`
		UserID int    `json:"user_id"`
	}

	ReviewResponse struct {
		ID         int              `json:"id"`
		Komentar   string           `json:"komentar"`
		Rating     int              `json:"rating"`
		User       UserReview       `json:"user"`
		UserReaksi UserReaksiReview `json:"user_reaksi"`
	}

	ReviewByFilmResponse struct {
		CountPage int              `json:"count_page"`
		Reviews   []ReviewResponse `json:"reviews"`
	}

	ReviewByUserResponse struct {
		CountReview int              `json:"count_review"`
		Reviews     []ReviewResponse `json:"reviews"`
	}

	CreateReviewRequest struct {
		FilmID   int    `json:"film_id" validate:"required" binding:"required"`
		Komentar string `json:"komentar" validate:"required" binding:"required"`
		Rating   int    `json:"rating" validate:"required" binding:"required"`
	}

	CreateReviewResponse struct {
		ID       int    `json:"id"`
		Komentar string `json:"komentar"`
		Rating   int    `json:"rating"`
	}

	UpdateReviewRequest struct {
		ID       int    `json:"id" validate:"required" binding:"required"`
		Komentar string `json:"komentar"`
		Rating   int    `json:"rating"`
	}

	UpdateReaksiReviewRequest struct {
		ID       int    `json:"id" validate:"required"`
		UserID   int    `json:"user_id" validate:"required"`
		ReviewID int    `json:"review_id" validate:"required"`
		Reaksi   string `json:"reaksi" validate:"required"`
	}
)
