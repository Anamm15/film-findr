package dto

import "errors"

const (
	MESSAGE_FAILED_CREATED_FILM_GENRE = "Failed to create film genre"
	MESSAGE_FAILED_DELETED_FILM_GENRE = "Failed to delete film genre"

	MESSAGE_SUCCESS_CREATED_FILM_GENRE = "Film genre created successfully"
	MESSAGE_SUCCESS_DELETED_FILM_GENRE = "Film genre deleted successfully"
)

var (
	ErrCreateFilmGenre = errors.New("Failed to create film genre")
	ErrDeleteFilmGenre = errors.New("Failed to delete film genre")
)

type (
	FilmGenreRequest struct {
		FilmId  int `json:"film_id" validate:"required"`
		GenreId int `json:"genre_id" validate:"required"`
	}

	FilmGenreResponse struct {
		FilmId  int `json:"film_id"`
		GenreId int `json:"genre_id"`
	}
)
