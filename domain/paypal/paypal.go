package paypal

import (
	"github.com/google/uuid"
	"github.com/luryon/go-ecommerce/model"
	"net/http"
)

type UseCase interface {
	ProcessRequest(header http.Header, body []byte) error
}

type UseCasePurchaseOrder interface {
	GetByID(ID uuid.UUID) (model.PurchaseOrder, error)
}
type UseCaseInvoice interface {
	Create(m *model.PurchaseOrder) error
}
