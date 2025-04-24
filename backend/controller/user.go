package controller

import (
	"net/http"

	"ReviewPiLem/dto"
	"ReviewPiLem/entity"
	"ReviewPiLem/service"
	"ReviewPiLem/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) GetAllUser(ctx *gin.Context) {
	users, err := c.userService.GetAllUser(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_USER, users)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	userResponse, err := c.userService.GetUserById(ctx, utils.StringToInt(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, userResponse)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserCreateRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	photoProfil, err := ctx.FormFile("photo_profil")
	var userResponse dto.UserResponse
	if err != nil {
		userResponse, err = c.userService.RegisterUser(ctx, user, nil)
	} else {
		userResponse, err = c.userService.RegisterUser(ctx, user, photoProfil)
	}

	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATED_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATED_USER, userResponse)
	ctx.JSON(http.StatusCreated, res)
}

func (c *userController) LoginUser(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID, user.Role)
	userResponse := entity.Authorization{
		Token: token,
		Role:  user.Role,
	}

	session := sessions.Default(ctx)
	session.Set("user_id", user.ID)
	session.Save()

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) LogoutUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user_id")

	if userID == nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_USER_NOT_LOGIN, dto.ErrUserNotLogin.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	session.Clear()
	if err := session.Save(); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGOUT, dto.ErrFailedSaveSession.Error(), err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGOUT, nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(int)

	var user dto.UserUpdateRequest
	user.ID = userId
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := c.userService.UpdateUser(ctx, user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATED_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATED_USER, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	err := c.userService.DeleteUser(ctx, utils.StringToInt(userId))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETED_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETED_USER, nil)
	ctx.JSON(http.StatusOK, res)
}
