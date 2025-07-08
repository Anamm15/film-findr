package controller

import (
	"strconv"

	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type UserFilmController interface {
	GetUserFilmByUserId(ctx *gin.Context)
	CreateUserFilm(ctx *gin.Context)
	UpdateStatusUserFilm(ctx *gin.Context)
}

type userFilmController struct {
	userFilmService service.UserFilmService
}

func NewUserFilmController(userFilmService service.UserFilmService) UserFilmController {
	return &userFilmController{userFilmService: userFilmService}
}

func (u *userFilmController) GetUserFilmByUserId(ctx *gin.Context) {
	userIdParam := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	userId := utils.StringToInt(userIdParam)

	userFilms, err := u.userFilmService.GetUserFilmByUserId(ctx, userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER_FILM, userFilms)
	ctx.JSON(dto.STATUS_OK, res)
}

func (u *userFilmController) CreateUserFilm(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int)

	var userFilmReq dto.UserFilmCreateRequest
	userFilmReq.UserID = userId
	if err := ctx.ShouldBindJSON(&userFilmReq); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	userFilm, err := u.userFilmService.CreateUserFilm(ctx, userFilmReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_USER_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_USER_FILM, userFilm)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (u *userFilmController) UpdateStatusUserFilm(ctx *gin.Context) {
	var userFilmReq dto.UserFilmUpdateStatusRequest
	if err := ctx.ShouldBindJSON(&userFilmReq); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := u.userFilmService.UpdateStatusUserFilm(ctx, userFilmReq)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_USER_FILM, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_USER_FILM, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
