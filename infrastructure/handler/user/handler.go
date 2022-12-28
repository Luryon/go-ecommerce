package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/user"
	"github.com/luryon/go-ecommerce/infrastructure/handler/response"
	"github.com/luryon/go-ecommerce/model"
)

type handler struct {
	useCase   user.UseCase
	responser response.API
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	err := c.Bind(&m)
	if err != nil {
		return h.responser.BindFailed(err)
	}

	//TODO: verificar que la contraseña y mail no sean nulos o iguales a ""

	err = h.useCase.Create(&m)
	if err != nil {
		return h.responser.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.responser.Created(m))
}

// MySelf returns the data from my profile
func (h handler) MySelf(c echo.Context) error {
	ID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return h.responser.Error(c, "c.Get().(uuid.UUID)", errors.New("couldn´t parse the ID"))
	}

	u, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.responser.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.responser.OK(u))
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responser.Error(c, "userCase.GetAll()", err)
	}

	return c.JSON(h.responser.OK(users))
}
