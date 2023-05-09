package handler

import (
	"net/http"

	"github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/service"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/labstack/echo/v4"
)

type (
	OrderHandler interface {
		GetCartSession(c echo.Context) error
		AddItemToCart(c echo.Context) error
		GetItemsCart(c echo.Context) error

		DeleteItemFromCart(c echo.Context) error
		CreateOrder(c echo.Context) error
		GetListOrders(c echo.Context) error
		GetOrderById(c echo.Context) error
		UploadReceipt(c echo.Context) error

		//test payment handler
		TestPayment(c echo.Context) error
	}

	orderHandler struct {
		service.OrderService
	}
)

func (h *orderHandler) GetCartSession(c echo.Context) error {
	//ctx := c.Request().Context()
	//uid := c.Get("user_id").(int64)

	return nil
}

func (h *orderHandler) AddItemToCart(c echo.Context) error {
	//ctx := c.Request().Context()
	//h.PayService.GeneratePayReq(ctx,"ibhe","ibhe@admin.com","123",30000)

	return nil
}

func (h *orderHandler) GetItemsCart(c echo.Context) error {
	return nil
}

func (h *orderHandler) DeleteItemFromCart(c echo.Context) error {
	return nil
}

func (h *orderHandler) CreateOrder(c echo.Context) error {
	return nil
}

func (h *orderHandler) GetListOrders(c echo.Context) error {
	return nil
}

func (h *orderHandler) GetOrderById(c echo.Context) error {
	return nil
}

func (h *orderHandler) UploadReceipt(c echo.Context) error {
	ctx := c.Request().Context()

	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    ctx,
	})
}

func (h *orderHandler) TestPayment(c echo.Context) error {
	ctx := c.Request().Context()

	type payment struct {
		Tokenpay string `json:"token"`
		Urlpay   string `json:"url"`
	}

	payID, payUrl, err := h.OrderService.TestPayment(ctx)

	pay := &payment{
		Tokenpay: payID,
		Urlpay:   payUrl,
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    utils.ErrInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, models.HttpResponse{
		Code:    http.StatusOK,
		Message: "Order Success",
		Data:    pay,
	})
}
