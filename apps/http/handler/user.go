package handler

import (
	"net/http"

	"github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/service"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/labstack/echo/v4"
)

type (
	UserHandler interface {
		GetListUsers(c echo.Context) error
		UserSignUp(c echo.Context) error
		LoginUser(c echo.Context) error
		GetUserProfile(c echo.Context) error
		UpdateUserProfile(c echo.Context) error
		//DeleteUser(c echo.Context) error
	}

	userHandler struct {
		service.UserService
	}
)

// Signup User
// @Summary Get Sign up for user.
// @Description Sign up user process.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Failure 500 {object} utils.Error
// @Router /api/v1/auth/user/signup [post]
func (handler *userHandler) UserSignUp(c echo.Context) error {
	ctx := c.Request().Context()
	req := models.UserRegRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    utils.ErrBadQueryParams,
		})
	}

	createdUser, err := handler.UserService.UserSignUp(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Sever Error",
			Data:    utils.ErrInternalServerError,
		})
	}
	return c.JSON(http.StatusCreated, models.HttpResponse{
		Code:    http.StatusCreated,
		Message: "Admin Registered Successfully",
		Data:    createdUser,
	})
}

// UserSignIn godoc
// @Summary User SignIn
// @Description User SignIn
// @Tags User
// @Accept json
// @Produce json
// @Param signin body request.SignInReq true "SignIn user"
// @Success 200
// @Router /api/v1/auth/user/signin [post]
func (handler *userHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := models.UserLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Admin Failed",
			Data:    utils.ErrBadRequest,
		})
	}

	adminAuth, err := handler.UserService.LoginUser(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    utils.ErrInternalServerError,
		})
	}
	if adminAuth.AccessToken == "" {
		return c.JSON(http.StatusUnauthorized, models.HttpResponse{
			Code:    http.StatusUnauthorized,
			Message: "Access Denied. Unauthorized",
			Data:    "Access token not found",
		})
	}
	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "Login Success",
		Data:    adminAuth,
	})
}

// GetUserProfile godoc
// @Summary User GetProfile
// @Description User GetProfile
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/user/profile [get]
// @Security JwtToken
func (handler *userHandler) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()
	uid := c.Get("user_id").(int64)

	userData, err := handler.UserService.GetUserProfile(ctx, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    utils.ErrInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "Profile User",
		Data:    userData,
	})
}

// GetListUsers godoc
// @Summary User GetListUsers
// @Description User GetListUsers
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/users [get]
// @Security JwtToken / Admin
func (handler *userHandler) GetListUsers(c echo.Context) error {
	ctx := c.Request().Context()

	listUsers, err := handler.UserService.FindListUsers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	if len(listUsers) <= 0 {
		return c.JSON(http.StatusOK, models.HttpResponse{
			Code:    http.StatusOK,
			Message: "Success With Empty Data",
			Data:    listUsers,
		})
	}
	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "List Users Found",
		Data:    listUsers,
	})
}

// UpdateUser godoc
// @Summary User Update User
// @Description User Update User
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/user/updateprofile [get]
// @Security JwtToken / User
func (handler *userHandler) UpdateUserProfile(c echo.Context) error {
	uid := c.Get("user_id").(int64)
	ctx := c.Request().Context()
	req := models.UpdateProfileReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}
	if err := handler.UserService.UpdateUserProfile(ctx, req, uid); err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.HTTPResponseWithoutData{
		Code:    http.StatusOK,
		Message: "Profile Updated Successfully",
	})

}
