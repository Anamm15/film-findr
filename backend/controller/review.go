package controller

import (
	"strconv"

	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type ReviewController interface {
	GetReviewByUserId(ctx *gin.Context)
	GetReviewByFilmId(ctx *gin.Context)
	CreateReview(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	UpdateReaksiReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}

type reviewController struct {
	reviewService service.ReviewService
	jwtService    service.JWTService
}

func NewReviewController(
	reviewService service.ReviewService,
	jwtService service.JWTService,
) ReviewController {
	return &reviewController{
		reviewService: reviewService,
		jwtService:    jwtService,
	}
}

func (c *reviewController) GetReviewByUserId(ctx *gin.Context) {
	userReviewIdParam := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	userReviewId, err := strconv.Atoi(userReviewIdParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_PARAMETER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	tokenStr, _ := ctx.Cookie("access_token")
	userId, _, _ := c.jwtService.GetDataByToken(tokenStr)
	reviews, err := c.reviewService.GetReviewByUserId(ctx, userReviewId, userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEW, reviews)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) GetReviewByFilmId(ctx *gin.Context) {
	filmIdParam := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	filmId, err := strconv.Atoi(filmIdParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_PARAMETER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	tokenStr, _ := ctx.Cookie("access_token")
	userId, _, _ := c.jwtService.GetDataByToken(tokenStr)
	reviews, err := c.reviewService.GetReviewByFilmId(ctx, filmId, userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_REVIEW, reviews)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	var userId int
	userId = ctx.MustGet("user_id").(int)

	var review dto.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	createdReview, err := c.reviewService.CreateReview(ctx, review, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_REVIEW, createdReview)
	ctx.JSON(dto.STATUS_CREATED, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	reviewIdParam := ctx.Param("id")
	reviewId, err := strconv.Atoi(reviewIdParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_PARAMETER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	var review dto.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err = c.reviewService.UpdateReview(ctx, reviewId, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) UpdateReaksiReview(ctx *gin.Context) {
	reaksiReviewIdParam := ctx.Param("id")
	reaksiReviewId, err := strconv.Atoi(reaksiReviewIdParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_PARAMETER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	var review dto.UpdateReaksiReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err = c.reviewService.UpdateReaksiReview(ctx, reaksiReviewId, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	reviewIdParam := ctx.Param("id")
	reviewId, err := strconv.Atoi(reviewIdParam)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_INVALID_PARAMETER, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err = c.reviewService.DeleteReview(ctx, reviewId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
