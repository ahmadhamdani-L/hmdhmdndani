package model

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/pkg/util/date"

	"gorm.io/gorm"
)

type CartEntity struct {
	UserID      int    `gorm:"type:integer" json:"user_id"`
	NamaUser    string `gorm:"type:varchar(100);not null"`
	ProductID   int    `gorm:"type:integer" json:"category_id"`
	Category    string `gorm:"type:string" json:"category"`
	NameProduct string `gorm:"type:string" json:"name_product"`
	Price       int    `gorm:"type:integer" json:"price"`
	Description string `gorm:"type:text" json:"description"`
}

type CartFilter struct {
	NameUser    *string `query:"name_user" filter:"ILIKE"`
	NameProduct *string `query:"name_product"`
	UserID      *int    `query:"user_id"`
	ProductID   *int    `query:"product_id"`
}

type CartEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	CartEntity

	Context *abstraction.Context `json:"-" gorm:"-"`
}

type CartFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	CartFilter
}

func (CartEntityModel) TableName() string {
	return "Cart"
}

func (m *CartEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.ID
	return
}

func (m *CartEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.ID
	return
}
