package repository

import (
	"context"
	_ "strings"

	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
	Create(ctx context.Context, data model.User) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
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



