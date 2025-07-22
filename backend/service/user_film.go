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
	GetUserFilmByUserId(ctx context.Context, userId int, page int) (dto.GetUserFilmResponse, error)
	CreateUserFilm(ctx context.Context, userFilm dto.UserFilmCreateRequest) (entity.UserFilm, error)
	UpdateStatusUserFilm(ctx context.Context, userFilmId int, userFilm dto.UserFilmUpdateStatusRequest) error
}

type userFilmService struct {
	userFilmRepository repository.UserFilmRepository
	filmRepository     repository.FilmRepository
}

func NewUserFilmService(
	userFilmRepository repository.UserFilmRepository,
	filmRepository repository.FilmRepository,
) UserFilmService {
	return &userFilmService{
		userFilmRepository: userFilmRepository,
		filmRepository:     filmRepository,
	}
}

func (s *userFilmService) GetUserFilmByUserId(ctx context.Context, userId int, page int) (dto.GetUserFilmResponse, error) {
	userFilms, countUserFilm, err := s.userFilmRepository.GetUserFilmByUserId(ctx, userId, page)
	if err != nil {
		return dto.GetUserFilmResponse{}, dto.ErrGetUserFilm
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

	var GetUserFilmResponse dto.GetUserFilmResponse
	GetUserFilmResponse.UserFilms = userFilmResponses
	GetUserFilmResponse.CountPage = int(countUserFilm)
	return GetUserFilmResponse, nil
}

func (s *userFilmService) CreateUserFilm(ctx context.Context, userFilmReq dto.UserFilmCreateRequest) (entity.UserFilm, error) {
	userFilm := entity.UserFilm{
		Status: userFilmReq.Status,
		UserID: userFilmReq.UserID,
		FilmID: userFilmReq.FilmID,
	}

	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilmReq.FilmID)
	if err != nil {
		return entity.UserFilm{}, dto.ErrCheckUserFilm
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return entity.UserFilm{}, dto.ErrStatusFilmNotYetAired
	}

	userFilmRes, err := s.userFilmRepository.CreateUserFilm(ctx, userFilm)
	if err != nil {
		return entity.UserFilm{}, dto.ErrCreateUserFilm
	}

	return userFilmRes, nil
}

func (s *userFilmService) UpdateStatusUserFilm(ctx context.Context, userFilmId int, userFilm dto.UserFilmUpdateStatusRequest) error {
	film, err := s.filmRepository.CheckStatusFilm(ctx, userFilm.FilmID)
	if err != nil {
		return err
	}

	if film.Status == helpers.ENUM_FILM_NOT_YET_AIRED && userFilm.Status != helpers.ENUM_LIST_FILM_PLAN_TO_WATCH {
		return dto.ErrUpdateStatusUserFilm
	}

	if err := s.userFilmRepository.UpdateStatusUserFilm(ctx, userFilmId, userFilm.Status); err != nil {
		return dto.ErrUpdateStatusUserFilm
	}

	return nil
}
