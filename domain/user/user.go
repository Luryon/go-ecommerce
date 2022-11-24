package user

import "github.com/luryon/go-ecommerce/model"

type UseCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Storage interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
