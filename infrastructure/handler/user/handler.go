package user

import (
	"github.com/labstack/echo/v4"
	"github.com/luryon/go-ecommerce/domain/user"
	"github.com/luryon/go-ecommerce/model"
	"net/http"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}
}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	err := c.Bind(&m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.useCase.Create(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, m)
}

func (h handler) GetAll(c echo.Context) error {
	users, err := h.useCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}
