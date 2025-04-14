package dto

import "errors"

const (
	// failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_GENRE_NOT_FOUND = "genre not found"
	MESSAGE_FAILED_PROSES_REQUEST = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS = "denied access"

	// success message
	MESSAGE_GENRE_CREATED = "genre created successfully"
	MESSAGE_GENRE_UPDATED = "genre updated successfully"
	MESSAGE_SUCCESS_GET_LIST_GENRE = "success get list genre"
)

var (
	ErrGetAllGenre = errors.New("failed to get all genre")
	ErrCreateGenre = errors.New("failed to create genre")
	ErrUpdateGenre = errors.New("failed to update genre")
)

