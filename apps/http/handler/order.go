package handler

import (
	"github.com/ibhesholihin/hevent/apps/service"
	"github.com/labstack/echo/v4"
)

type (
	OrderHandler interface {
		GetUserOrder(c echo.Context) error
		AddUserOrder(c echo.Context) error
	}

	orderHandler struct {
		service.OrderService
	}
)

func (h *orderHandler) GetUserOrder(c echo.Context) error {
	return nil
}

func (h *orderHandler) AddUserOrder(c echo.Context) error {
	return nil
}
