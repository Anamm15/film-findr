package service

import (
	"context"
	"math"
	"mime/multipart"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/repository"
	"FilmFindr/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type FilmService interface {
	GetAllFilm(ctx context.Context, page int) ([]dto.FilmResponse, error)
	GetFilmByID(ctx context.Context, id int) (dto.FilmResponse, error)
	CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmResponse, error)
	UpdateFilm(ctx context.Context, film dto.UpdateFilmRequest) (entity.Film, error)
	DeleteFilm(ctx context.Context, id int) error
	UpdateStatus(ctx context.Context, id int, req dto.UpdateStatusFilmRequest) error
	SearchFilm(ctx context.Context, req dto.SearchFilmRequest) ([]dto.FilmResponse, error)
}

type filmService struct {
	filmRepository       repository.FilmRepository
	filmGambarRepository repository.FilmGambarRepository
	filmGenreRepository  repository.FilmGenreRepository
	reviewRepository     repository.ReviewRepository
	cloudinary           *cloudinary.Cloudinary
	db                   *gorm.DB
}

func NewFilmService(
	db *gorm.DB,
	cloudinary *cloudinary.Cloudinary,
	filmRepository repository.FilmRepository,
	filmGambarRepository repository.FilmGambarRepository,
	filmGenreRepository repository.FilmGenreRepository,
	reviewRepository repository.ReviewRepository,
) FilmService {
	return &filmService{
		db:                   db,
		cloudinary:           cloudinary,
		filmRepository:       filmRepository,
		filmGambarRepository: filmGambarRepository,
		filmGenreRepository:  filmGenreRepository,
		reviewRepository:     reviewRepository,
	}
}

func (s *filmService) GetAllFilm(ctx context.Context, page int) ([]dto.FilmResponse, error) {
	films, err := s.filmRepository.GetAllFilm(ctx, page)
	if err != nil {
		return []dto.FilmResponse{}, dto.ErrGetAllGenre
	}

	var filmResponses []dto.FilmResponse

	for _, film := range films {
		var fileResponses []dto.FilmGambarResponse
		var genreResponses []dto.GenreResponse

		rating, _ := s.reviewRepository.GetRatingByFilmID(ctx, film.ID)
		rating = math.Round(rating*100) / 100

		for _, file := range film.FilmGambar {
			fileResponses = append(fileResponses, dto.FilmGambarResponse{
				ID:  file.ID,
				Url: file.Url,
			})
		}

		for _, genre := range film.FilmGenre {
			genreResponses = append(genreResponses, dto.GenreResponse{
				ID:   genre.Genre.ID,
				Nama: genre.Genre.Nama,
			})
		}

		formattedDate := utils.FormatDate(film.TanggalRilis)
		filmResponses = append(filmResponses, dto.FilmResponse{
			ID:           film.ID,
			Judul:        film.Judul,
			TanggalRilis: formattedDate,
			Durasi:       film.Durasi,
			Status:       film.Status,
			Rating:       rating,
			Gambar:       fileResponses,
			Genres:       genreResponses,
		})
	}

	return filmResponses, nil
}

func (s *filmService) GetFilmByID(ctx context.Context, id int) (dto.FilmResponse, error) {
	film, err := s.filmRepository.GetFilmByID(ctx, id)
	if err != nil {
		return dto.FilmResponse{}, dto.ErrGetFilmByID
	}

	rating, _ := s.reviewRepository.GetRatingByFilmID(ctx, film.ID)
	rating = math.Round(rating*100) / 100

	var fileResponses []dto.FilmGambarResponse
	var genreResponses []dto.GenreResponse
	for _, file := range film.FilmGambar {
		fileResponses = append(fileResponses, dto.FilmGambarResponse{
			ID:  file.ID,
			Url: file.Url,
		})
	}

	for _, genre := range film.FilmGenre {
		genreResponses = append(genreResponses, dto.GenreResponse{
			ID:   genre.Genre.ID,
			Nama: genre.Genre.Nama,
		})
	}

	formattedDate := utils.FormatDate(film.TanggalRilis)
	filmResponse := dto.FilmResponse{
		ID:           film.ID,
		Judul:        film.Judul,
		Sinopsis:     film.Sinopsis,
		Sutradara:    film.Sutradara,
		TanggalRilis: formattedDate,
		TotalEpisode: film.TotalEpisode,
		Durasi:       film.Durasi,
		Status:       film.Status,
		Rating:       rating,
		Gambar:       fileResponses,
		Genres:       genreResponses,
	}

	return filmResponse, nil
}

func (s *filmService) CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmResponse, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return dto.FilmResponse{}, tx.Error
	}

	film := entity.Film{
		Judul:        filmReq.Judul,
		Sinopsis:     filmReq.Sinopsis,
		Sutradara:    filmReq.Sutradara,
		Status:       filmReq.Status,
		Durasi:       filmReq.Durasi,
		TotalEpisode: filmReq.TotalEpisode,
		TanggalRilis: filmReq.TanggalRilis,
	}

	createdFilm, err := s.filmRepository.CreateFilm(ctx, tx, film)
	var filmGambarResponse []dto.FilmGambarResponse
	var genreResponse []dto.GenreResponse

	if err != nil {
		tx.Rollback()
		return dto.FilmResponse{}, err
	}

	for _, genreID := range filmReq.Genre {
		filmGenre := entity.FilmGenre{
			FilmID:  createdFilm.ID,
			GenreID: genreID,
		}
		genre, err := s.filmGenreRepository.CreateFilmGenre(ctx, tx, filmGenre)
		genreResponse = append(genreResponse, dto.GenreResponse{
			ID:   genre.Genre.ID,
			Nama: genre.Genre.Nama,
		})
		if err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, err
		}
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, err
		}

		uniqueName := utils.GenerateUniqueImageName(createdFilm.Judul, file.Filename)
		uploadResult, err := s.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{
			Folder:   "ReviewFilem",
			PublicID: uniqueName,
		})
		src.Close()

		if err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, err
		}

		filmGambar := entity.FilmGambar{
			FilmID: createdFilm.ID,
			Url:    uploadResult.SecureURL,
		}

		if err := s.filmGambarRepository.Save(ctx, tx, filmGambar); err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, err
		}

		filmGambarResponse = append(filmGambarResponse, dto.FilmGambarResponse{
			ID:  filmGambar.ID,
			Url: filmGambar.Url,
		})
	}

	// Commit transaction jika semua berhasil
	if err := tx.Commit().Error; err != nil {
		return dto.FilmResponse{}, err
	}

	formattedDate := utils.FormatDate(createdFilm.TanggalRilis)
	return dto.FilmResponse{
		ID:           createdFilm.ID,
		Judul:        createdFilm.Judul,
		Sinopsis:     createdFilm.Sinopsis,
		Sutradara:    createdFilm.Sutradara,
		Status:       createdFilm.Status,
		Durasi:       createdFilm.Durasi,
		TotalEpisode: createdFilm.TotalEpisode,
		TanggalRilis: formattedDate,
		Rating:       0,
		Gambar:       filmGambarResponse,
		Genres:       genreResponse,
	}, nil
}

func (s *filmService) UpdateFilm(ctx context.Context, film dto.UpdateFilmRequest) (entity.Film, error) {
	updatedFilm, err := s.filmRepository.UpdateFilm(ctx, film)
	if err != nil {
		return entity.Film{}, dto.ErrUpdateFilm
	}

	return updatedFilm, nil
}

func (s *filmService) DeleteFilm(ctx context.Context, id int) error {
	err := s.filmRepository.DeleteFilm(ctx, id)
	if err != nil {
		return dto.ErrDeleteFilm
	}

	return nil
}

func (s *filmService) UpdateStatus(ctx context.Context, id int, req dto.UpdateStatusFilmRequest) error {
	err := s.filmRepository.UpdateStatus(ctx, id, req.Status)
	if err != nil {
		return dto.ErrUpdateStatusFilm
	}

	return nil
}

func (s *filmService) SearchFilm(ctx context.Context, req dto.SearchFilmRequest) ([]dto.FilmResponse, error) {
	films, err := s.filmRepository.SearchFilm(ctx, req)
	if err != nil {
		return nil, err
	}

	var filmResponses []dto.FilmResponse

	for _, film := range films {
		var fileResponses []dto.FilmGambarResponse
		var genreResponses []dto.GenreResponse

		rating, _ := s.reviewRepository.GetRatingByFilmID(ctx, film.ID)
		rating = math.Round(rating*100) / 100

		for _, file := range film.FilmGambar {
			fileResponses = append(fileResponses, dto.FilmGambarResponse{
				ID:  file.ID,
				Url: file.Url,
			})
		}

		for _, genre := range film.FilmGenre {
			genreResponses = append(genreResponses, dto.GenreResponse{
				ID:   genre.Genre.ID,
				Nama: genre.Genre.Nama,
			})
		}

		formattedDate := utils.FormatDate(film.TanggalRilis)
		filmResponses = append(filmResponses, dto.FilmResponse{
			ID:           film.ID,
			Judul:        film.Judul,
			TanggalRilis: formattedDate,
			Durasi:       film.Durasi,
			Status:       film.Status,
			Rating:       rating,
			Gambar:       fileResponses,
			Genres:       genreResponses,
		})
	}

	return filmResponses, nil
}
