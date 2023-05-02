package service

import (
	"time"

	"github.com/ibhesholihin/hevent/apps/repository"
	"github.com/ibhesholihin/hevent/utils/crypto"
	"github.com/ibhesholihin/hevent/utils/jwt"
)

type Services struct {
	Admin AdminService
	User  UserService
	Event EventService
	Order OrderService
}

func NewService(s *repository.Stores, cryptoSvc crypto.CryptoService, jwtSvc jwt.JWTService, contextTimeout time.Duration) *Services {
	return &Services{
		Admin: &adminService{s, cryptoSvc, jwtSvc, contextTimeout},
		User:  &userService{s, cryptoSvc, jwtSvc, contextTimeout},
		Event: &eventService{s, contextTimeout},
		Order: &orderService{s, contextTimeout},
	}
}
