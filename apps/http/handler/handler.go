package handler

import "github.com/ibhesholihin/hevent/apps/service"

type Handlers struct {
	AdminHandler
	UserHandler
	EventHandler
	OrderHandler
}

func NewHandler(s *service.Services) *Handlers {
	return &Handlers{
		AdminHandler: &adminHandler{s.Admin},
		UserHandler:  &userHandler{s.User},
		EventHandler: &eventHandler{s.Event},
		OrderHandler: &orderHandler{s.Order},
	}
}
