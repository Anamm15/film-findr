package dto

import "errors"

const (
	MESSAGE_FAILED_UPDATED_STATUS_USER_FILM = "Failed to update status user film"
	MESSAGE_FAILED_CREATED_USER_FILM        = "Failed to add film"
	MESSAGE_FAILED_UPDATED_USER_FILM        = "Failed to update status film"
	MESSAGE_FAILED_GET_USER_FILM            = "Failed to get user film"

	MESSAGE_SUCCESS_UPDATED_STATUS_USER_FILM = "Status user film updated successfully"
	MESSAGE_SUCCESS_CREATED_USER_FILM        = "Success add film"
	MESSAGE_SUCCESS_UPDATED_USER_FILM        = "Status film updated successfully"
	MESSAGE_SUCCESS_GET_USER_FILM            = "Success get user film"
)

var (
	ErrGetUserFilm           = errors.New("Failed to get user film")
	ErrCreateUserFilm        = errors.New("Failed to create user film")
	ErrUpdateStatusUserFilm  = errors.New("Failed to update status user film")
	ErrCheckUserFilm         = errors.New("Failed to check user film")
	ErrStatusFilmNotYetAired = errors.New("Film with not yet aired must be plan to watch")
)

type (
	UserFilmCreateRequest struct {
		UserID int    `json:"user_id" validate:"required"`
		FilmID int    `json:"film_id" validate:"required"`
		Status string `json:"status" validate:"required"`
	}

	UserFilmUpdateStatusRequest struct {
		ID     int    `json:"id" validate:"required"`
		Status string `json:"status" validate:"required"`
		FilmID int    `json:"film_id" validate:"required"`
	}

	UserFilmResponse struct {
		ID     int          `json:"id"`
		Status string       `json:"status"`
		Film   FilmResponse `json:"film"`
	}

	GetUserFilmResponse struct {
		UserFilms []UserFilmResponse `json:"user_films"`
		CountPage int                `json:"count_page"`
	}
)
