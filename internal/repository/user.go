package repository

import (
	"lion-super-app/internal/abstraction"
	"lion-super-app/internal/model"

	"gorm.io/gorm"
)

type User interface {
	checkTrx(ctx *abstraction.Context) *gorm.DB
	FindByUsername(ctx *abstraction.Context, username *string) (*model.UserEntityModel, error)
	Create(ctx *abstraction.Context, data *model.UserEntityModel) (*model.UserEntityModel, error)
	Delete(ctx *abstraction.Context, id *int, e *model.UserEntityModel) (*model.UserEntityModel, error)
	FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error)
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *user) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}

func (r *user) FindByUsername(ctx *abstraction.Context, username *string) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.UserEntityModel
	err := conn.Where("username = ? OR email = ?", username, username).Find(&data).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx *abstraction.Context, e *model.UserEntityModel) (*model.UserEntityModel, error) {
	conn := r.checkTrx(ctx)

	err := conn.Create(&e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&e).First(&e).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *user) Delete(ctx *abstraction.Context, id *int, e *model.UserEntityModel) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	if err := conn.Model(e).Where("id =?", id).Delete(e).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *user) FindByEmail(ctx *abstraction.Context, email *string) (*model.UserEntityModel, error) {
	conn := r.CheckTrx(ctx)

	var data model.UserEntityModel
	if err := conn.Where("email = ?", &email).First(&data).WithContext(ctx.Request().Context()).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
