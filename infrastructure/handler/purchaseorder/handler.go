package purchaseorder

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/purchaseorder"
	"github.com/luryon/go-ecommerce/infrastructure/handler/response"
	"github.com/luryon/go-ecommerce/model"
)

type handler struct {
	useCase  purchaseorder.UseCase
	response response.API
}

func newHandler(useCase purchaseorder.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.PurchaseOrder{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)

	}

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return h.response.Error(c, "c.Get().(uuid.UUID)", errors.New("can't parse uuid"))
	}

	m.UserID = userID
	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "userCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

//TODO: crear GetAll para recibir todas las ordenes, neceita token JWt, para un cliente en especifico
