package cart

import (
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
	Repository        repository.Cart
	ProductRepository repository.Product
	CategoryRepository repository.Category
	Db                *gorm.DB
}

type Service interface {
	Find(ctx *abstraction.Context, payload *dto.CartGetRequest) (*dto.CartGetResponse, error)
	FindByID(ctx *abstraction.Context, payload *dto.CartGetByIDRequest) (*dto.CartGetByIDResponse, error)
	Create(ctx *abstraction.Context, payload *dto.CartCreateRequest) (*dto.CartCreateResponse, error)
	Update(ctx *abstraction.Context, payload *dto.CartUpdateRequest) (*dto.CartUpdateResponse, error)
	Delete(ctx *abstraction.Context, payload *dto.CartDeleteRequest) (*dto.CartDeleteResponse, error)
}

func NewService(f *factory.Factory) *service {
	repository := f.CartRepository
	productRepository := f.ProductRepository
	categoryRepository := f.CategoryRepository

	db := f.Db
	return &service{
		Repository:        repository,
		ProductRepository: productRepository,
		CategoryRepository: categoryRepository,
		Db:                db,
	}
}

func (s *service) Find(ctx *abstraction.Context, payload *dto.CartGetRequest) (*dto.CartGetResponse, error) {
	data, info, err := s.Repository.Find(ctx, &payload.CartFilterModel, &payload.Pagination)
	if err != nil {
		return &dto.CartGetResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CartGetResponse{
		Datas:          *data,
		PaginationInfo: *info,
	}
	return result, nil
}

func (s *service) FindByID(ctx *abstraction.Context, payload *dto.CartGetByIDRequest) (*dto.CartGetByIDResponse, error) {
	data, err := s.Repository.FindByID(ctx, &payload.ID)
	if err != nil {
		return &dto.CartGetByIDResponse{}, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	result := &dto.CartGetByIDResponse{
		CartEntityModel: *data,
	}
	return result, nil
}
func (s *service) Create(ctx *abstraction.Context, payload *dto.CartCreateRequest) (*dto.CartCreateResponse, error) {
    var data []model.CartEntityModel

    if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
        for _, productID := range payload.ProductID {
            product, err := s.ProductRepository.FindByID(ctx, &productID)
            if err != nil {
                return err
            }

            category, err := s.CategoryRepository.FindByID(ctx, &product.CategoryID)
            if err != nil {
                return err
            }

            // Create a new cart entry for each product
            cart := model.CartEntityModel{
                Context:      ctx,
                CartEntity: model.CartEntity{
					UserID: ctx.Auth.ID,
                    Category:    category.Name,
                    NamaUser:    ctx.Auth.Name,
                    NameProduct: product.Name,
                    Price:       product.Price,
                    ProductID:   product.ID,
                },
            }

            result, err := s.Repository.Create(ctx, &cart)
            if err != nil {
                return response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
            }

            data = append(data, *result)
        }
        return nil
    }); err != nil {
        return &dto.CartCreateResponse{}, err
    }

    result := &dto.CartCreateResponse{
        Datas: data,
    }
    return result, nil
}

func (s *service) Update(ctx *abstraction.Context, payload *dto.CartUpdateRequest) (*dto.CartUpdateResponse, error) {
	var data model.CartEntityModel

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		if _, err := s.Repository.FindByID(ctx, &payload.ID); err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err)
		}
		data.Context = ctx
		data.CartEntity = payload.CartEntity
		result, err := s.Repository.Update(ctx, &payload.ID, &data)
		if err != nil {
			return response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
		data = *result
		return nil
	}); err != nil {
		return &dto.CartUpdateResponse{}, err
	}
	result := &dto.CartUpdateResponse{
		CartEntityModel: data,
	}
	return result, nil
}
func (s *service) Delete(ctx *abstraction.Context, payload *dto.CartDeleteRequest) (*dto.CartDeleteResponse, error) {
	var data model.CartEntityModel

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
		return &dto.CartDeleteResponse{}, err
	}
	result := &dto.CartDeleteResponse{
		CartEntityModel: data,
	}
	return result, nil
}
