package model

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/pkg/util/date"

	"gorm.io/gorm"
)

type CategoryEntity struct {
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Products    []ProductEntityModel `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

type CategoryFilter struct {
	Name *string `query:"name" filter:"ILIKE"`
}

type CategoryEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	CategoryEntity

	// relations
	Product []ProductEntityModel `json:"category_id" gorm:"foreignKey:CategoryID"`
	
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type CategoryFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	CategoryFilter
}

func (CategoryEntityModel) TableName() string {
	return "Category"
}

func (m *CategoryEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.ID
	return
}

func (m *CategoryEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.ID
	return
}
