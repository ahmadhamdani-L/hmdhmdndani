package model

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/pkg/util/date"

	"gorm.io/gorm"
)

type ProductEntity struct {
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	CategoryID  int    `gorm:"type:integer" json:"category_id"`
	Price       int    `gorm:"type:integer" json:"price"`
}

type ProductFilter struct {
	Name *string `query:"name" filter:"ILIKE"`
}

type ProductEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	ProductEntity

	Context *abstraction.Context `json:"-" gorm:"-"`
}

type ProductFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	ProductFilter
}

func (ProductEntityModel) TableName() string {
	return "Product"
}

func (m *ProductEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.ID
	return
}

func (m *ProductEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.ID
	return
}
