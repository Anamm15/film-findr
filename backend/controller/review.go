package controller

import (
	"net/http"
	"strconv"

	"ReviewPiLem/dto"
	"ReviewPiLem/service"
	"ReviewPiLem/utils"

	"github.com/gin-contrib/sessions"
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
}

func NewReviewController(reviewService service.ReviewService) ReviewController {
	return &reviewController{
		reviewService: reviewService,
	}
}

func (c *reviewController) GetReviewByUserId(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	userIdSession := session.Get("user_id")
	userId, ok := userIdSession.(int)
	if !ok {
		res := utils.BuildResponseFailed("Gagal mengambil session user", "user_id tidak ditemukan atau bukan string", nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	review, err := c.reviewService.GetReviewByUserId(ctx, utils.StringToInt(id), userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_REVIEW, review)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) GetReviewByFilmId(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id := ctx.Param("id")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	userIdSession := session.Get("user_id")
	userId, ok := userIdSession.(int)
	if !ok {
		res := utils.BuildResponseFailed("Gagal mengambil session user", "user_id tidak ditemukan atau bukan string", nil)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	reviews, err := c.reviewService.GetReviewByFilmId(ctx, utils.StringToInt(id), userId, page)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_REVIEW, reviews)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	var userId int
	userId = ctx.MustGet("user_id").(int)

	var review dto.CreateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdReview, err := c.reviewService.CreateReview(ctx, review, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_REVIEW, createdReview)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) UpdateReview(ctx *gin.Context) {
	var review dto.UpdateReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.reviewService.UpdateReview(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) UpdateReaksiReview(ctx *gin.Context) {
	var review dto.UpdateReaksiReviewRequest
	if err := ctx.ShouldBindJSON(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.reviewService.UpdateReaksiReview(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_REVIEW, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.reviewService.DeleteReview(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_REVIEW, nil)
	ctx.JSON(http.StatusOK, res)
}
