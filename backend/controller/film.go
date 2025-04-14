package controller

import (
	"ReviewPiLem/service"

	"ReviewPiLem/dto"
	"ReviewPiLem/utils"
	"ReviewPiLem/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FilmController interface {
	GetAllFilm(ctx *gin.Context)
	GetFilmById(ctx *gin.Context)
	CreateFilm(ctx *gin.Context)
	UpdateFilm(ctx *gin.Context)
	DeleteFilm(ctx *gin.Context)
}

type filmController struct {
	filmService service.FilmService
}

func NewFilmController(filmService service.FilmService) FilmController {
	return &filmController{filmService: filmService}
}

func (s *filmController) GetAllFilm(ctx *gin.Context) {
	films, err := s.filmService.GetAllFilm(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get all film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_FILM, films)
	ctx.JSON(http.StatusOK, res)
}

func (s *filmController) GetFilmById(ctx *gin.Context) {
	id := ctx.Param("id")
	film, err := s.filmService.GetFilmByID(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get film by id", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_FILM, film)
	ctx.JSON(http.StatusOK, res)
}

func (s *filmController) CreateFilm(ctx *gin.Context) {
	var film entity.Film
	if err := ctx.ShouldBindJSON(&film); err != nil {
		res := utils.BuildResponseFailed("Failed to create film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}	

	film, err := s.filmService.CreateFilm(ctx, film)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_FILM, film)
	ctx.JSON(http.StatusOK, res)
}

func (s *filmController) UpdateFilm(ctx *gin.Context) {
	var film entity.Film
	if err := ctx.ShouldBindJSON(&film); err != nil {
		res := utils.BuildResponseFailed("Failed to update film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	film, err := s.filmService.UpdateFilm(ctx, film)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_FILM, film)
	ctx.JSON(http.StatusOK, res)
}

func (s *filmController) DeleteFilm(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.filmService.DeleteFilm(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to delete film", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_FILM, nil)
	ctx.JSON(http.StatusOK, res)
}