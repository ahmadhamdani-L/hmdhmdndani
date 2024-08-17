package abstraction

import (
	"lion-super-app/pkg/util/date"
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt          time.Time  `json:"created_at"`
	CreatedBy          int        `json:"created_by"`
	UserCreatedString  string     `json:"user_created" gorm:"-"`
	ModifiedAt         *time.Time `json:"modified_at"`
	ModifiedBy         *int       `json:"modified_by"`
	UserModifiedString *string    `json:"user_modified" gorm:"-"`
}

type Filter struct {
	CreatedAt          *time.Time `query:"created_at" filter:"DATE" example:"2022-08-17T15:04:05Z"`
	CreatedBy          *int       `query:"created_by" example:"1"`
	UserCreatedString  *string    `query:"user_created" filter:"CUSTOM" example:"Lelouch Lampergouee"`
	ModifiedAt         *time.Time `query:"modified_at" filter:"DATE" example:"2022-08-17T15:04:05Z"`
	ModifiedBy         *int       `query:"modified_by" example:"1"`
	UserModifiedString *string    `query:"user_modified" filter:"CUSTOM" example:"Lelouch Lampergouee"`
}

type EntityImportedWorksheetDetail struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement;"`
	ModifiedAt *time.Time `json:"modified_at"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}
