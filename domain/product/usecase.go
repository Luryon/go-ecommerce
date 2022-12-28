package product

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luryon/go-ecommerce/model"
	"time"
)

type Product struct {
	storage Storage
}

func New(s Storage) Product {
	return Product{s}
}

func (p Product) Create(m *model.Product) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	m.ID = ID
	if len(m.Images) == 0 {
		m.Images = []byte("[]")
	}
	if len(m.Features) == 0 {
		m.Features = []byte("{}")
	}

	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) Update(m *model.Product) error {
	if !m.HasID() {
		return fmt.Errorf("product: %w", model.ErrInvalidID)
	}

	if len(m.Images) == 0 {
		m.Images = []byte("[]")
	}
	if len(m.Features) == 0 {
		m.Features = []byte("{}")
	}
	m.UpdatedAt = time.Now().Unix()

	err := p.storage.Update(m)
	if err != nil {
		return err
	}

	return nil
}

func (p Product) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return fmt.Errorf("product: %w", model.ErrInvalidID)
	}

	return nil
}

func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	product, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (p Product) GetAll() (model.Products, error) {
	products, err := p.storage.GetAll()
	if err != nil {
		return model.Products{}, err
	}

	return products, nil
}
