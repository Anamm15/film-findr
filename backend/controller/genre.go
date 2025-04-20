package controller

import (
	"net/http"

	"ReviewPiLem/dto"
	"ReviewPiLem/service"
	"ReviewPiLem/utils"

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
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_GENRE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_GENRE, genres)
	ctx.JSON(http.StatusOK, res)
}

func (s *genreController) CreateGenre(ctx *gin.Context) {
	var genre dto.GenreRequest
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdGenre, err := s.genreService.CreateGenre(ctx, genre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_GENRE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_CREATED, createdGenre)
	ctx.JSON(http.StatusOK, res)
}

func (s *genreController) UpdateGenre(ctx *gin.Context) {
	genreId := ctx.Param("id")
	var genre dto.GenreRequest
	genre.ID = utils.StringToInt(genreId)
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	updatedGenre, err := s.genreService.UpdateGenre(ctx, genre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_GENRE, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_GENRE_UPDATED, updatedGenre)
	ctx.JSON(http.StatusOK, res)
}
