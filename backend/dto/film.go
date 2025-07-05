package dto

import (
	"errors"
	"time"
)

const (
	// failed message
	MESSAGE_FAILED_FILM_NOT_FOUND      = "Film not found"
	MESSAGE_FAILED_GET_ALL_FILM        = "Failed get all film"
	MESSAGE_FAILED_CREATED_FILM        = "Failed created film"
	MESSAGE_FAILED_UPDATED_FILM        = "Failed updated film"
	MESSAGE_FAILED_DELETED_FILM        = "Failed deleted film"
	MESSAGE_FAILED_UPDATED_STATUS_FIlM = "Failed to update status film"
	MESSAGE_FAILED_SEARCH_FILM         = "Failed to search film"

	// success message
	MESSAGE_SUCCESS_GET_LIST_FILM       = "Success get list film"
	MESSAGE_SUCCESS_GET_FILM            = "Success get film"
	MESSAGE_SUCCESS_CREATED_FILM        = "Film created successfully"
	MESSAGE_SUCCESS_UPDATED_FILM        = "Film updated successfully"
	MESSAGE_SUCCESS_DELETED_FILM        = "Film deleted successfully"
	MESSAGE_SUCCESS_UPDATED_STATUS_FILM = "Status film updated successfully"
	MESSAGE_SUCCESS_SEARCH_FILM         = "Success search film"
)

var (
	ErrGetAllFilm       = errors.New("Failed to get all film")
	ErrGetFilmByID      = errors.New("Failed to get film by id")
	ErrCreateFilm       = errors.New("Failed to create film")
	ErrUpdateFilm       = errors.New("Failed to update film")
	ErrDeleteFilm       = errors.New("Failed to delete film")
	ErrUpdateStatusFilm = errors.New("Failed to update status film")
	ErrGetImageRequest  = errors.New("No image files found")
)

type (
	CreateFilmRequest struct {
		Judul        string    `form:"judul" validate:"required"`
		Sinopsis     string    `form:"sinopsis" validate:"required"`
		Sutradara    string    `form:"sutradara" validate:"required"`
		Status       string    `form:"status" validate:"required"`
		Durasi       int       `form:"durasi" validate:"required"`
		TotalEpisode int       `form:"total_episode" validate:"required"`
		TanggalRilis time.Time `form:"tanggal_rilis" time_format:"2006-01-02" validate:"required"`
		Genre        []int     `form:"genres" validate:"required"`
	}

	FilmGambarResponse struct {
		ID  int    `json:"id"`
		Url string `json:"url"`
	}

	FilmResponse struct {
		ID           int                  `json:"id"`
		Judul        string               `json:"judul"`
		Sinopsis     string               `json:"sinopsis"`
		Sutradara    string               `json:"sutradara"`
		Status       string               `json:"status"`
		Durasi       int                  `json:"durasi"`
		TotalEpisode int                  `json:"total_episode"`
		TanggalRilis string               `json:"tanggal_rilis"`
		Rating       float64              `json:"rating"`
		Gambar       []FilmGambarResponse `json:"film_gambar"`
		Genres       []GenreResponse      `json:"genres"`
	}

	UpdateFilmRequest struct {
		ID           int       `json:"id" validate:"required"`
		Judul        string    `json:"judul" validate:"required"`
		Sinopsis     string    `json:"sinopsis" validate:"required"`
		Sutradara    string    `json:"sutradara" validate:"required"`
		Status       string    `json:"status" validate:"required"`
		Durasi       int       `json:"durasi" validate:"required"`
		TotalEpisode int       `json:"total_episode" validate:"required"`
		TanggalRilis time.Time `json:"tanggal_rilis" time_format:"2006-01-02" validate:"required"`
		Genre        []int     `json:"genres" validate:"required"`
	}

	UpdateStatusFilmRequest struct {
		Status string `json:"status" binding:"required"`
	}

	SearchFilmRequest struct {
		Keyword string
		Status  string
		Genres  []int
	}
)
