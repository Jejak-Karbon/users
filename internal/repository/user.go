package repository

import (
	"context"
	_ "strings"

	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByID(ctx context.Context, id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
	Create(ctx context.Context, data model.User) error
	Update(ctx context.Context, ID uint, data map[string]interface{}) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var data model.User
	err := r.Db.WithContext(ctx).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
	var data model.User
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(ctx context.Context, data model.User) error{
	return r.Db.WithContext(ctx).Model(&model.User{}).Create(&data).Error
}

func (r *user) Update(ctx context.Context, ID uint, data map[string]interface{}) error {

	err := r.Db.WithContext(ctx).Where("id = ?", ID).Model(&model.User{}).Updates(data).Error
	return err

}



