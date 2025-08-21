package dto

import "errors"

const (
	// failed
	MESSAGE_FAILED_GET_LIST_GENRE = "Failed get list genre"
	MESSAGE_FAILED_GET_GENRE      = "Failed get genre"
	MESSAGE_FAILED_CREATED_GENRE  = "Failed created genre"
	MESSAGE_FAILED_UPDATED_GENRE  = "Failed updated genre"
	MESSAGE_FAILED_DELETED_GENRE  = "Failed deleted genre"

	// success message
	MESSAGE_SUCCESS_GET_LIST_GENRE = "Success get list genre"
	MESSAGE_GENRE_CREATED          = "Genre created successfully"
	MESSAGE_GENRE_UPDATED          = "Genre updated successfully"
)

var (
	ErrGetAllGenre = errors.New("Failed to get all genre")
	ErrCreateGenre = errors.New("Failed to create genre")
	ErrDeleteGenre = errors.New("Failed to delete genre")
)

type (
	GenreRequest struct {
		ID   int    `json:"id"`
		Nama string `json:"nama" validate:"required" binding:"required"`
	}

	GenreResponse struct {
		ID     int    `json:"id"`
		FilmID int    `json:"film_id"`
		Nama   string `json:"nama"`
	}

	GenreListAndCountResponse struct {
		Nama  string `json:"nama" gorm:"column:nama"`
		Count int64  `json:"count" gorm:"column:count"`
	}
)
