package controller

import (
	"strconv"

	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type FilmController interface {
	GetAllFilm(ctx *gin.Context)
	GetFilmById(ctx *gin.Context)
	CreateFilm(ctx *gin.Context)
	UpdateFilm(ctx *gin.Context)
	DeleteFilm(ctx *gin.Context)
	AddFilmGenre(ctx *gin.Context)
	DeleteFilmGenre(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
	SearchFilm(ctx *gin.Context)
}

type filmController struct {
	filmService      service.FilmService
	filmGenreService service.FilmGenreService
}

func NewFilmController(filmService service.FilmService, filmGenreService service.FilmGenreService) FilmController {
	return &filmController{
		filmService:      filmService,
		filmGenreService: filmGenreService,
	}
}

func (s *filmController) GetAllFilm(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	films, err := s.filmService.GetAllFilm(ctx, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_FILM, films, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) GetFilmById(ctx *gin.Context) {
	id := ctx.Param("id")
	film, err := s.filmService.GetFilmByID(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed("Failed to get film by id", err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_FILM, film, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) CreateFilm(ctx *gin.Context) {
	var film dto.CreateFilmRequest
	if err := ctx.ShouldBind(&film); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INTERNAL_ERROR, err.Error(), nil)
		ctx.JSON(dto.STATUS_INTERNAL_SERVER_ERROR, res)
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_IMAGE_NOT_FOUND, dto.ErrGetImageRequest.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdFilm, err := s.filmService.CreateFilm(ctx, film, files)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_FILM, createdFilm, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) UpdateFilm(ctx *gin.Context) {
	var film dto.UpdateFilmRequest
	if err := ctx.ShouldBindJSON(&film); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	updatedFilm, err := s.filmService.UpdateFilm(ctx, film)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_FILM, updatedFilm, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) DeleteFilm(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.filmService.DeleteFilm(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_FILM, nil, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) AddFilmGenre(ctx *gin.Context) {
	var filmGenre dto.FilmGenreRequest
	if err := ctx.ShouldBindJSON(&filmGenre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdFilmGenre, err := s.filmGenreService.CreateFilmGenre(ctx, filmGenre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_FILM_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_FILM_GENRE, createdFilmGenre, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) DeleteFilmGenre(ctx *gin.Context) {
	var filmGenre dto.FilmGenreRequest
	if err := ctx.ShouldBindJSON(&filmGenre); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := s.filmGenreService.DeleteFilmGenre(ctx, filmGenre)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_FILM_GENRE, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_FILM_GENRE, nil, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var req dto.UpdateStatusFilmRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := s.filmService.UpdateStatus(ctx, utils.StringToInt(id), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_STATUS_FIlM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_STATUS_FILM, nil, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (s *filmController) SearchFilm(c *gin.Context) {
	keyword := c.Query("keyword")
	status := c.Query("status")
	genreIDs := c.QueryArray("genre_ids")

	var genresId []int
	if len(genreIDs) > 0 {
		genresId = make([]int, len(genreIDs))
		for i, genreID := range genreIDs {
			genresId[i] = utils.StringToInt(genreID)
		}
	}

	var req dto.SearchFilmRequest
	req = dto.SearchFilmRequest{
		Keyword: keyword,
		Status:  status,
		Genres:  genresId,
	}

	films, err := s.filmService.SearchFilm(c, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEARCH_FILM, err.Error(), nil)
		c.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	if len(films) == 0 {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_FILM_NOT_FOUND, dto.MESSAGE_FAILED_FILM_NOT_FOUND, nil)
		c.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_SEARCH_FILM, films, nil)
	c.JSON(dto.STATUS_OK, res)
}
