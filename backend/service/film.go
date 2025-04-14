package service

import (
	"ReviewPiLem/dto"
	"ReviewPiLem/entity"
	"ReviewPiLem/repository"
	"context"
)

type FilmService interface {
	GetAllFilm(ctx context.Context) ([]entity.Film, error)
	CreateFilm(ctx context.Context, film entity.Film) (entity.Film, error)
	UpdateFilm(ctx context.Context, film entity.Film) (entity.Film, error)
	DeleteFilm(ctx context.Context, id uint64) error
	GetFilmByID(ctx context.Context, id uint64) (entity.Film, error)
}

type filmService struct {
	filmRepository repository.FilmRepository
}

func NewFilmService(filmRepository repository.FilmRepository) FilmService {
	return &filmService{filmRepository: filmRepository}
}

func (s *filmService) GetAllFilm(ctx context.Context) ([]entity.Film, error) {
	films, err := s.filmRepository.GetAllFilm(ctx)
	if err != nil {
		return nil, dto.ErrGetAllGenre
	}

	return films, nil
}

func (s *filmService) CreateFilm(ctx context.Context, film entity.Film) (entity.Film, error) {
	createdFilm, err := s.filmRepository.CreateFilm(ctx, film)
	if err != nil {
		return entity.Film{}, dto.ErrCreateFilm
	}

	return createdFilm, nil
}

func (s *filmService) UpdateFilm(ctx context.Context, film entity.Film) (entity.Film, error) {
	updatedFilm, err := s.filmRepository.UpdateFilm(ctx, film)
	if err != nil {
		return entity.Film{}, dto.ErrUpdateFilm
	}

	return updatedFilm, nil
}

func (s *filmService) DeleteFilm(ctx context.Context, id uint64) error {
	err := s.filmRepository.DeleteFilm(ctx, id)
	if err != nil {
		return dto.ErrDeleteFilm
	}

	return nil
}

func (s *filmService) GetFilmByID(ctx context.Context, id uint64) (entity.Film, error) {
	film, err := s.filmRepository.GetFilmByID(ctx, id)
	if err != nil {
		return entity.Film{}, dto.ErrGetFilmByID
	}

	return film, nil
}