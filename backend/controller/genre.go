package controller

import (
	"ReviewPiLem/dto"
	"ReviewPiLem/utils"
	"ReviewPiLem/entity"
	"ReviewPiLem/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type GenreController interface {
	GetAllGenre(ctx *gin.Context)
	CreateGenre(ctx *gin.Context)
	UpdateGenre(ctx *gin.Context)
}

type genreController struct {
	genreService service.GenreService
}

func NewGenreController(genreService service.GenreService) GenreController {
	return &genreController{
		genreService: genreService,
	}
}

func (s *genreController) GetAllGenre(ctx *gin.Context) {
	genres, err := s.genreService.GetAllGenre(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get all genres", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_GENRE, genres)
	ctx.JSON(http.StatusOK, res)
}

func (s *genreController) CreateGenre(ctx *gin.Context) {
	var genre entity.Genre
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed("Failed to create genre", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	genre, err := s.genreService.CreateGenre(ctx, genre)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to create genre", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_CREATED, genre)
	ctx.JSON(http.StatusOK, res)
}

func (s *genreController) UpdateGenre(ctx *gin.Context) {
	var genre entity.Genre
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed("Failed to update genre", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	genre, err := s.genreService.UpdateGenre(ctx, genre)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to update genre", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_UPDATED, genre)
	ctx.JSON(http.StatusOK, res)
}