package service

// import (
// 	"ReviewPiLem/dto"
// 	"ReviewPiLem/repository"
// 	"context"
// )

// type ReviewReaksiService interface {
// 	CheckUserReaksi(ctx context.Context, userId uint) (dto.UserReaksiReview, error)
// 	UpdateUserReaksi(ctx context.Context, review dto.UpdateReaksiReviewRequest) error
// }

// type reviewReaksiService struct {
// 	reviewRepository repository.ReviewRepository
// 	reviewReaksiRepository repository.ReaksiReviewRepository
// }

// func NewReviewReaksiService(reviewReaksiRepository repository.ReaksiReviewRepository) ReviewReaksiService {
// 	return &reviewReaksiService{reviewReaksiRepository: reviewReaksiRepository}
// }

// func (s *reviewReaksiService) CheckUserReaksi(ctx context.Context, userId uint) (dto.UserReaksiReview, error) {
// 	return s.reviewReaksiRepository.CheckUserReaksi(ctx, userId)
// }

// func (s *reviewReaksiService) UpdateUserReaksi(ctx context.Context, review dto.UpdateReaksiReviewRequest) error {
// 	return s.reviewReaksiRepository.UpdateUserReaksi(ctx, review)
// }
