package dto

import "errors"

const (
	// failed message
	MESSAGE_FAILED_FILM_NOT_FOUND = "film not found"
	MESSAGE_FAILED_CREATED_FILM = "failed created film"
	MESSAGE_FAILED_UPDATED_FILM = "failed updated film"
	MESSAGE_FAILED_DELETED_FILM = "failed deleted film"

	// success message
	MESSAGE_SUCCESS_GET_LIST_FILM = "success get list film" 
	MESSAGE_SUCCESS_GET_FILM = "success get film"
	MESSAGE_SUCCESS_CREATED_FILM = "film created successfully"
	MESSAGE_SUCCESS_UPDATED_FILM = "film updated successfully"
	MESSAGE_SUCCESS_DELETED_FILM  = "film deleted successfully"
)

var (
	ErrGetAllFilm = errors.New("failed to get all film")
	ErrGetFilmByID = errors.New("failed to get film by id")
	ErrCreateFilm = errors.New("failed to create film")
	ErrUpdateFilm = errors.New("failed to update film")
	ErrDeleteFilm = errors.New("failed to delete film")
)