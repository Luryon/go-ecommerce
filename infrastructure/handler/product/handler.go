package product

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/product"
	"github.com/luryon/go-ecommerce/infrastructure/handler/response"
	"github.com/luryon/go-ecommerce/model"
)

type handler struct {
	useCase  product.UseCase
	response response.API
}

func newHandler(useCase product.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) Update(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}
	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "UseCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(productData))

}

func (h handler) GetAll(c echo.Context) error {
	products, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAll()", err)
	}

	//TODO: Crear paginacion del listado, ejemplo: LIMITE 100 productos por pagina, y poder listarlos

	return c.JSON(h.response.OK(products))
}
