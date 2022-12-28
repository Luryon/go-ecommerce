package purchaseorder

import (
	"github.com/google/uuid"
	"github.com/luryon/go-ecommerce/model"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}

type Storage interface {
	Create(m *model.PurchaseOrder) error

	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}
