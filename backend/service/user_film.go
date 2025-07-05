package service

import (
	"context"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
	"FilmFindr/utils"
)

type UserFilmService interface {
	GetUserFilmByUserId(ctx context.Context, userId int) ([]dto.UserFilmResponse, error)
	CreateUserFilm(ctx context.Context, userFilm dto.UserFilmCreateRequest) (entity.UserFilm, error)
	UpdateStatusUserFilm(ctx context.Context, userFilm dto.UserFilmUpdateStatusRequest) error
}

type userFilmService struct {
	userFilmRepository repository.UserFilmRepository
	filmRepository     repository.FilmRepository
}

func NewUserFilmService(userFilmRepository repository.UserFilmRepository, filmRepository repository.FilmRepository) UserFilmService {
	return &userFilmService{userFilmRepository: userFilmRepository, filmRepository: filmRepository}
}

func (s *userFilmService) GetUserFilmByUserId(ctx context.Context, userId int) ([]dto.UserFilmResponse, error) {
	userFilms, err := s.userFilmRepository.GetUserFilmByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	var userFilmResponses []dto.UserFilmResponse
	for _, userFilm := range userFilms {
		var FilmGambarResponse []dto.FilmGambarResponse
		var FilmGenreResponse []dto.GenreResponse
		var FilmResponse dto.FilmResponse

		for _, FilmGambar := range userFilm.Film.FilmGambar {
			FilmGambarResponse = append(FilmGambarResponse, dto.FilmGambarResponse{
				ID:  FilmGambar.ID,
				Url: FilmGambar.Url,
			})
		}

		for _, FilmGenre := range userFilm.Film.FilmGenre {
			FilmGenreResponse = append(FilmGenreResponse, dto.GenreResponse{
				ID:   FilmGenre.Genre.ID,
				Nama: FilmGenre.Genre.Nama,
			})
		}

		formattedDate := utils.FormatDate(userFilm.Film.TanggalRilis)
		FilmResponse = dto.FilmResponse{
			ID:           userFilm.Film.ID,
			Judul:        userFilm.Film.Judul,
			TanggalRilis: formattedDate,
			Durasi:       userFilm.Film.Durasi,
			Status:       userFilm.Film.Status,
			Gambar:       FilmGambarResponse,
			Genres:       FilmGenreResponse,
		}

		userFilmResponses = append(userFilmResponses, dto.UserFilmResponse{
			ID:     userFilm.ID,
			Status: userFilm.Status,
			Film:   FilmResponse,
		})
	}

	return userFilmResponses, nil
}

func (s *userFilmService) CreateUserFilm(ctx context.Context, userFilmReq dto.UserFilmCreateRequest) (entity.UserFilm, error) {
	userFilm := entity.UserFilm{
		Status: userFilmReq.Status,
		UserID: userFilmReq.UserID,
		FilmID: userFilmReq.FilmID,
	}

	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilmReq.FilmID)
	if err != nil {
		return entity.UserFilm{}, err
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return entity.UserFilm{}, dto.ErrCreateUserFilm
	}

	userFilmRes, err := s.userFilmRepository.CreateUserFilm(ctx, userFilm)
	if err != nil {
		return entity.UserFilm{}, err
	}

	return userFilmRes, nil
}

func (s *userFilmService) UpdateStatusUserFilm(ctx context.Context, userFilm dto.UserFilmUpdateStatusRequest) error {
	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilm.FilmID)
	if err != nil {
		return err
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return dto.ErrUpdateStatusUserFilm
	}

	if err := s.userFilmRepository.UpdateStatusUserFilm(ctx, userFilm.ID, userFilm.Status); err != nil {
		return dto.ErrUpdateStatusUserFilm
	}

	return nil
}
