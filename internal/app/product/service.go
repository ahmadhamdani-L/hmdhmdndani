package product

import (
	"fmt"
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/dto"
	"lion-super-app/internal/factory"
	"lion-super-app/internal/model"
	"lion-super-app/internal/repository"
	"lion-super-app/pkg/util/response"
	"lion-super-app/pkg/util/trxmanager"
	
	"gorm.io/gorm"
)

type service struct {
	Repository        repository.Product
	// ProductRepository repository.Product
	Db                *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.ProductGetRequest) (*dto.ProductGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.ProductGetByIDRequest) (*dto.ProductGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.ProductCreateRequest) (*dto.ProductCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.ProductUpdateRequest) (*dto.ProductUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.ProductDeleteRequest) (*dto.ProductDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.ProductRepository

	db := f.Db
	return &service{
		Repository:        repository,
		Db:                db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.ProductGetRequest) (*dto.ProductGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.ProductFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.ProductGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.ProductGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.ProductGetByIDRequest) (*dto.ProductGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.ProductGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.ProductGetByIDResponse{
		ProductEntityModel: *data,
	}
	return result, nil
}
func (s *service) Create(ctx *abstraction.Context, payload *dto.ProductCreateRequest) (*dto.ProductCreateResponse, error) {

	var data model.ProductEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		fmt.Println(err)

		data.Context = ctx
		data.ProductEntity = payload.ProductEntity

		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			fmt.Println(err)
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil

	}); err != nil {
		fmt.Println(err)
		return &dto.ProductCreateResponse{}, err
	}
	result := &dto.ProductCreateResponse{
		ProductEntityModel: data,
	}
	return result, nil
}
func (s *service) Update(ctx *abstraction.Context, payload *dto.ProductUpdateRequest) (*dto.ProductUpdateResponse, error) {
	var data model.ProductEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.ProductEntity = payload.ProductEntity
		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.ProductUpdateResponse{}, err
	}
	result := &dto.ProductUpdateResponse{
		ProductEntityModel: data,
	}
	return result, nil
}
func (s *service) Delete(ctx *abstraction.Context, payload *dto.ProductDeleteRequest) (*dto.ProductDeleteResponse, error) {
	var data model.ProductEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		findById, err := s.Repository.FindByID(ctx, &payload.ID)
		if err != nil {
			return err
		}

		if findById == nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}

		data.Context = ctx
		result, err := s.Repository.Delete(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.ProductDeleteResponse{}, err
	}
	result := &dto.ProductDeleteResponse{
		ProductEntityModel: data,
	}
	return result, nil
}
