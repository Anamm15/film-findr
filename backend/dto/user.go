package dto

import (
	"errors"
)

const (
	// failed
	MESSAGE_USER_NOT_FOUND        = "User not found"
	MESSAGE_FAILED_GET_LIST_USER  = "Failed get list user"
	MESSAGE_FAILED_GET_USER       = "Failed get user"
	MESSAGE_FAILED_CREATED_USER   = "Failed created user"
	MESSAGE_FAILED_UPDATED_USER   = "Failed updated user"
	MESSAGE_FAILED_DELETED_USER   = "Failed deleted user"
	MESSAGE_FAILED_LOGIN          = "Failed to login"
	MESSAGE_FAILED_LOGOUT         = "Failed to logout"
	MESSAGE_FAILED_USER_NOT_LOGIN = "User not login"

	// success message
	MESSAGE_SUCCESS_GET_LIST_USER = "Success get list of user"
	MESSAGE_SUCCESS_GET_USER      = "Success get user"
	MESSAGE_SUCCESS_LOGIN         = "Login successfully"
	MESSAGE_SUCCESS_CREATED_USER  = "User created successfully"
	MESSAGE_SUCCESS_UPDATED_USER  = "User updated successfully"
	MESSAGE_SUCCESS_DELETED_USER  = "User deleted successfully"
	MESSAGE_SUCCESS_LOGOUT        = "Logout successfully"
)

var (
	ErrGetAllUser        = errors.New("Failed to get all user")
	ErrGetUserByID       = errors.New("Failed to get user by id")
	ErrCreateUser        = errors.New("Failed to create user")
	ErrCheckUsername     = errors.New("Failed to check username")
	ErrUpdateUser        = errors.New("Failed to update user")
	ErrDeleteUser        = errors.New("Failed to delete user")
	ErrFailedLogin       = errors.New("Failed to login")
	ErrPasswordNotMatch  = errors.New("Password not match")
	ErrEmailOrPassword   = errors.New("Email or password not match")
	ErrUserNotFound      = errors.New("User not found")
	ErrUserNotLogin      = errors.New("User not login")
	ErrFailedSaveSession = errors.New("Failed to save session")
)

type (
	UserCreateRequest struct {
		Nama     string `json:"nama" form:"nama" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
		Bio      string `json:"bio" form:"bio"`
	}

	UserLoginRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserFilm struct {
		ID     int          `json:"id"`
		Status string       `json:"status"`
		Film   FilmResponse `json:"film"`
	}

	UserRequest struct {
		ID          int    `json:"id"`
		Nama        string `json:"nama" form:"nama" binding:"required"`
		Username    string `json:"username" form:"username" binding:"required"`
		Bio         string `json:"bio" form:"bio"`
		PhotoProfil string `json:"photo_profil" form:"photo_profil"`
	}

	UserResponse struct {
		ID          int    `json:"id"`
		Nama        string `json:"nama"`
		Username    string `json:"username"`
		Bio         string `json:"bio"`
		PhotoProfil string `json:"photo_profil"`
	}

	UserUpdateRequest struct {
		ID       int    `json:"id"`
		Nama     string `json:"nama" form:"nama"`
		Bio      string `json:"bio" form:"bio"`
		Username string `json:"username" form:"username"`
	}

	UserResponseWithFilm struct {
		ID          int        `json:"id"`
		Nama        string     `json:"nama"`
		Username    string     `json:"username"`
		Bio         string     `json:"bio"`
		PhotoProfil string     `json:"photo_profil"`
		UserFilm    []UserFilm `json:"user_film"`
	}
)
