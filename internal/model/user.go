package model

import (
	"encoding/base64"
	"lion-super-app/configs"
	"lion-super-app/internal/abstraction"
	"lion-super-app/pkg/util/date"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Username     string `json:"username" validate:"required" example:"administrator"`
	Name         string `json:"name" validate:"required" example:"Lelouch Lampergouee"`
	Password     string `json:"password" validate:"required" gorm:"-" example:"nevemor3"`
	Phone        int    `json:"phone" validate:"required" example:"1"`
	PasswordHash string `json:"-" gorm:"column:password"`
	Email        string `json:"email" validate:"required" example:"admin@lion.code"`
	IsActive     *bool  `json:"is_active" validate:"required" gorm:"default:true" example:"true"`
}

type UserFilter struct {
	Username          *string    `query:"username" filter:"ILIKE"`
	Name              *string    `query:"name" filter:"ILIKE"`
	Email             *string    `query:"email" filter:"ILIKE" example:"admin@lion.code"`
	RoleID            *int       `query:"role_id" example:"1"`
	IsActive          *bool      `query:"is_active"`
	CreatedAt         *time.Time `query:"created_at" filter:"DATE" example:"2022-08-17T15:04:05Z"`
	CreatedBy         *int       `query:"created_by" example:"1"`
	UserCreatedString *string    `query:"user_created" filter:"CUSTOM" example:"Lelouch Lampergouee"`
	Search            *string    `query:"search" filter:"CUSTOM" example:"Lelouch Lampergouee"`
}

type UserEntityModel struct {
	// abstraction
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt         time.Time        `json:"created_at"`
	UserCreated       *UserEntityModel `json:"-" gorm:"foreignKey:CreatedBy"`
	UserCreatedString string           `json:"user_created" gorm:"-"`
	CreatedBy         int              `json:"created_by"`

	// entity
	UserEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type UserFilterModel struct {
	UserFilter
}

func (UserEntityModel) TableName() string {
	return "users"
}

func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = 1

	m.hashPassword()
	m.Password = ""
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.Password != "" {
		m.hashPassword()
		m.Password = ""
	}
	return
}

func (m *UserEntityModel) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserEntityModel) GenerateToken() (string, error) {
	jwtKey := configs.Jwt().SecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  m.ID,
		"pid": m.Phone,
		"mid": m.Email,
		"exp": time.Now().Add(time.Minute * 50).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}

func Encode(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}
