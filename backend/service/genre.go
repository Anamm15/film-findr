package service

import (
	"context"

	"ReviewPiLem/dto"
	"ReviewPiLem/entity"
	"ReviewPiLem/repository"
)

type GenreService interface {
	GetAllGenre(ctx context.Context) ([]entity.Genre, error)
	CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error)
	UpdateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error)
}

type genreService struct {
	genreRepository repository.GenreRepository
}

func NewGenreService(genreRepository repository.GenreRepository) GenreService {
	return &genreService{
		genreRepository: genreRepository,
	}
}

func (s *genreService) GetAllGenre(ctx context.Context) ([]entity.Genre, error) {
	genres, err := s.genreRepository.GetAllGenre(ctx)
	if err != nil {
		return nil, dto.ErrGetAllGenre
	}

	return genres, nil
}

func (s *genreService) CreateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error) {
	createdGenre, err := s.genreRepository.CreateGenre(ctx, genre)
	if err != nil {
		return entity.Genre{}, dto.ErrGetAllGenre
	}

	return createdGenre, nil
}

func (s *genreService) UpdateGenre(ctx context.Context, genre entity.Genre) (entity.Genre, error) {
	genre, err := s.genreRepository.UpdateGenre(ctx, genre)
	if err != nil {
		return entity.Genre{}, dto.ErrUpdateGenre
	}

	return genre, nil
}