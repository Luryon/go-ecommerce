package purchaseorder

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luryon/go-ecommerce/model"
)

type PurchaseOrder struct {
	storage Storage
}

func New(s Storage) PurchaseOrder {
	return PurchaseOrder{storage: s}
}
func (p PurchaseOrder) Create(m *model.PurchaseOrder) error {
	if err := m.Validate(); err != nil {
		return fmt.Errorf("purchaseOrder: %w", err)
	}

	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID

	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

func (p PurchaseOrder) GetByID(ID uuid.UUID) (model.PurchaseOrder, error) {
	purchaseOrder, err := p.storage.GetByID(ID)
	if err != nil {
		return model.PurchaseOrder{}, fmt.Errorf("purchaseOrder: %w", err)
	}

	return purchaseOrder, nil
}
