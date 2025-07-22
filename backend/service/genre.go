package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/repository"
)

type GenreService interface {
	GetAllGenre(ctx context.Context) ([]dto.GenreRequest, error)
	CreateGenre(ctx context.Context, genre dto.GenreRequest) (dto.GenreResponse, error)
	DeleteGenre(ctx context.Context, genreId int) error
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

	var response []dto.GenreRequest
	for _, genre := range genres {
		response = append(response, dto.GenreRequest{
			ID:   genre.ID,
			Nama: genre.Nama,
		})
	}

	return response, nil
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

	response := dto.GenreResponse{
		ID:   createdGenre.ID,
		Nama: createdGenre.Nama,
	}

	return response, nil
}

func (s *genreService) DeleteGenre(ctx context.Context, genreId int) error {
	err := s.genreRepository.DeleteGenre(ctx, genreId)
	if err != nil {
		return dto.ErrDeleteGenre
	}

	return nil
}
