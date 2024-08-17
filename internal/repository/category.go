package repository

import (
	"fmt"
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"

	"gorm.io/gorm"
)

type Category interface {
	Find(ctx *abstraction.Context, m *model.CategoryFilterModel, p *abstraction.Pagination) (*[]model.CategoryEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.CategoryEntityModel, error)
	Create(ctx *abstraction.Context, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error)
	Get(ctx *abstraction.Context, m *model.CategoryFilterModel) (*[]model.CategoryEntityModel, error)
}

type category struct {
	abstraction.Repository
}

func NewCategory(db *gorm.DB) *category {
	return &category{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *category) Find(ctx *abstraction.Context, m *model.CategoryFilterModel, p *abstraction.Pagination) (*[]model.CategoryEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var datas []model.CategoryEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.CategoryEntityModel{})

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
	query = query.Order(sortBy)

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
	countQuery := conn.Model(&model.CategoryEntityModel{})
	countQuery = r.Filter(ctx, countQuery, *m)
	err := countQuery.Count(&totalData).Error
	if err != nil {
		return nil, nil, err
	}

	// Fetch data with pagination
	dataQuery := conn.Model(&model.CategoryEntityModel{})
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


func (r *category) FindByID(ctx *abstraction.Context, id *int) (*model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.CategoryEntityModel
	if err := conn.Where("id = ?", &id).First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
	// if err := conn.Where("id = ?", &id).Preload("Product").First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func (r *category) Create(ctx *abstraction.Context, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Create(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *category) Update(ctx *abstraction.Context, id *int, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id = ?", &id).Updates(e).Preload("Company").WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	if err := conn.Model(e).Where("id = ?", &id).First(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return e, nil

}

func (r *category) Delete(ctx *abstraction.Context, id *int, e *model.CategoryEntityModel) (*model.CategoryEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id =?", id).Update("status", 4).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	
	return e, nil
}

func (r *category) Get(ctx *abstraction.Context, m *model.CategoryFilterModel) (*[]model.CategoryEntityModel, error) {
	var datas []model.CategoryEntityModel

	conn := r.CheckTrx(ctx)
	query := conn.Model(&model.CategoryEntityModel{})
	
	query = r.Filter(ctx, query, *m)

	if err := query.Find(&datas).WithContext(ctx.Request().Context()).Error; err != nil {
		return &datas, err
	}

	return &datas, nil
}

func (r *category) GetCount(ctx *abstraction.Context, m *model.CategoryFilterModel) (*int64, error) {
	var jmlData int64
	conn := r.CheckTrx(ctx)
	query := conn.Model(&model.CategoryEntityModel{})
	//filter
	query = r.Filter(ctx, query, *m)

	if err := query.Count(&jmlData).WithContext(ctx.Request().Context()).Error; err != nil {
		return &jmlData, err
	}

	return &jmlData, nil
}
