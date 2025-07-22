package controller

import (
	"FilmFindr/dto"
	"FilmFindr/service"
	"FilmFindr/utils"

	"github.com/gin-gonic/gin"
)

type DashboardController interface {
	GetDashboard(ctx *gin.Context)
}

type dashboardController struct {
	dashboardService service.DashboardService
	jwtService       service.JWTService
}

func NewDashboardController(dashboardService service.DashboardService, jwtService service.JWTService) DashboardController {
	return &dashboardController{dashboardService: dashboardService, jwtService: jwtService}
}

func (c *dashboardController) GetDashboard(ctx *gin.Context) {
	var dashboardResponse dto.GetDashboardResponse
	dashboardResponse, err := c.dashboardService.GetDashboard(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DASHBOARD, err.Error(), nil)
		ctx.JSON(dto.STATUS_BAD_REQUEST, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_DASHBOARD, dashboardResponse)
	ctx.JSON(dto.STATUS_OK, res)
}
