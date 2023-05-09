package handler

import (
	"net/http"
	"strconv"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/service"
	"github.com/labstack/echo/v4"
)

type (
	EventHandler interface {
		GetListCategory(c echo.Context) error
		GetCategory(c echo.Context) error
		AddCategory(c echo.Context) error
		UpdateEventCategory(c echo.Context) error
		RemoveEventCategory(c echo.Context) error

		GetListEvents(c echo.Context) error
		GetEvent(c echo.Context) error
		AddEvent(c echo.Context) error
		UpdateEvent(c echo.Context) error
		RemoveEvent(c echo.Context) error

		GetEventPrice(c echo.Context) error
		GetEventPriceTipe(c echo.Context) error
		AddEventPrice(c echo.Context) error
		/*
			UpdateEventPrice(c echo.Context) error
			RemoveEventPrice(c echo.Context) error
		*/
	}

	eventHandler struct {
		service.EventService
	}
)

//////////////////
//EVENTS HANDLER//
//////////////////

// GetEvent godoc
// @Summary Event GetEvent
// @Description Event GetEvent
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/public/events/:eventid [get]
func (h *eventHandler) GetEvent(c echo.Context) error {
	eventid, _ := strconv.Atoi(c.Param("eventid"))
	ctx := c.Request().Context()
	event, err := h.EventService.GetEvent(ctx, uint(eventid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}

	if event.ID == 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code:    http.StatusOK,
			Message: "Success With Empty Data",
			Data:    event,
		})
	}

	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "Event Found",
		Data:    event,
	})
}

// GetListEvents godoc
// @Summary Event GetListEvents
// @Description Event GetListEvents
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/events [get]
func (h *eventHandler) GetListEvents(c echo.Context) error {
	ctx := c.Request().Context()

	listEvents, err := h.EventService.FindListEvents(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}

	if len(listEvents) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code:    http.StatusOK,
			Message: "Success With Empty Data",
			Data:    listEvents,
		})
	}

	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "List Events Found",
		Data:    listEvents,
	})
}

// AddEvent godoc
// @Summary Event AddEvent
// @Description Event AddEvent
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/events [post]
// @Security JwtToken / Admin
func (h *eventHandler) AddEvent(c echo.Context) error {
	ctx := c.Request().Context()
	evt := md.CreateEventReq{}
	if err := c.Bind(&evt); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}
	eventRes, err := h.EventService.CreateEvent(ctx, evt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code:    http.StatusCreated,
		Message: "Event Created Successfully",
		Data:    eventRes,
	})
}

// UpdateEvent godoc
// @Summary Event Update Event
// @Description Event Update Event
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/events/:eventid [put]
// @Security JwtToken / Admin
func (h *eventHandler) UpdateEvent(c echo.Context) error {
	//uid, _ := strconv.Atoi(c.Param("eventid"))

	return nil
}

// RemoveEvent godoc
// @Summary Event Remove Event
// @Description Event Remove Event
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/events/:eventid [delete]
// @Security JwtToken / Admin
func (h *eventHandler) RemoveEvent(c echo.Context) error {
	//uid, _ := strconv.Atoi(c.Param("eventid"))
	return nil
}

//////////////////////////
//EVENT CATEGORY HANDLER//
//////////////////////////

// GetListCategory godoc
// @Summary Event GetListCategory
// @Description Event GetListCategory
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/public/category [get]
func (h *eventHandler) GetListCategory(c echo.Context) error {
	ctx := c.Request().Context()
	listCategory, err := h.EventService.FindListCategory(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	if len(listCategory) <= 0 {
		return c.JSON(http.StatusOK, md.HttpResponse{
			Code:    http.StatusOK,
			Message: "Success With Empty Data",
			Data:    listCategory,
		})
	}
	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "List Categories Found",
		Data:    listCategory,
	})
}

// GetCategory godoc
// @Summary Event Get Category
// @Description Event Get Detail Category
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/public/category/:categoryid [get]
func (h *eventHandler) GetCategory(c echo.Context) error {
	ctx := c.Request().Context()
	uid, _ := strconv.Atoi(c.Param("categoryid"))

	category, err := h.EventService.GetEventCategory(ctx, uint(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "Category Found",
		Data:    category,
	})
}

// AddCategory godoc
// @Summary Event AddCategory
// @Description Event AddCategory
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/category [post]
// @Security JwtToken / Admin
func (h *eventHandler) AddCategory(c echo.Context) error {
	ctx := c.Request().Context()
	ctv := md.CreateEventCategoryReq{}

	if err := c.Bind(&ctv); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}

	catRes, err := h.EventService.AddCategory(ctx, ctv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code:    http.StatusCreated,
		Message: "Product Created Successfully",
		Data:    catRes,
	})
}

// UpdateEventCategory godoc
// @Summary Event Update Event Category
// @Description Event Update Event Category
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/category/:categoryid [put]
// @Security JwtToken / Admin
func (h *eventHandler) UpdateEventCategory(c echo.Context) error {
	uid, _ := strconv.Atoi(c.Param("categoryid"))
	ctx := c.Request().Context()
	req := md.UpdateEventReq{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}

	if err := h.EventService.UpdateEvent(ctx, req, uint(uid)); err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, md.HTTPResponseWithoutData{
		Code:    http.StatusOK,
		Message: "Event Updated Successfully",
	})
}

// RemoveEventCategory godoc
// @Summary Event Remove Event Category
// @Description Event Remove Event Category
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/category/:categoryid [delete]
// @Security JwtToken / Admin
func (h *eventHandler) RemoveEventCategory(c echo.Context) error {
	//ctx := c.Request().Context()
	//uid, _ := strconv.Atoi(c.Param("categoryid"))
	return nil
}

////////////////////////////
//EVENT PRICE TYPE HANDLER//
////////////////////////////

// GetEventPrice godoc
// @Summary Event GetEventPrice
// @Description Event GetEventPrice
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/public/events/price/:eventid [get]
func (h *eventHandler) GetEventPrice(c echo.Context) error {
	//ctx := c.Request().Context()
	uid, _ := strconv.Atoi(c.Param("eventid"))

	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "Event Price Found",
		Data:    uid,
	})
}

// GetEventPriceTipe godoc
// @Summary Event GetEventPriceTipe
// @Description Event GetEventPriceTipe
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/public/events/price/:tipeid [get]
func (h *eventHandler) GetEventPriceTipe(c echo.Context) error {
	ctx := c.Request().Context()
	uid, _ := strconv.Atoi(c.Param("tipeid"))

	price, err := h.EventService.FindListPrice(ctx, uint(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, md.HttpResponse{
		Code:    http.StatusOK,
		Message: "Event Price Found",
		Data:    price,
	})
}

// AddEventPrice godoc
// @Summary Event AddEvent
// @Description Event AddEvent
// @Tags Event
// @Accept json
// @Produce json
// @Success 200
// @Router api/v1/admin/events/price/:eventid [post]
// @Security JwtToken / Admin
func (h *eventHandler) AddEventPrice(c echo.Context) error {
	ctx := c.Request().Context()
	uid, _ := strconv.Atoi(c.Param("eventid"))

	evt := md.EventPriceTipe{}
	evt.EventID = uint(uid)
	if err := c.Bind(&evt); err != nil {
		return c.JSON(http.StatusBadRequest, md.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Body Request",
			Data:    err.Error(),
		})
	}
	Res, err := h.EventService.AddEventPrice(ctx, evt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, md.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, md.HttpResponse{
		Code:    http.StatusCreated,
		Message: "Event Price Created Successfully",
		Data:    Res,
	})
}
