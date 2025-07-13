package service

import (
	"context"
	"math"
	"mime/multipart"

	"FilmFindr/dto"
	"FilmFindr/entity"
	"FilmFindr/helpers"
	"FilmFindr/repository"
	"FilmFindr/utils"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

type FilmService interface {
	GetAllFilm(ctx context.Context, page int) (dto.GetFilmResponse, error)
	GetFilmByID(ctx context.Context, id int) (dto.FilmResponse, error)
	CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmResponse, error)
	UpdateFilm(ctx context.Context, film dto.UpdateFilmRequest) (entity.Film, error)
	DeleteFilm(ctx context.Context, id int) error
	UpdateStatus(ctx context.Context, id int, req dto.UpdateStatusFilmRequest) error
	SearchFilm(ctx context.Context, req dto.SearchFilmRequest, page int) (dto.GetFilmResponse, error)
	GetTopFilm(ctx context.Context) ([]dto.TopFilm, error)
	GetTrendingFilm(ctx context.Context) ([]dto.TrendingFilm, error)
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

func (s *filmService) GetAllFilm(ctx context.Context, page int) (dto.GetFilmResponse, error) {
	offset := (page - 1) * helpers.LIMIT_FILM

	countFilm, err := s.filmRepository.CountFilm(ctx)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	filmsFlat, err := s.filmRepository.GetAllFilm(ctx, offset)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	filmIDs := make([]int, len(filmsFlat))
	for i, f := range filmsFlat {
		filmIDs[i] = f.ID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	genreMap := make(map[int][]dto.GenreResponse)
	for _, g := range genres {
		genreMap[g.FilmID] = append(genreMap[g.FilmID], g)
	}

	gambarMap := make(map[int][]dto.FilmGambarResponse)
	for _, img := range gambar {
		gambarMap[img.FilmID] = append(gambarMap[img.FilmID], img)
	}

	var results []dto.FilmResponse
	for _, flat := range filmsFlat {
		formattedDate := utils.FormatDate(flat.TanggalRilis)
		results = append(results, dto.FilmResponse{
			ID:           flat.ID,
			Judul:        flat.Judul,
			Sinopsis:     flat.Sinopsis,
			Sutradara:    flat.Sutradara,
			Status:       flat.Status,
			Durasi:       flat.Durasi,
			TotalEpisode: flat.TotalEpisode,
			TanggalRilis: formattedDate,
			Rating:       math.Round(flat.Rating*100) / 100,
			Genres:       genreMap[flat.ID],
			Gambar:       gambarMap[flat.ID],
		})
	}

	totalPage := int(math.Ceil(float64(countFilm) / float64(helpers.LIMIT_FILM)))
	GetFilmResponses := dto.GetFilmResponse{
		Film:      results,
		CountPage: totalPage,
	}
	return GetFilmResponses, nil
}

func (s *filmService) GetFilmByID(ctx context.Context, id int) (dto.FilmResponse, error) {
	filmFlat, err := s.filmRepository.GetFilmByID(ctx, id)
	if err != nil {
		return dto.FilmResponse{}, dto.ErrGetFilm
	}

	filmIDs := []int{filmFlat.ID}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.FilmResponse{}, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return dto.FilmResponse{}, err
	}

	genreMap := make(map[int][]dto.GenreResponse)
	for _, g := range genres {
		genreMap[g.FilmID] = append(genreMap[g.FilmID], g)
	}

	gambarMap := make(map[int][]dto.FilmGambarResponse)
	for _, img := range gambar {
		gambarMap[img.FilmID] = append(gambarMap[img.FilmID], img)
	}

	formattedDate := utils.FormatDate(filmFlat.TanggalRilis)
	filmResponse := dto.FilmResponse{
		ID:           filmFlat.ID,
		Judul:        filmFlat.Judul,
		Sinopsis:     filmFlat.Sinopsis,
		Sutradara:    filmFlat.Sutradara,
		TanggalRilis: formattedDate,
		TotalEpisode: filmFlat.TotalEpisode,
		Durasi:       filmFlat.Durasi,
		Status:       filmFlat.Status,
		Rating:       filmFlat.Rating,
		Gambar:       gambarMap[filmFlat.ID],
		Genres:       genreMap[filmFlat.ID],
	}

	return filmResponse, nil
}

func (s *filmService) CreateFilm(ctx context.Context, filmReq dto.CreateFilmRequest, files []*multipart.FileHeader) (dto.FilmResponse, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return dto.FilmResponse{}, dto.ErrCreateFilm
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
		return dto.FilmResponse{}, dto.ErrCreateFilm
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
			return dto.FilmResponse{}, dto.ErrCreateFilm
		}
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, dto.ErrCreateFilm
		}

		uniqueName := utils.GenerateUniqueImageName(createdFilm.Judul, file.Filename)
		uploadResult, err := s.cloudinary.Upload.Upload(ctx, src, uploader.UploadParams{
			Folder:   "ReviewFilem",
			PublicID: uniqueName,
		})
		src.Close()

		if err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, dto.ErrFailedUploadFile
		}

		filmGambar := entity.FilmGambar{
			FilmID: createdFilm.ID,
			Url:    uploadResult.SecureURL,
		}

		if err := s.filmGambarRepository.Save(ctx, tx, filmGambar); err != nil {
			tx.Rollback()
			return dto.FilmResponse{}, dto.ErrCreateFilm
		}

		filmGambarResponse = append(filmGambarResponse, dto.FilmGambarResponse{
			ID:  filmGambar.ID,
			Url: filmGambar.Url,
		})
	}

	if err := tx.Commit().Error; err != nil {
		return dto.FilmResponse{}, dto.ErrCreateFilm
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

func (s *filmService) SearchFilm(ctx context.Context, req dto.SearchFilmRequest, page int) (dto.GetFilmResponse, error) {
	countFilm, err := s.filmRepository.CountFilm(ctx)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	films, err := s.filmRepository.SearchFilm(ctx, req, page)
	if err != nil {
		return dto.GetFilmResponse{}, err
	}

	var filmResponses []dto.FilmResponse

	for _, film := range films {
		var fileResponses []dto.FilmGambarResponse
		var genreResponses []dto.GenreResponse

		rating, _ := s.reviewRepository.GetRatingFromMaterializedView(ctx, film.ID)
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

	totalPage := int(math.Ceil(float64(countFilm) / float64(helpers.LIMIT_FILM)))
	getFilmResponses := dto.GetFilmResponse{
		CountPage: totalPage,
		Film:      filmResponses,
	}
	return getFilmResponses, nil
}

func (s *filmService) GetTopFilm(ctx context.Context) ([]dto.TopFilm, error) {
	topFilmsFlat, err := s.filmRepository.GetTopFilm(ctx)
	if err != nil {
		return nil, err
	}

	filmIDs := make([]int, len(topFilmsFlat))
	for i, f := range topFilmsFlat {
		filmIDs[i] = f.FilmID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	genreMap := make(map[int][]dto.GenreResponse)
	for _, g := range genres {
		genreMap[g.FilmID] = append(genreMap[g.FilmID], g)
	}

	gambarMap := make(map[int][]dto.FilmGambarResponse)
	for _, img := range gambar {
		gambarMap[img.FilmID] = append(gambarMap[img.FilmID], img)
	}

	var results []dto.TopFilm
	for _, flat := range topFilmsFlat {
		results = append(results, dto.TopFilm{
			ID:           flat.FilmID,
			Judul:        flat.Judul,
			Sinopsis:     flat.Sinopsis,
			Sutradara:    flat.Sutradara,
			Status:       flat.Status,
			Durasi:       flat.Durasi,
			TotalEpisode: flat.TotalEpisode,
			TanggalRilis: flat.TanggalRilis,
			Rating:       math.Round(flat.Rating*100) / 100,
			Genres:       genreMap[flat.FilmID],
			Gambar:       gambarMap[flat.FilmID],
		})
	}

	return results, nil
}

func (s *filmService) GetTrendingFilm(ctx context.Context) ([]dto.TrendingFilm, error) {
	trendingFilmFlat, err := s.filmRepository.GetTrendingFilm(ctx)
	if err != nil {
		return nil, err
	}

	filmIDs := make([]int, len(trendingFilmFlat))
	for i, f := range trendingFilmFlat {
		filmIDs[i] = f.FilmID
	}

	genres, err := s.filmGenreRepository.FindGenreByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	gambar, err := s.filmGambarRepository.FindFilmGambarByFilmIDs(ctx, filmIDs)
	if err != nil {
		return nil, err
	}

	genreMap := make(map[int][]dto.GenreResponse)
	for _, g := range genres {
		genreMap[g.FilmID] = append(genreMap[g.FilmID], g)
	}

	gambarMap := make(map[int][]dto.FilmGambarResponse)
	for _, img := range gambar {
		gambarMap[img.FilmID] = append(gambarMap[img.FilmID], img)
	}

	var results []dto.TrendingFilm
	for _, flat := range trendingFilmFlat {
		results = append(results, dto.TrendingFilm{
			ID:           flat.FilmID,
			Judul:        flat.Judul,
			Sinopsis:     flat.Sinopsis,
			Sutradara:    flat.Sutradara,
			Status:       flat.Status,
			Durasi:       flat.Durasi,
			TotalEpisode: flat.TotalEpisode,
			TanggalRilis: flat.TanggalRilis,
			Rating:       math.Round(flat.Rating*100) / 100,
			Genres:       genreMap[flat.FilmID],
			Gambar:       gambarMap[flat.FilmID],
		})
	}

	return results, nil
}
