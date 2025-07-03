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
	reviewService  service.ReviewService
	sessionService service.SessionService
}

func NewReviewController(
	reviewService service.ReviewService,
	sessionService service.SessionService,
) ReviewController {
	return &reviewController{
		reviewService:  reviewService,
		sessionService: sessionService,
	}
}

func (c *reviewController) GetReviewByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	userId, err := c.sessionService.GetUserID(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SESSION_EXPIRED, dto.MESSAGE_FAILED_SESSION_EXPIRED, nil)
		ctx.JSON(dto.STATUS_UNAUTHORIZED, res)
		return
	}

	review, err := c.reviewService.GetReviewByUserId(ctx, utils.StringToInt(id), userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEW, review)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) GetReviewByFilmId(ctx *gin.Context) {
	id := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	userId, err := c.sessionService.GetUserID(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SESSION_EXPIRED, dto.MESSAGE_FAILED_SESSION_EXPIRED, nil)
		ctx.JSON(dto.STATUS_UNAUTHORIZED, res)
		return
	}

	reviews, err := c.reviewService.GetReviewByFilmId(ctx, utils.StringToInt(id), userId, page)
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
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	var review dto.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := c.reviewService.UpdateReview(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) UpdateReaksiReview(ctx *gin.Context) {
	var review dto.UpdateReaksiReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REQUIRED_FIELD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	err := c.reviewService.UpdateReaksiReview(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.reviewService.DeleteReview(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_REVIEW, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_REVIEW, nil)
	ctx.JSON(dto.STATUS_OK, res)
}
