package user

import (
	"lion-super-app/internal/factory"
	"lion-super-app/internal/repository"
	"gorm.io/gorm"
)

type service struct {
	Repository                  repository.User
	Db                          *gorm.DB
}

type Service interface {
}

func NewService(f *factory.Factory) *service {
	repository := f.UserRepository
	db := f.Db
	return &service{
		Repository:                  repository,
		Db:                          db,
	}
}

