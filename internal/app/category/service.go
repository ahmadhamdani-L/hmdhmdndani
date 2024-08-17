package category

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
	Repository        repository.Category
	// ProductRepository repository.Product
	Db                *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.CategoryGetRequest) (*dto.CategoryGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.CategoryGetByIDRequest) (*dto.CategoryGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CategoryCreateRequest) (*dto.CategoryCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.CategoryUpdateRequest) (*dto.CategoryUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.CategoryDeleteRequest) (*dto.CategoryDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.CategoryRepository
	// ProductRepository := f.ProductRepository

	db := f.Db
	return &service{
		Repository:        repository,
		// ProductRepository: ProductRepository,
		Db:                db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.CategoryGetRequest) (*dto.CategoryGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.CategoryFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.CategoryGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CategoryGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.CategoryGetByIDRequest) (*dto.CategoryGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.CategoryGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CategoryGetByIDResponse{
		CategoryEntityModel: *data,
	}
	return result, nil
}
func (s *service) Create(ctx *abstraction.Context, payload *dto.CategoryCreateRequest) (*dto.CategoryCreateResponse, error) {

	var data model.CategoryEntityModel
	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		fmt.Println(err)

		data.Context = ctx
		data.CategoryEntity = payload.CategoryEntity

		result, err := s.Repository.Create(ctx, &data)
		if err != nil {
			fmt.Println(err)
			return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
		}
		data = *result
		return nil

	}); err != nil {
		fmt.Println(err)
		return &dto.CategoryCreateResponse{}, err
	}
	result := &dto.CategoryCreateResponse{
		CategoryEntityModel: data,
	}
	return result, nil
}
func (s *service) Update(ctx *abstraction.Context, payload *dto.CategoryUpdateRequest) (*dto.CategoryUpdateResponse, error) {
	var data model.CategoryEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.CategoryEntity = payload.CategoryEntity
		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CategoryUpdateResponse{}, err
	}
	result := &dto.CategoryUpdateResponse{
		CategoryEntityModel: data,
	}
	return result, nil
}
func (s *service) Delete(ctx *abstraction.Context, payload *dto.CategoryDeleteRequest) (*dto.CategoryDeleteResponse, error) {
	var data model.CategoryEntityModel

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
		return &dto.CategoryDeleteResponse{}, err
	}
	result := &dto.CategoryDeleteResponse{
		CategoryEntityModel: data,
	}
	return result, nil
}
