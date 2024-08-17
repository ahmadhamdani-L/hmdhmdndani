package dto

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"
	res "lion-super-app/pkg/util/response"
)

// Get
type CategoryGetRequest struct {
	abstraction.Pagination
	model.CategoryFilterModel
}
type CategoryGetResponse struct {
	Datas          []model.CategoryEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type CategoryGetResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data []model.CategoryEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type CategoryGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CategoryGetByIDResponse struct {
	model.CategoryEntityModel
}
type CategoryGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data CategoryGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type CategoryCreateRequest struct {
	model.CategoryEntity
	Product []model.ProductEntity
}
type CategoryCreateResponse struct {
	model.CategoryEntityModel
}
type CategoryCreateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CategoryCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type CategoryUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.CategoryEntity
}
type CategoryUpdateResponse struct {
	model.CategoryEntityModel
}
type CategoryUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CategoryUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type CategoryDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CategoryDeleteResponse struct {
	model.CategoryEntityModel
}
type CategoryDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CategoryDeleteResponse `json:"data"`
	} `json:"body"`
}
