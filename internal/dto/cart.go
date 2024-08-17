package dto

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"
	res "lion-super-app/pkg/util/response"
)

// Get
type CartGetRequest struct {
	abstraction.Pagination
	model.CartFilterModel
}
type CartGetResponse struct {
	Datas          []model.CartEntityModel
	PaginationInfo abstraction.PaginationInfo
}
type CartGetResponseDoc struct {
	Body struct {
		Meta res.Meta                    `json:"meta"`
		Data []model.CartEntityModel `json:"data"`
	} `json:"body"`
}

// GetByID
type CartGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CartGetByIDResponse struct {
	model.CartEntityModel
}
type CartGetByIDResponseDoc struct {
	Body struct {
		Meta res.Meta                `json:"meta"`
		Data CartGetByIDResponse `json:"data"`
	} `json:"body"`
}

// Create
type CartCreateRequest struct {
	model.CartEntity
	ProductID []int `json:"product_id" validate:"required"`
}
type CartCreateResponse struct {
	Datas []model.CartEntityModel
}
type CartCreateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CartCreateResponse `json:"data"`
	} `json:"body"`
}

// Update
type CartUpdateRequest struct {
	ID int `param:"id" validate:"required,numeric"`
	model.CartEntity
}
type CartUpdateResponse struct {
	model.CartEntityModel
}
type CartUpdateResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CartUpdateResponse `json:"data"`
	} `json:"body"`
}

// Delete
type CartDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}
type CartDeleteResponse struct {
	model.CartEntityModel
}
type CartDeleteResponseDoc struct {
	Body struct {
		Meta res.Meta               `json:"meta"`
		Data CartDeleteResponse `json:"data"`
	} `json:"body"`
}
