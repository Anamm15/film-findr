package service

import (
	"context"

	"ReviewPiLem/dto"
	"ReviewPiLem/entity"
	"ReviewPiLem/repository"
)

type GenreService interface {
	GetAllGenre(ctx context.Context) ([]dto.GenreRequest, error)
	CreateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error)
	UpdateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error)
}

type genreService struct {
	genreRepository repository.GenreRepository
}

func NewGenreService(genreRepository repository.GenreRepository) GenreService {
	return &genreService{
		genreRepository: genreRepository,
	}
}

func (s *genreService) GetAllGenre(ctx context.Context) ([]dto.GenreRequest, error) {
	genres, err := s.genreRepository.GetAllGenre(ctx)
	if err != nil {
		return nil, dto.ErrGetAllGenre
	}

	var genreResponse []dto.GenreRequest
	for _, genre := range genres {
		genreResponse = append(genreResponse, dto.GenreRequest{
			ID:   genre.ID,
			Nama: genre.Nama,
		})
	}

	return genreResponse, nil
}

func (s *genreService) CreateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error) {
	GenreRequest := entity.Genre{
		ID:   genre.ID,
		Nama: genre.Nama,
	}

	createdGenre, err := s.genreRepository.CreateGenre(ctx, GenreRequest)
	if err != nil {
		return dto.GenreResponse{}, dto.ErrGetAllGenre
	}

	return dto.GenreResponse{ID: createdGenre.ID, Nama: createdGenre.Nama}, nil
}

func (s *genreService) UpdateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error) {
	genreEntity := entity.Genre{
		ID:   genre.ID,
		Nama: genre.Nama,
	}

	updatedGenre, err := s.genreRepository.UpdateGenre(ctx, genreEntity)
	if err != nil {
		return dto.GenreResponse{}, dto.ErrUpdateGenre
	}

	return dto.GenreResponse{
		ID:   updatedGenre.ID,
		Nama: updatedGenre.Nama,
	}, nil
}
