package factory

import (
	"lion-super-app/internal/database"
	"lion-super-app/internal/repository"

	"gorm.io/gorm"
)

type Factory struct {
	Db                                         *gorm.DB
	UserRepository                             repository.User
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("LION-SUPER-APP")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
}