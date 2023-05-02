package service

import (
	"context"
	"time"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/repository"
	"github.com/ibhesholihin/hevent/utils"
	"github.com/ibhesholihin/hevent/utils/crypto"
	"github.com/ibhesholihin/hevent/utils/jwt"
)

type (
	//init service contract
	UserService interface {
		FindListUsers(c context.Context) ([]md.User, error)
		GetUserProfile(c context.Context, id int64) (md.User, error)
		UserSignUp(c context.Context, userReq md.UserRegRequest) (md.UserRegResponse, error)
		LoginUser(c context.Context, userReq md.UserLoginRequest) (md.UserLoginResponse, error)

		UpdateUserProfile(c context.Context, profileReq md.UpdateProfileReq, uid int64) error
	}

	//init service parameter
	userService struct {
		stores         *repository.Stores
		cryptoSvc      crypto.CryptoService
		jwtSvc         jwt.JWTService
		contextTimeout time.Duration
	}
)

// user sign up function
func (s *userService) UserSignUp(c context.Context, userReq md.UserRegRequest) (md.UserRegResponse, error) {
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

	passwordHash, hashErr := s.cryptoSvc.CreatePasswordHash(ctx, userReq.Password)
	if hashErr != nil {
		return md.UserRegResponse{}, hashErr
	}

	usr := md.User{
		Username: userReq.Username,
		Password: passwordHash,
		UserProfile: md.UserProfile{
			Fullname: userReq.Fullname,
			Email:    userReq.Email,
		},
	}
	user, err := s.stores.User.CreateUser(usr)
	if err != nil {
		return md.UserRegResponse{}, err
	}
	userRes := md.UserRegResponse{
		UID:      user.ID,
		Username: user.Username,
	}

	return userRes, nil
}

// user login
func (service *userService) LoginUser(c context.Context, userReq md.UserLoginRequest) (md.UserLoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()

	user, err := service.stores.User.GetUserByUsername(userReq.Username)
	if err != nil || user.ID < 1 {
		return md.UserLoginResponse{}, err
	}

	if !service.cryptoSvc.ValidatePassword(ctx, user.Password, user.Password) {
		err = utils.NewBadRequestError("password not valid")
		return md.UserLoginResponse{}, err
	}

	//generate token user is (admin:false)
	accesstoken, refreshtoken, _ := service.jwtSvc.GenerateToken(ctx, user.ID, false)

	adm := md.User{
		ID:            user.ID,
		Access_Token:  accesstoken,
		Refresh_Token: refreshtoken,
	}

	adminUpdate, err := service.stores.User.UpdateUser(adm)
	if err != nil {
		return md.UserLoginResponse{}, err
	}

	return md.UserLoginResponse{
		AccessToken: adminUpdate.Access_Token,
	}, nil
}

// Users Get List
func (s *userService) FindListUsers(c context.Context) ([]md.User, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listUsers, err := s.stores.User.GetUserList()
	if err != nil {
		return []md.User{}, err
	}
	return listUsers, nil
}

// Users Get List
func (s *userService) GetUserProfile(c context.Context, id int64) (md.User, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	userData, err := s.stores.User.FindUserProfileByUID(id)
	if err != nil {
		return md.User{}, err
	}
	return userData, nil
}

// User update
func (s *userService) UpdateUserProfile(c context.Context, profileReq md.UpdateProfileReq, uid int64) error {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	profileID, err := s.stores.User.FindUserProfileID(uid)
	if err != nil {
		return err
	}
	updatedProfile := md.UserProfile{
		ID:       profileID,
		Fullname: profileReq.Fullname,
		Email:    profileReq.Email,
		Gender:   profileReq.Gender,
		Phone:    profileReq.Phone,
	}
	_, err = s.stores.User.UpdateProfile(updatedProfile)
	if err != nil {
		return err
	}
	return nil
}
