package dto

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"
	res "lion-super-app/pkg/util/response"
)

// Get
type ProductGetRequest struct {
	abstraction.Pagination
	model.ProductFilterModel
}
type ProductGetResponse struct {
	Datas          []model.ProductEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type ProductGetResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data []model.ProductEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type ProductGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ProductGetByIDResponse struct {
	model.ProductEntityModel
}
type ProductGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data ProductGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type ProductCreateRequest struct {
	model.ProductEntity
}
type ProductCreateResponse struct {
	model.ProductEntityModel
}
type ProductCreateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data ProductCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type ProductUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.ProductEntity
}
type ProductUpdateResponse struct {
	model.ProductEntityModel
}
type ProductUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data ProductUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type ProductDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type ProductDeleteResponse struct {
	model.ProductEntityModel
}
type ProductDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data ProductDeleteResponse `json:"data"`
	} `json:"body"`
}
