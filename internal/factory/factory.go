package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	ProductRepository repository.Product
	UserRepository    repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		ProductRepository: repository.NewProduct(db),
		UserRepository:    repository.NewUser(db),
	}
}
