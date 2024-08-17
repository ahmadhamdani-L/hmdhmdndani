package auth

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/dto"
	"lion-super-app/internal/factory"
	"lion-super-app/internal/model"
	"lion-super-app/internal/repository"
	res "lion-super-app/pkg/util/response"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error)
}

type service struct {
	Repository repository.User
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	var result *dto.AuthLoginResponse

	data, err := s.Repository.FindByUsername(ctx, &payload.Username)
	if data == nil {
		return nil, res.CustomErrorBuilder(http.StatusBadRequest, res.E_BAD_REQUEST, "Username is incorrect")
	}

	if data.IsActive == nil || (data.IsActive != nil && !*data.IsActive) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password))
	if err != nil {
		return nil, res.CustomErrorBuilder(http.StatusBadRequest, res.E_BAD_REQUEST, "Password is incorrect")
	}

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthLoginResponse{
		Token:           token,
		UserEntityModel: *data,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.AuthRegisterRequest) (*dto.AuthRegisterResponse, error) {
	var result *dto.AuthRegisterResponse

	existingUser, err := s.Repository.FindByUsername(ctx, &payload.Username)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if existingUser != nil {
		return nil, res.CustomErrorBuilder(http.StatusBadRequest, res.E_BAD_REQUEST, "Username is already taken")
	}

	active := true

	newUser := &model.UserEntityModel{
		UserEntity: model.UserEntity{
			Username: payload.Username,
			Name:     payload.Name,
			Email:    payload.Email,
			Phone:    payload.Phone,
			Password: payload.Password,
			IsActive: &active,
		},
	}

	_, err = s.Repository.Create(ctx, newUser)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.AuthRegisterResponse{
		UserEntityModel: *newUser,
	}

	return result, nil
}
