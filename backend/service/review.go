package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
)

type ReviewService interface {
	CreateReview(ctx context.Context, review dto.CreateReviewRequest, userId int) (dto.CreateReviewResponse, error)
	GetReviewByUserId(ctx context.Context, id int, userId int, page int) (dto.ReviewByUserResponse, error)
	GetReviewByFilmId(ctx context.Context, filmId int, userId int, page int) (dto.ReviewByFilmResponse, error)
	UpdateReview(ctx context.Context, review dto.UpdateReviewRequest) error
	UpdateReaksiReview(ctx context.Context, review dto.UpdateReaksiReviewRequest) error
	DeleteReview(ctx context.Context, id int) error
}

type reviewService struct {
	reviewRepository       repository.ReviewRepository
	reaksiReviewRepository repository.ReaksiReviewRepository
	userFilmRepository     repository.UserFilmRepository
	filmRepository         repository.FilmRepository
}

func NewReviewService(
	reviewRepository repository.ReviewRepository,
	reaksiReviewRepository repository.ReaksiReviewRepository,
	userFilmRepository repository.UserFilmRepository,
	filmRepository repository.FilmRepository,
) ReviewService {
	return &reviewService{
		reviewRepository:       reviewRepository,
		reaksiReviewRepository: reaksiReviewRepository,
		userFilmRepository:     userFilmRepository,
		filmRepository:         filmRepository,
	}
}

func (s *reviewService) CreateReview(ctx context.Context, reviewReq dto.CreateReviewRequest, userId int) (dto.CreateReviewResponse, error) {
	review := entity.Review{
		FilmID:   reviewReq.FilmID,
		UserID:   userId,
		Komentar: reviewReq.Komentar,
		Rating:   reviewReq.Rating,
	}

	film, err := s.filmRepository.CheckStatusFilm(ctx, reviewReq.FilmID)
	if err != nil {
		return dto.CreateReviewResponse{}, err
	}
	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED {
		return dto.CreateReviewResponse{}, dto.ErrCreateReview
	}

	checkUserFilm, err := s.userFilmRepository.CheckUserFilm(ctx, userId, reviewReq.FilmID)
	if err != nil || !checkUserFilm {
		return dto.CreateReviewResponse{}, dto.ErrCheckUserFilm
	}

	createdReview, err := s.reviewRepository.CreateReview(ctx, review)
	if err != nil {
		return dto.CreateReviewResponse{}, dto.ErrCreateReview
	}

	return dto.CreateReviewResponse{
		ID:       createdReview.ID,
		Komentar: createdReview.Komentar,
		Rating:   createdReview.Rating,
	}, nil
}

func (s *reviewService) GetReviewByUserId(ctx context.Context, id int, userId int, page int) (dto.ReviewByUserResponse, error) {
	reviews, countReview, err := s.reviewRepository.GetReviewByUserId(ctx, id, page)
	if err != nil {
		return dto.ReviewByUserResponse{}, dto.ErrGetReviewByUserId
	}

	var reviewsResponse dto.ReviewByUserResponse
	reviewsResponse.CountReview = int(countReview)

	for _, review := range reviews {
		userReaksiReview, _ := s.reaksiReviewRepository.GetReaksiReviewByUserId(ctx, review.ID, userId)
		reviewResponse := dto.ReviewResponse{
			ID:       review.ID,
			Komentar: review.Komentar,
			Rating:   review.Rating,
			User: dto.UserReview{
				ID:       review.User.ID,
				Username: review.User.Username,
			},
			UserReaksi: userReaksiReview,
		}

		reviewsResponse.Reviews = append(reviewsResponse.Reviews, reviewResponse)
	}

	return reviewsResponse, nil
}

func (s *reviewService) GetReviewByFilmId(ctx context.Context, filmId int, userId int, page int) (dto.ReviewByFilmResponse, error) {
	reviews, countPage, err := s.reviewRepository.GetReviewByFilmId(ctx, filmId, page)
	if err != nil {
		return dto.ReviewByFilmResponse{}, dto.ErrGetReviewFilmById
	}

	var reviewsResponse dto.ReviewByFilmResponse
	reviewsResponse.CountPage = int(countPage)

	for _, review := range reviews {
		userReaksiReview, _ := s.reaksiReviewRepository.GetReaksiReviewByUserId(ctx, review.ID, userId)

		reviewResponse := dto.ReviewResponse{
			ID:       review.ID,
			Komentar: review.Komentar,
			Rating:   review.Rating,
			User: dto.UserReview{
				ID:       review.User.ID,
				Username: review.User.Username,
			},
			UserReaksi: userReaksiReview,
		}

		reviewsResponse.Reviews = append(reviewsResponse.Reviews, reviewResponse)
	}

	return reviewsResponse, nil
}

func (s *reviewService) UpdateReview(ctx context.Context, review dto.UpdateReviewRequest) error {
	return s.reviewRepository.UpdateReview(ctx, review)
}

func (s *reviewService) UpdateReaksiReview(ctx context.Context, review dto.UpdateReaksiReviewRequest) error {
	reaksiReview := entity.ReaksiReview{
		ID:       review.ID,
		Reaksi:   review.Reaksi,
		UserID:   review.UserID,
		ReviewID: review.ReviewID,
	}

	err := s.reaksiReviewRepository.UpdateOrCreateUserReaksi(ctx, reaksiReview)
	if err != nil {
		return dto.ErrUpdateReview
	}

	return nil
}

func (s *reviewService) DeleteReview(ctx context.Context, id int) error {
	return s.reviewRepository.DeleteReview(ctx, id)
}
