package repository

import (
	"fmt"
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"

	"gorm.io/gorm"
)

type Cart interface {
	Find(ctx *abstraction.Context, m *model.CartFilterModel, p *abstraction.Pagination) (*[]model.CartEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.CartEntityModel, error)
	Create(ctx *abstraction.Context, e *model.CartEntityModel) (*model.CartEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.CartEntityModel) (*model.CartEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.CartEntityModel) (*model.CartEntityModel, error)
	Get(ctx *abstraction.Context, m *model.CartFilterModel) (*[]model.CartEntityModel, error)
}

type cart struct {
	abstraction.Repository
}

func NewCart(db *gorm.DB) *cart {
	return &cart{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *cart) Find(ctx *abstraction.Context, m *model.CartFilterModel, p *abstraction.Pagination) (*[]model.CartEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.CartEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.CartEntityModel{})

	// Filter
	query = r.Filter(ctx, query, *m)

	// Sorting
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}
	if p.SortBy == nil {
		sortBy := "created_at"
		p.SortBy = &sortBy
	}
	sortBy := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)

	// Pagination
	if p.Page == nil {
		page := 1
		p.Page = &page
	}
	if p.PageSize == nil {
		pageSize := 10
		p.PageSize = &pageSize
	}
	limit := *p.PageSize
	offset := limit * (*p.Page - 1)

	var totalData int64
	countQuery := conn.Model(&model.CartEntityModel{})
	countQuery = r.Filter(ctx, countQuery, *m)
	err := countQuery.Count(&totalData).Error
	if err != nil {
		return nil, nil, err
	}

	// Fetch data with pagination
	dataQuery := conn.Model(&model.CartEntityModel{})
	dataQuery = r.Filter(ctx, dataQuery, *m)
	dataQuery = dataQuery.Order(sortBy).Limit(limit).Offset(offset)
	err = dataQuery.Find(&datas).Error
	if err != nil {
		return nil, nil, err
	}

	info = abstraction.PaginationInfo{
		Pagination: p,
		Count:      int(totalData),
		MoreRecords: totalData > int64(offset+limit),
	}

	return &datas, &info, nil
}


func (r *cart) FindByID(ctx *abstraction.Context, id *int) (*model.CartEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.CartEntityModel
	if err := conn.Where("id = ?", &id).First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
	// if err := conn.Where("id = ?", &id).Preload("Cart").First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *cart) Create(ctx *abstraction.Context, e *model.CartEntityModel) (*model.CartEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *cart) Update(ctx *abstraction.Context, id *int, e *model.CartEntityModel) (*model.CartEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).Preload("Company").WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Where("id = ?", &id).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil

}

func (r *cart) Delete(ctx *abstraction.Context, id *int, e *model.CartEntityModel) (*model.CartEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id =?", id).Update("status", 4).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	
	return e, nil
}

func (r *cart) Get(ctx *abstraction.Context, m *model.CartFilterModel) (*[]model.CartEntityModel, error) {
	var datas []model.CartEntityModel

	conn := r.CheckTrx(ctx)
	query := conn.Model(&model.CartEntityModel{})
	
	query = r.Filter(ctx, query, *m)

	if err := query.Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
		return &datas, err
	}

	return &datas, nil
}

func (r *cart) GetCount(ctx *abstraction.Context, m *model.CartFilterModel) (*int64, error) {
	var jmlData int64
	conn := r.CheckTrx(ctx)
	query := conn.Model(&model.CartEntityModel{})
	//filter
	query = r.Filter(ctx, query, *m)

	if err := query.Count(&jmlData).WithContext(ctx.Request().Context()).Error; err != nil {
		return &jmlData, err
	}

	return &jmlData, nil
}
