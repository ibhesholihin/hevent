package handler

import (
	"net/http"

	"github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/service"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/labstack/echo/v4"
)

type (
	AdminHandler interface {
		//GetAdmin(c echo.Context) error
		AdminSignUp(c echo.Context) error
		GetAdminProfile(c echo.Context) error
		LoginAdmin(c echo.Context) error
		UpdateAdmin(c echo.Context) error
		//DeleteAdmin(c echo.Context) error
	}

	adminHandler struct {
		service.AdminService
	}
)

// Signup Admin
// @Summary Get Sign up for admin.
// @Description Sign up admin process.
// @Tags Admin
// @Accept */*
// @Produce json
// @Success 200 {object} models.Admin
// @Failure 500 {object} utils.Error
// @Router /api/v1/auth/admin/signup [post]
func (handler *adminHandler) AdminSignUp(c echo.Context) error {
	//var req *models.AdminRegReq
	ctx := c.Request().Context()
	req := models.AdminRegReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    utils.ErrBadQueryParams,
		})
	}

	createdAdmin, err := handler.AdminService.AdminSignUp(ctx, req)
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
		Data:    createdAdmin,
	})
}

// AdminSignIn godoc
// @Summary Admin SignIn
// @Description Admin SignIn
// @Tags Admin
// @Accept json
// @Produce json
// @Param signin body request.AdminLoginReq true "SignIn admin"
// @Success 200
// @Router /api/v1/auth/admin/signin [post]
func (handler *adminHandler) LoginAdmin(c echo.Context) error {
	ctx := c.Request().Context()
	req := models.AdminLoginReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Admin Failed",
			Data:    utils.ErrBadRequest,
		})
	}

	adminAuth, err := handler.AdminService.LoginAdmin(ctx, req)
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

// GetAdminProfile godoc
// @Summary Admin GetProfile
// @Description Admin GetProfile
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/admin/profile [get]
// @Security JwtToken
func (handler *adminHandler) GetAdminProfile(c echo.Context) error {
	ctx := c.Request().Context()
	uid := c.Get("user_id").(int64)
	userData, err := handler.AdminService.GetAdminProfile(ctx, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    utils.ErrInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "Profile Admin",
		Data:    userData,
	})
}

// UpdateAdmin godoc
// @Summary Admin Update Admin
// @Description Admin Update Admin
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/admin/updateprofile [get]
// @Security JwtToken
func (handler *adminHandler) UpdateAdmin(c echo.Context) error {
	uid := c.Get("user_id").(int64)
	ctx := c.Request().Context()

	req := models.UpdateAdminReq{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}
	if err := handler.AdminService.UpdateAdmin(ctx, req, uid); err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.HTTPResponseWithoutData{
		Code:    http.StatusOK,
		Message: "Admin Profile Updated Successfully",
	})
}
