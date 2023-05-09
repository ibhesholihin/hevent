package service

import (
	"context"
	"time"

	"github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/repository"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/ibhesholihin/hevent/utils/crypto"
	"github.com/ibhesholihin/hevent/utils/jwt"
)

type (

	//init service contract
	AdminService interface {
		AdminSignUp(c context.Context, adminReq models.AdminRegReq) (models.AdminRegRes, error)
		LoginAdmin(c context.Context, adminReq models.AdminLoginReq) (models.AdminLoginRes, error)

		GetAdminProfile(c context.Context, id int64) (models.Admin, error)
		UpdateAdmin(c context.Context, profileReq models.UpdateAdminReq, uid int64) error
	}

	//init service parameter
	adminService struct {
		stores         *repository.Stores
		cryptoSvc      crypto.CryptoService
		jwtSvc         jwt.JWTService
		contextTimeout time.Duration
	}
)

// admin sign up function
func (s *adminService) AdminSignUp(c context.Context, adminReq models.AdminRegReq) (models.AdminRegRes, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	/*
		user, err := s.userRepo.GetByEmail(ctx, request.Email)
		if err != nil && err != sql.ErrNoRows {
			return
		}
		if user.ID != 0 {
			err = utils.NewBadRequestError("email already registered")
			return
		}
	*/

	passwordHash, hashErr := s.cryptoSvc.CreatePasswordHash(ctx, adminReq.Password)
	if hashErr != nil {
		return models.AdminRegRes{}, hashErr
	}

	adm := models.Admin{
		Username: adminReq.Username,
		Email:    adminReq.Email,
		Password: passwordHash,
		Fullname: adminReq.Fullname,
	}
	admin, err := s.stores.Admin.CreateAdmin(adm)
	if err != nil {
		return models.AdminRegRes{}, err
	}

	adminRes := models.AdminRegRes{
		ID: admin.ID, Username: admin.Username,
	}

	return adminRes, nil
}

func (service adminService) LoginAdmin(c context.Context, adminReq models.AdminLoginReq) (models.AdminLoginRes, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()

	admin, err := service.stores.Admin.GetAdminByUsername(adminReq.Username)
	if err != nil || admin.ID < 1 {
		err = utils.NewBadRequestError("email and password not match")
		return models.AdminLoginRes{}, err
	}

	//err =  utils.CheckPasswordHash(adminReq.Password, admin.Password)
	//if err != nil {
	//	return md.AdminLoginRes{}, err
	//}

	if !service.cryptoSvc.ValidatePassword(ctx, admin.Password, adminReq.Password) {
		err = utils.NewBadRequestError("password not valid")
		return models.AdminLoginRes{}, err
	}

	accesstoken, refreshtoken, _ := service.jwtSvc.GenerateToken(ctx, admin.ID, true)

	adm := models.Admin{
		ID:            admin.ID,
		Access_Token:  accesstoken,
		Refresh_Token: refreshtoken,
		UpdatedAt:     time.Now(),
	}

	adminUpdate, err := service.stores.Admin.UpdateAdmin(adm)
	if err != nil {
		return models.AdminLoginRes{}, err
	}

	return models.AdminLoginRes{
		AccessToken: adminUpdate.Access_Token,
	}, nil
}

func (service adminService) GetAdminProfile(c context.Context, id int64) (models.Admin, error) {
	_, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()

	userData, err := service.stores.Admin.GetAdminById(id)
	if err != nil {
		return models.Admin{}, err
	}
	return userData, nil
}

// User update
func (s *adminService) UpdateAdmin(c context.Context, adminReq models.UpdateAdminReq, uid int64) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	adm, err := s.stores.Admin.GetAdminById(uid)
	if err != nil {
		return err
	}

	passwordHash, hashErr := s.cryptoSvc.CreatePasswordHash(ctx, adminReq.Password)
	if hashErr != nil {
		return hashErr
	}

	updatedProfile := models.Admin{
		ID:       adm.ID,
		Email:    adminReq.Email,
		Password: passwordHash,
	}
	_, err = s.stores.Admin.UpdateAdmin(updatedProfile)
	if err != nil {
		return err
	}
	return nil
}
